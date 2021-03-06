package main

import (
	//"regexp"
	"context"
	"encoding/json"
	pgx "github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	//"log"
)

func (a *App) register(w http.ResponseWriter, r *http.Request) {
    var creds User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&creds); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	// TODO: check with regexp
	/*match, _ := regexp.MatchString("ты гандон", creds.Username)
	if !match {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}*/

	if len(creds.Username) < 4 || creds.Username[0] == '_' {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// check if user exists
	var found bool
	result := a.DB.QueryRow(context.Background(),
			"select exists(select 1 from users where username=$1)", creds.Username)
	if err := result.Scan(&found); err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	if found {
		http.Error(w, "User exists", http.StatusBadRequest)
		return
	}

	// bcrypt hash the password using default 10 complexity
	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 10)

	 _, err = a.DB.Query(context.Background(),
	 	"insert into users (username, password, role, status) values ($1, $2, 'guest', 'active')",
		creds.Username, string(hash));
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

    w.Write([]byte("success"))
}


func (a *App) login(w http.ResponseWriter, r *http.Request) {
	session, err := a.getSession(r);
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

    var creds User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&creds); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	result := a.DB.QueryRow(context.Background(),
			"select id, password from users where username=$1", creds.Username)

	var found_creds User
	var user_id int
	err = result.Scan(&user_id, &found_creds.Password)
	if err != nil {
		// If an entry with the username does not exist, send 401 status
		if err == pgx.ErrNoRows {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		// If any other error occurred, send a 500 status
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := bcrypt.CompareHashAndPassword([]byte(found_creds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	// Authorization successful
    session.Values["user_id"] = user_id
    session.Values["auth"] = true
    session.Save(r, w)
    w.Write([]byte("success"))
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.getSession(r);
	if  err != nil || session.Values["auth"] != true {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}
	// unauthorization successful
    //session.Values["user_id"] = 0
    //session.Values["auth"] = false
	session.Options.MaxAge = -1
    session.Save(r, w)
    w.Write([]byte("success"))
}
