package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"homeWork2/middleware"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	Middlewares *middleware.Middleware
	Config      *Env
}

func (a *App) Initialize(e *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Config = e
	a.Router = mux.NewRouter()
	a.Middlewares = &middleware.Middleware{}

	a.InitializeRouter()
}

func (a *App) InitializeRouter() {
	m := alice.New(a.Middlewares.LoggingHandler)

	a.Router.Handle("/version", m.ThenFunc(a.version)).Methods("get")
	a.Router.Handle("/healthz", m.ThenFunc(a.healthz)).Methods("get")
}

func (a *App) Run(addr string) {
	//err := http.ListenAndServe(":" + os.Getenv("MY_SERVICE_PORT"), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
