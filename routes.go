package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/volatiletech/authboss/lock"
	"github.com/volatiletech/authboss/remember"
)

func (s *server) loadRoutes() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(nosurfing)

	// Add the general auth middleware
	s.router.Use(s.auth.LoadClientStateMiddleware)
	// Remember middleware
	s.router.Use(remember.Middleware(s.auth))
	// Inject data such as csrf token to context
	s.router.Use(commonDataInjector)
	// Redirects logged in users if visiting the login or register page
	s.router.Use(s.redirectIfLoggedIn)

	s.router.NotFound(s.Handlers.use("404"))
	s.router.Mount("/auth", http.StripPrefix("/auth", s.auth.Config.Core.Router))
	s.router.Mount("/static/", http.FileServer(s.Files))

	s.router.Group(func(protect chi.Router) {
		// protected pages
		protect.Use(s.authMW)
		protect.Use(lock.Middleware(s.auth))

		protect.Get("/", s.Handlers.use("Dashboard"))

		protect.Get("/transactions/add", s.Handlers.use("AddTransactionForm"))
		protect.Post("/transactions/add", s.Handlers.use("AddTransaction"))
		protect.Post("/transactions", s.Handlers.use("AddTransaction"))

		protect.Get("/transaction/{id}", s.Handlers.use("ViewTransaction"))
		protect.Post("/transaction/{id}", s.Handlers.use("EditTransaction"))
		protect.Post("/transaction/{id}/delete", s.Handlers.use("DeleteTransaction"))

		protect.Get("/transactions", s.Handlers.use("ListTransactions"))
	})
}
