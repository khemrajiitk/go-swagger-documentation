package controllers

import (
	"doc-generation/request"
	"doc-generation/response"
	"encoding/json"
	"log"
	"net/http"
)

type PingCtrl struct {
	l *log.Logger
}

func NewPingCtrl(l *log.Logger) *PingCtrl {
	return &PingCtrl{l}
}

func (pc *PingCtrl) PingGetController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	pingRes := response.Ping{
		Name: "Hello, Testing",
	}

	response, _ := json.Marshal(pingRes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (pc *PingCtrl) PingController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer r.Body.Close()
	var pingReq request.Ping

	err := json.NewDecoder(r.Body).Decode(&pingReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	pingRes := response.Ping{
		Name: "Hello, " + pingReq.Name,
	}

	response, _ := json.Marshal(pingRes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}