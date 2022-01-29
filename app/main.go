package main

import (
	"log"
	"os"

	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/store/postgres"
)

func main() {
	l := log.New(os.Stdout, "COMMENTIFY: ", log.LstdFlags|log.Lshortfile)
	if err := run(l); err != nil {
		l.Fatal(err)
	}
}

func run(logger *log.Logger) error {
	// config
	cfg, err := config.LoadConf("./.config.json")
	if err != nil {
		return err
	}

	// init database
	store, err := postgres.Create(cfg.Store.Dsn)
	if err != nil {
		return err
	}

	app := NewApp(logger, cfg, store)
	return app.start()
}
