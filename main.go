package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bitcode-bin/expmgr/inmemory"
	"github.com/bitcode-bin/expmgr/logger"
)

func main() {
	//var port string
	//flag.StringVar(&port, "port", "9000", "server port")
	//flag.Parse()
	port := os.Getenv("PORT")

	if err := run(port); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start: %v", err)
		os.Exit(1)
	}
}

func run(port string) error {
	s := &server{
		wallet: inmemory.NewWallet(100),
		log:    logger.NewDefaultLogger(),
	}
	s.Init()

	s.log.Info("listening on " + port)
	http.ListenAndServe(":"+port, s)

	return nil
}
