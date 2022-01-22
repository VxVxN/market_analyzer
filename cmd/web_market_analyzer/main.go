package main

import (
	slog "log"

	"github.com/VxVxN/log"

	httpserver "github.com/VxVxN/market_analyzer/internal/server"
)

func main() {
	server, err := httpserver.InitServer()
	if err != nil {
		slog.Fatalln("Failed to init md server", err)
	}

	server.SetRoutes()

	if err = server.ListenAndServe("localhost:3030"); err != nil {
		log.Fatal.Printf("Failed to run router: %v", err)
	}
}
