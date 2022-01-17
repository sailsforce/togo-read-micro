package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/quasilyte/go-ruleguard/dsl"
	"github.com/sailsforce/togo-read-micro/internal"
	"github.com/sailsforce/togo-read-micro/internal/config"
)

func main() {
	// 'true' is to indicate that we want a database with this config.
	if err := config.NewServiceConfig(true); err != nil {
		log.Fatalf("Error creating service config: %v\n", err)
	}

	router := internal.Routes()

	config.Logger.Info("TOGO-READ-MICRO running on PORT: ", os.Getenv("PORT"))
	config.Logger.Debug(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
