package main

import (
	"github.com/ashhhleyyy/clickd/model"
	"github.com/ashhhleyyy/clickd/routes"
)

func main() {
	model.InitDB()

	r := routes.SetupRoutes()
	r.Run()
}
