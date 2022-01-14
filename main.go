package main

import (
	"log"

	_ "github.com/quasilyte/go-ruleguard/dsl"
	"github.com/sailsforce/togo-read-micro/internal/config"
)

func main() {
	// 'true' is to indicate that we want a database with this config.
	if err := config.NewServiceConfig(false); err != nil {
		log.Fatalf("Error creating service config: %v\n", err)
	}
}
