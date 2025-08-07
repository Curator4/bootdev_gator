package main

import (
	"log"
	"fmt"

	"github.com/curator4/bootdev_gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("couldn't read config err: %v", err)
	}

	cfg.SetUser("curator")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("couldn't read config err: %v", err)
	}

	fmt.Printf("ConfigL %+v\n", cfg)
}
