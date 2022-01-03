package app

import (
	"log"
	"os"

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

func New(config *config.Config, store store.Store) *Application {
	l := log.New(os.Stdout, "COMMENTIFY : ", log.LstdFlags|log.Lshortfile)
	cmder := command.New("commentify")

	app := &Application{
		config:    config,
		logger:    l,
		store:     store,
		commander: cmder,
	}

	return app
}

func (app *Application) Start() {
	routes := handlers.Routes(app.store, app.logger)

	// register commands
	createAdminCmd := newAdminCmd(app.store)
	serverCmd := NewServerCmd(app.logger, routes)
	app.commander.Register("admin", createAdminCmd)
	app.commander.Register("server", serverCmd)

	app.logger.Fatalln(app.commander.Run())
}
