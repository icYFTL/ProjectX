package main

import (
	"net/http"
	"encoding/json"
)

func (a *App) login(w http.ResponseWriter, r *http.Request) {
    session, _ := a.getSession(r)

    var cred User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cred); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

    if cred.Username != "saveliy" ||  cred.Password != "pidor" {
		http.Error(w, "Wrong credentials", http.StatusUnauthorized)
		return
    }

    session.Values["auth"] = true
    session.Save(r, w)
    w.Write([]byte("success"))
}

func (a *App) register(w http.ResponseWriter, r *http.Request) {
    var cred User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cred); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

    // TODO: check if not exists and is ok

    // TODO: create user

    w.Write([]byte("success"))
}
