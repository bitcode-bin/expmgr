package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start: %v", err)
		os.Exit(1)
	}
}

func run() error {
	s := &server{
		wallet: NewWallet(100),
		log:    NewLogger(),
	}
	s.Init()

	fmt.Println("listening")
	http.ListenAndServe(":9000", s)
	return nil
}
