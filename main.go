package main

import (
	"context"
	routers "doc-generation/router"
	"fmt"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "document-generation-testing", log.LstdFlags)

	fmt.Println("Starting server...")

	// router := routers.InitRoutes()
	router := routers.NewRouter(l)

	n := negroni.Classic()
	n.UseHandler(router.InitRoutes())

	// create a new server
	s := http.Server{
		Addr:         ":8000",           // configure the bind address
		Handler:      n,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  60 * time.Second,  // max time to read request from the client
		WriteTimeout: 60 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Started server")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
