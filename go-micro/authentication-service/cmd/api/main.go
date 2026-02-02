package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8082"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting Auth service...")

	// TODO: connect to DB
	// set up config
	app := Config{}

	// set up the web-server

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
