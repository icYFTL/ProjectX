package main

import (
	"net/http"
)

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := a.getSession(r)
        if err != nil {
            http.Error(w, "Invalid session", http.StatusBadRequest)
            return
        }
        if session.Values["auth"] != true {
            http.Error(w, "Not authorized", http.StatusUnauthorized)
            return
        }
		next.ServeHTTP(w, r)
	})
}
