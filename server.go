package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultServer = "/Users/JP/www"
	defaultPort   = "8080"
)

func main() {
	serveDir := defaultServer
	servePort := defaultPort

	if len(os.Args) > 1 {
		serveDir = os.Args[1]
	}

	if len(os.Args) > 2 {
		servePort = os.Args[2]
	}

	server := http.FileServer(http.Dir(serveDir))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", servePort), server))
}
