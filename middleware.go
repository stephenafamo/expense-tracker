package main

import (
	"context"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/volatiletech/authboss"
)

func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate XSRF Token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}

func (s *server) authMW(h http.Handler) http.Handler {
	return authboss.Middleware(s.auth, true, false, false)(h)
}

func (s *server) redirectIfLoggedIn(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pid, err := s.auth.CurrentUserID(r)
		checkError(err)

		mountPath := s.auth.Config.Paths.Mount
		switch r.URL.Path {
		case mountPath + "/login", mountPath + "/register":
			if pid != "" {
				ro := authboss.RedirectOptions{
					Code:             http.StatusTemporaryRedirect,
					RedirectPath:     s.auth.Paths.AuthLoginOK,
					FollowRedirParam: true,
				}
				if err := s.auth.Core.Redirector.Redirect(w, r, ro); err != nil {
					checkError(err)
				}
			}
		}
		h.ServeHTTP(w, r)
	})
}

func commonDataInjector(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := authboss.HTMLData{
			"csrf_token":    nosurf.Token(r),
			"baseUrl":       getBaseUrl().String(),
			"flash_success": authboss.FlashSuccess(w, r),
			"flash_error":   authboss.FlashError(w, r),
		}
		r = r.WithContext(context.WithValue(r.Context(), authboss.CTXKeyData, data))
		handler.ServeHTTP(w, r)
	})
}
