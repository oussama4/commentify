package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/oussama4/commentify/app/admin/commands"
	"github.com/oussama4/commentify/base/command"
	"github.com/oussama4/commentify/business/data/database"
)

func main() {
	l := log.New(os.Stdout, "COMMENTIFY: ", log.LstdFlags|log.Lshortfile)
	cfg := struct {
		DB struct {
			User         string `conf:"default:commentify"`
			Password     string `conf:"default:secret,mask"`
			Host         string `conf:"default:localhost"`
			Name         string `conf:"default:commentify"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}{}

	help, err := conf.Parse("", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
		}
		l.Println(fmt.Errorf("parsing config: %w", err))
	}

	db, err := database.Open(database.Config(cfg.DB))
	if err != nil {
		l.Fatal(err)
	}

	commander := command.New("admin")
	s := commands.NewSeed(db)
	commander.Register("seed", s)
	if err := commander.Run(); err != nil {
		log.Fatal(err)
	}
}
