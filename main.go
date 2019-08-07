package main

import (
	"./app"
	"./config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initiailize(config)
	app.Run(":3000")
}
