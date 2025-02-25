package main

import (
	"flag"
	"log"

	"github.com/demyforge/category-service/internal/app"
	"github.com/demyforge/category-service/internal/config"
)

func main() {
	configPath := flag.String("config", "configs/stub.toml", "path to config file")
	flag.Parse()

	conf, err := config.Load(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	a := app.New(conf.DSN())
	log.Fatal(a.Listen(conf.ListenAddr))
}
