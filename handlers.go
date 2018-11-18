package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/stephenafamo/expense-tracker/models"
	"github.com/volatiletech/authboss"
)

type Handlers map[string]func() http.HandlerFunc

func (h Handlers) use(name string) http.HandlerFunc {
	method, ok := h[name]
	if ok == false {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Handler not found", 500)
		}
	}
	return method()
}

func (s *server) registerHandlers() {
	s.Handlers = make(map[string]func() http.HandlerFunc)

	s.Handlers["Dashboard"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error

			info := make(map[string]interface{})
			info["BaseUrl"] = getBaseUrl().String()
			info["IsHomePage"] = true

			data, err := s.render("index", info)
			checkError(err)

			w.Write([]byte(data))
		}
	}

	s.Handlers["ListTransactions"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error

			ctx := r.Context()

			info := make(map[string]interface{})
			info["BaseUrl"] = getBaseUrl().String()
			info["auth_data"] = ctx.Value(authboss.CTXKeyData)

			tConn, err := s.getTransactions(ctx, r)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve transactions", 500)
				return
			}
			info["tConn"] = tConn

			types, err := models.Types().All(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve types", 500)
				return
			}
			info["types"] = types

			categories, err := models.Categories().All(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve categories", 500)
				return
			}
			info["categories"] = categories

			data, err := s.render("transactions", info)
			checkError(err)

			w.Write([]byte(data))
		}
	}

	s.Handlers["DeleteTransaction"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error

			ctx := r.Context()
			id := chi.URLParam(r, "id")

			transaction, err := models.FindTransaction(ctx, s.DB, id)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve transaction", 500)
				return
			}

			_, err = transaction.Delete(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot delete transaction", 500)
				return
			}

			http.Redirect(w, r, r.Header.Get("Referer"), 302)
		}
	}

	s.Handlers["ViewTransaction"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error

			ctx := r.Context()
			id := chi.URLParam(r, "id")

			info := make(map[string]interface{})
			info["BaseUrl"] = getBaseUrl().String()
			info["auth_data"] = ctx.Value(authboss.CTXKeyData)

			transaction, err := models.Transactions(
				qm.Load("Category"),
				qm.Load("Type"),
				qm.Where("id=?", id),
			).One(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve transaction", 500)
				return
			}
			info["transaction"] = transaction
			info["tDate"] = transaction.CreatedAt.Time.Format("2006-01-02")

			types, err := models.Types().All(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve types", 500)
				return
			}
			info["types"] = types

			categories, err := models.Categories().All(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve categories", 500)
				return
			}
			info["categories"] = categories

			data, err := s.render("view-transaction", info)
			checkError(err)

			w.Write([]byte(data))
		}
	}

	s.Handlers["AddTransactionForm"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error

			ctx := r.Context()
			info := make(map[string]interface{})
			info["BaseUrl"] = getBaseUrl().String()
			info["auth_data"] = ctx.Value(authboss.CTXKeyData)

			types, err := models.Types().All(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve types", 500)
				return
			}
			info["types"] = types

			categories, err := models.Categories().All(ctx, s.DB)
			if err != nil {
				checkError(err)
				http.Error(w, "Cannot retrieve categories", 500)
				return
			}
			info["categories"] = categories

			data, err := s.render("add-transaction", info)
			checkError(err)

			w.Write([]byte(data))
		}
	}

	s.Handlers["AddTransaction"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			ctx := r.Context()

			t := &models.Transaction{}

			t.Narration = r.FormValue("narration")

			err = t.Amount.Scan(r.FormValue("amount"))
			if err != nil {
				checkError(err)
				http.Error(w, "Error adding product", 500)
				return
			}

			if r.FormValue("exchange-rate") != "" {
				err = t.ExchangeRate.Scan(r.FormValue("exchange-rate"))
				if err != nil {
					checkError(err)
					http.Error(w, "Error adding product", 500)
					return
				}
			}

			if r.FormValue("category") != "" {
				t.CategoryID = null.StringFrom(r.FormValue("category"))
			}

			if r.FormValue("type") != "" {
				t.TypeID = null.StringFrom(r.FormValue("type"))
			}

			if r.FormValue("currency") != "" {
				t.Currency = r.FormValue("currency")
			}

			if r.FormValue("date") != "" {
				var theTime time.Time
				theTime, err = time.Parse("2006-01-02", r.FormValue("date"))
				if err != nil {
					checkError(err)
					http.Error(w, "Error adding product", 500)
					return
				}

				t.CreatedAt = null.TimeFrom(theTime)
			}

			err = t.Insert(ctx, s.DB, boil.Blacklist("amount_local"))
			if err != nil {
				checkError(err)
				http.Error(w, "Error adding product", 500)
				return
			}

			authboss.PutSession(w, authboss.FlashSuccessKey, "Transaction added successfully")

			http.Redirect(w, r, r.Header.Get("Referer"), 302)
		}
	}

	s.Handlers["404"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "404: Page not found", 404)
		}
	}
}
