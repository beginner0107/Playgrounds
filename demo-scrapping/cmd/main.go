package main

import (
	"demo-scrapping/config"
	"flag"
	"fmt"
)

var pathFlag = flag.String("config", "./config.toml", "set toml path")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)
	fmt.Println(c)
}
