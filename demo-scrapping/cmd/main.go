package main

import (
	"demo-scrapping/cmd/app"
	"demo-scrapping/config"
	"flag"
)

var pathFlag = flag.String("config", "./config.toml", "set toml path")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)
	app := app.NewApp(c)
	go app.Wait()
	app.Run()
}
