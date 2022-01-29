package main

import (
	"log"

	"github.com/oussama4/command"
	"github.com/oussama4/commentify/app/handlers"
	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/store"
)

type Application struct {
	config    *config.Config
	logger    *log.Logger
	store     store.Store
	commander *command.Commander
}

func NewApp(logger *log.Logger, config *config.Config, store store.Store) *Application {
	cmder := command.New("commentify")

	app := &Application{
		config:    config,
		logger:    logger,
		store:     store,
		commander: cmder,
	}

	return app
}

func (app *Application) start() error {
	routes := handlers.Routes(app.store, app.logger)

	// register commands
	createAdminCmd := newAdminCmd(app.store)
	serverCmd := NewServerCmd(app.logger, routes)
	app.commander.Register("admin", createAdminCmd)
	app.commander.Register("server", serverCmd)

	return app.commander.Run()
}
