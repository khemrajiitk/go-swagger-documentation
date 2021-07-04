package routers

import (
	controllers "doc-generation/controller"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
)

type PingRouter struct {
	l  *log.Logger
	pc *controllers.PingCtrl
}

func NewPingRouter(l *log.Logger) *PingRouter {
	pc := controllers.NewPingCtrl(l)

	return &PingRouter{l, pc}
}

func (pr *PingRouter) Init(router *mux.Router) *mux.Router {
	router.Handle("/ping-check",
		negroni.New(
			negroni.HandlerFunc(pr.pc.PingGetController),
		)).Methods("GET")

	router.Handle("/ping",
		negroni.New(
			negroni.HandlerFunc(pr.pc.PingController),
		)).Methods("POST")

	return router
}
