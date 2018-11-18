package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/stephenafamo/expense-tracker/models"
)

type ConnQuery struct {
	First    *int
	Last     *int
	Before   *string
	After    *string
	Type     *string
	Category *string
	From     *time.Time
	To       *time.Time
}

type ConnectionParams struct {
	PerPage       int
	Pid           string
	OrderDir      string
	OrderCol      string
	FwdPagination bool
	SortSign      *string
	Data          map[string]interface{}
	Type          *string
	Category      *string
	From          *time.Time
	To            *time.Time
}

type TransactionConnection struct {
	Transactions    []*models.Transaction
	HasPreviousPage bool
	HasNextPage     bool
	Total           int
	FirstPid        string
	LastPid         string
	Type            string
	Category        string
	From            string
	To              string
}

func defaultData() map[string]interface{} {
	data := make(map[string]interface{})
	data["order_col"] = "created_at"
	data["order_dir"] = "desc"

	return data
}

func getConnParams(ctx context.Context, args ConnQuery, mods ...string) (params *ConnectionParams, err error) {

	params = &ConnectionParams{
		Data: defaultData(),
	}
	params.FwdPagination = true

	sort_sign_asc := true
	sign := ">"
	params.SortSign = &sign

	if args.After != nil {
		params.Pid = *args.After
	} else if args.Before != nil {
		params.Pid = *args.Before
		sort_sign_asc = toggle(sort_sign_asc)
	}

	params.OrderCol = params.Data["order_col"].(string)
	params.OrderDir = params.Data["order_dir"].(string)
	if len(mods) > 0 {
		baseTable := mods[0]
		params.OrderCol = baseTable + "." + params.OrderCol
	}

	params.PerPage = 24
	if args.First != nil {
		params.PerPage = *args.First
	} else if args.Last != nil {
		params.PerPage = *args.Last
		params.FwdPagination = false
	}

	if params.PerPage > 24 {
		params.PerPage = 24
	}

	if params.OrderDir == "desc" {
		sort_sign_asc = toggle(sort_sign_asc)
	}

	if !sort_sign_asc {
		sign := "<"
		params.SortSign = &sign
	}

	if args.After == nil && args.Before == nil {
		params.SortSign = nil
	}

	if !params.FwdPagination {
		if params.OrderDir == "desc" {
			params.OrderDir = "asc"
		} else {
			params.OrderDir = "desc"
		}
	}

	params.Type = args.Type
	params.Category = args.Category
	params.From = args.From
	params.To = args.To

	return
}

func getConnQuery(ctx context.Context, r *http.Request) (args ConnQuery, err error) {
	first := r.FormValue("first")
	if first != "" {
		var i int
		i, err = strconv.Atoi(first)
		if err != nil {
			return
		}
		args.First = &i
	}

	last := r.FormValue("last")
	if last != "" {
		var i int
		i, err = strconv.Atoi(last)
		if err != nil {
			return
		}
		args.Last = &i
	}

	before := r.FormValue("before")
	if before != "" {
		args.Before = &before
	}

	after := r.FormValue("after")
	if after != "" {
		args.After = &after
	}

	theType := r.FormValue("type")
	if theType != "" {
		args.Type = &theType
	}

	category := r.FormValue("category")
	if category != "" {
		args.Category = &category
	}

	var From, To time.Time

	from := r.FormValue("from")
	if from != "" {
		From, err = time.Parse("2006-01-02", from)
		if err != nil {
			return
		}
		args.From = &From
	}

	to := r.FormValue("to")
	if to != "" {
		To, err = time.Parse("2006-01-02", to)
		if err != nil {
			return
		}
		args.To = &To
	}

	return
}
