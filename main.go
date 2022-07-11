package main

import (
	"mini_project/config"
	"mini_project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":80")
}
