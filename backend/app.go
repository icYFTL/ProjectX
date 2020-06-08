package main

import (
	"github.com/gorilla/handlers"
    "github.com/gorilla/sessions"
	"github.com/gorilla/mux"
    //"github.com/jackc/pgx/v4" // PostgreSQL driver and toolkit
    "database/sql"
	"net/http"
	"log"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
    Store  sessions.Store
	Logger http.Handler
}

func (a *App) InitializeRoutes() {
    a.Router.HandleFunc("/api/user",    a.register).Methods("POST")
	a.Router.HandleFunc("/api/user/login", a.login).Methods("POST")
}

func (a *App) Initialize(storeKey []byte) {
	// TODO: connect to db

    a.Store  = sessions.NewCookieStore(storeKey)
	a.Router = mux.NewRouter()
	a.Logger = handlers.CombinedLoggingHandler(os.Stdout, a.Router)
	//a.Router.Use(a.authMiddleware)
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	// https://stackoverflow.com/questions/38376226/how-to-allow-options-method-from-mobile-using-gorilla-handler
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(addr,
		handlers.CORS(headersOk, originsOk, methodsOk)(a.Logger)))
}
