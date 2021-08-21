package main

import (
	"log"
	"os"

	"github.com/oussama4/commentify/app"
	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/store/sqlite"
)

func main() {
	l := log.New(os.Stdout, "COMMENTIFY : ", log.LstdFlags|log.Lshortfile)

	// config
	cfg := config.New()

	// init database
	store, err := sqlite.Create(cfg.Store)
	if err != nil {
		l.Fatalln(err)
	}

	app := app.New(cfg, store)
	app.Start()
}
