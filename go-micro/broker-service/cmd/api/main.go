package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort           = "8080" //80
	HeaderContentType = "Content-Type"
	JSONContentType   = "application/json"
)

type Config struct{}

func main() {

	app := Config{}
	log.Printf("Starting AMQP service on port %s\n", webPort)

	// http server

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
		return
	}

}
