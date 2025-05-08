package main

import (
	"github.com/Tairascii/url-shortener/internal"
	"github.com/Tairascii/url-shortener/internal/infra/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	app := internal.NewApp(cfg)
	if err = app.Start(); err != nil {
		panic(err)
	}
}
