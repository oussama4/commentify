package main

import (
	"log"
	"os"

	"github.com/oussama4/commentify/app"
	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/store/postgres"
)

func main() {
	l := log.New(os.Stdout, "COMMENTIFY : ", log.LstdFlags|log.Lshortfile)

	// config
	cfg, err := config.LoadConf("./config.json")
	if err != nil {
		l.Fatalln(err)
	}

	// init database
	store, err := postgres.Create(cfg.Store.Dsn)
	if err != nil {
		l.Fatalln(err)
	}

	app := app.New(cfg, store)
	app.Start()
}
