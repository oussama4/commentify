package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/oussama4/commentify/app/handlers"
	"github.com/oussama4/commentify/business/data/store/postgres"
)

func main() {
	l := log.New(os.Stdout, "COMMENTIFY: ", log.LstdFlags|log.Lshortfile)
	if err := run(l); err != nil {
		l.Fatal(err)
	}
}

func run(logger *log.Logger) error {
	// config
	cfg := struct {
		Server struct {
			Addr string `conf:"default:0.0.0.0:8888"`
		}
		Store struct {
			Name string `conf:"default:postgres"`
			Url  string `conf:"default:postgresql://commentify:password@localhost:5432/commentify?sslmode=disable"`
		}
	}{}

	help, err := conf.Parse("", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// database
	store, err := postgres.Create(cfg.Store.Url)
	if err != nil {
		return err
	}

	// start server
	routes := handlers.Routes(store, logger)
	s := http.Server{
		Addr:    cfg.Server.Addr,
		Handler: routes,
	}

	logger.Printf("server started listening on: %v", cfg.Server.Addr)
	return s.ListenAndServe()
}
