package main

import (
	"github.com/gorilla/handlers"
    "github.com/gorilla/sessions"
	"github.com/gorilla/mux"
    "github.com/jackc/pgx/v4/pgxpool" // PostgreSQL driver and toolkit
    "net/http"
	"context"
	"log"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *pgxpool.Pool
    Store  sessions.Store
	Logger http.Handler
}

func (a *App) InitializeRoutes() {
    a.Router.HandleFunc("/api/user",        a.register).Methods("POST")
	a.Router.HandleFunc("/api/user/login",  a.login).Methods("POST")
	a.Router.HandleFunc("/api/user/logout", a.logout).Methods("GET")
}

func (a *App) Initialize(storeKey []byte, dbURL string) {
    pool, err := pgxpool.Connect(context.Background(), dbURL)
    if err != nil {
        //fmt.Printf("pgerr: %v\n", err)
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    // defer pool.Close()
    log.Printf("Connected to database.")

	a.DB	 = pool
    //a.Store  = sessions.NewCookieStore(storeKey)
	a.Store  = sessions.NewFilesystemStore("", storeKey)
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
