package main

import (
	"api/config"
	"api/core"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if err := config.ValidateEnv(); err != nil {
		log.Println("environment variables validation failure:\n", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	core.InitRoutes(mux)

	core.InitHttpServer(net.JoinHostPort(config.HOST, config.PORT), mux)

	if err := core.Srv.ListenAndServe(); err != nil {
		log.Println("server exiting:", err)
		os.Exit(1)
	}
}
