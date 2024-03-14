package main

import "github.com/deduardo/gobrax/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
