package routers

import (
	"github.com/gorilla/mux"
	"log"
)

type Router struct {
	l   *log.Logger
}

func NewRouter(l *log.Logger) *Router {
	return &Router{l}
}

func (r *Router) InitRoutes() *mux.Router {
	pr := NewPingRouter(r.l)

	muxRouter := mux.NewRouter()
	router := muxRouter.PathPrefix("/v1").Subrouter()
	router = pr.Init(router)

	return router
}
