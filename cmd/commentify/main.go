package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/oussama4/commentify/cmd/commentify/handlers"
	"github.com/oussama4/commentify/config"
	"github.com/oussama4/commentify/store/sqlite"
)

func main() {
	l := log.New(os.Stdout, "COMMENTIFY : ", log.LstdFlags|log.Lshortfile)

	// config
	a := flag.String("a", ":8888", `TCP address for the server to listen on in the form "host:port".`)
	e := flag.String("e", "prod", `Environment, either "prod" or "dev".`)
	flag.Parse()

	cfg := config.New()
	cfg.Environment = *e
	cfg.Server.Address = *a

	// init database
	store, err := sqlite.Create(cfg.Store)
	if err != nil {
		l.Fatalln(err)
	}
	l.Println("database initialized")

	// start server
	l.Println("starting server")

	s := http.Server{
		Addr:    cfg.Server.Address,
		Handler: handlers.Routes(store),
	}
	l.Fatal(s.ListenAndServe())
}
