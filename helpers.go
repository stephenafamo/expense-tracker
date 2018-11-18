package main

import (
	"context"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/stephenafamo/expense-tracker/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type PartialProvider struct {
	server
}

func (p PartialProvider) Get(name string) (string, error) {
	return p.getTemplate("/partials/" + name)
}

func getBaseUrl() *url.URL {

	baseUrl, err := url.Parse(viper.GetString("panel_url"))
	checkError(err)

	baseUrl.Path = filepath.Join(baseUrl.Path, "/") // will be "/" if empty

	if baseUrl.Path != "/" {
		baseUrl.Path = baseUrl.Path + "/" // add trailing slash
	}

	return baseUrl
}

func toggle(val bool) bool {
	if val {
		return false
	} else {
		return true
	}
}

func (s *server) getTransactions(ctx context.Context, r *http.Request) (tConn TransactionConnection, err error) {
	var hasPrevPage, hasNextPage bool
	var firstPid, lastPid string

	args, err := getConnQuery(ctx, r)
	if err != nil {
		return
	}

	params, err := getConnParams(ctx, args)
	if err != nil {
		return
	}

	if args.After != nil {
		hasPrevPage = true
	} else if args.Before != nil {
		hasNextPage = true
	}

	baseTable := "transactions"

	query := models.Transactions()
	qm.Apply(query.Query, qm.Select("transactions.*"))
	qm.Apply(query.Query, qm.Load("Type"))
	qm.Apply(query.Query, qm.Load("Category"))
	if params.Category != nil {
		qm.Apply(query.Query, qm.InnerJoin("categories on categories.id = transactions.category_id"))
		qm.Apply(query.Query, qm.Where("categories.name=?", params.Category))
	}
	if params.Type != nil {
		qm.Apply(query.Query, qm.InnerJoin("types on types.id = transactions.type_id"))
		qm.Apply(query.Query, qm.Where("types.name=?", params.Type))
	}
	if params.From != nil {
		qm.Apply(query.Query, qm.Where("transactions.created_at > ?", params.From))
	}
	if params.To != nil {
		qm.Apply(query.Query, qm.Where("transactions.created_at  < ?", params.To))
	}

	countQuery := models.Transactions()
	countQueryCopy := *query.Query     //create a copy so we don't override
	countQuery.Query = &countQueryCopy //create a copy so we don't override
	count, err := countQuery.Count(ctx, s.DB)
	if err != nil {
		return
	}

	// Now order the query
	qm.Apply(query.Query, qm.OrderBy(params.OrderCol+" "+params.OrderDir))
	qm.Apply(query.Query, qm.Limit(params.PerPage+1))
	if params.SortSign != nil {
		qm.Apply(query.Query, qm.Where(params.OrderCol+" "+*params.SortSign+"  (select "+params.OrderCol+" from "+baseTable+" where id =?)", params.Pid))
	}

	transactions, err := query.All(ctx, s.DB)

	if len(transactions) > params.PerPage {
		if params.FwdPagination {
			hasNextPage = true
		} else {
			hasPrevPage = true
		}
		transactions = transactions[:len(transactions)-1]
	}

	if !params.FwdPagination {
		for i := len(transactions)/2 - 1; i >= 0; i-- {
			opp := len(transactions) - 1 - i
			transactions[i], transactions[opp] = transactions[opp], transactions[i]
		}
	}

	if len(transactions) > 0 {
		firstPid = transactions[0].ID
		lastPid = transactions[len(transactions)-1].ID
	}

	tConn = TransactionConnection{
		Transactions:    transactions,
		HasPreviousPage: hasPrevPage,
		HasNextPage:     hasNextPage,
		Total:           int(count),
		FirstPid:        firstPid,
		LastPid:         lastPid,
	}

	if params.Type != nil {
		tConn.Type = *params.Type
	}

	if params.Category != nil {
		tConn.Category = *params.Category
	}

	if params.From != nil {
		tConn.From = params.From.Format("2006-01-02")
	}

	if params.To != nil {
		tConn.To = params.To.Format("2006-01-02")
	}

	return
}
