package main

import (
	// "encoding/json"
	"net/http"
	// "github.com/go-chi/chi"
	// "github.com/spf13/viper"
	// "github.com/volatiletech/null"
	// "github.com/volatiletech/sqlboiler/boil"
	// amodels "github.com/stephenafamo/expense-tracker/auth_models"
	// "github.com/stephenafamo/expense-tracker/models"
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

	s.Handlers["Home"] = func() http.HandlerFunc {
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

	s.Handlers["404"] = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "404: Page not found", 404)
		}
	}
}
