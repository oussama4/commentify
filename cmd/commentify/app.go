package main

import (
	"log"
	"net/http"

	"github.com/oussama4/commentify/cmd/commentify/handlers"
	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/store"
)

type application struct {
	config config.Conf
	logger *log.Logger
	store  store.Store
}

func (app *application) start() {
	routes := handlers.Routes(app.store, app.logger)
	s := http.Server{
		Addr:    app.config.Server.Address,
		Handler: routes,
	}

	app.logger.Printf("server started listening on: %v", app.config.Server.Address)
	app.logger.Fatal(s.ListenAndServe())
}
