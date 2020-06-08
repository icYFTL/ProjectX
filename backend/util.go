package main

import (
    "github.com/gorilla/sessions"
    "net/http"
)

func (a *App) getSession(r *http.Request) (*sessions.Session, error) {
    session, err := a.Store.Get(r, "session")
    return session, err
}
