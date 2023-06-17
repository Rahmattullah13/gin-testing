package main

import (
	"gin-test/config"
	"gin-test/routes"
)

func main() {
	config.DBConnect()
	r := routes.SetupRoutes()
	r.Run("localhost:8000")
}
