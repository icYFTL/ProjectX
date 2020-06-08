package main

import (
	"net/http"
)

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := a.getSession(r)
        if err != nil {
            http.Error(w, "Failed to get session", http.StatusInternalServerError)
            return
        }
        if session.Values["auth"] != true {
            http.Error(w, "Please authenticate", http.StatusUnauthorized)
            return
        }
		next.ServeHTTP(w, r)
	})
}
