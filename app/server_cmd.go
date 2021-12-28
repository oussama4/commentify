package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type serverCmd struct {
	synopsis string
	port     int
	handler  http.Handler
	logger   *log.Logger
	fs       *flag.FlagSet
}

func (sc *serverCmd) Synopsis() string {
	return sc.synopsis
}

func NewServerCmd(logger *log.Logger, handler http.Handler) *serverCmd {
	fs := flag.NewFlagSet("server", flag.PanicOnError)
	p := fs.Int("p", 8888, "TCP port for the server to listen on.")

	if len(os.Args) > 1 {
		fs.Parse(os.Args[2:])
	}

	serverCmd := &serverCmd{
		synopsis: "start commentify server",
		port:     *p,
		handler:  handler,
		logger:   logger,
		fs:       flag.NewFlagSet("server", flag.PanicOnError),
	}
	return serverCmd
}

func (sc *serverCmd) Help() string {
	u := `usage: commentify server -p <PORT>
		start the API server`
	return u
}

func (sc *serverCmd) Run() error {
	address := fmt.Sprintf(":%v", sc.port)
	s := http.Server{
		Addr:    address,
		Handler: sc.handler,
	}

	sc.logger.Printf("server started listening on: %v", address)
	return s.ListenAndServe()
}
