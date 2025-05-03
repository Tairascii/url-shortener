package main

import "github.com/Tairascii/url-shortener/internal"

func main() {
	app := internal.NewApp()
	app.Start()
}
