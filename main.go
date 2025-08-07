package main

import (
	"log"

	"github.com/curator4/bootdev_gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser("curator")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ConfigL %+v\n", cfg)
}
