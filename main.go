package main

import (
	"log"

	"github.com/ddcad2030/gin3/initalizers"
	"github.com/ddcad2030/gin3/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config, err := initalizers.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load environment", err)
	}
	initalizers.ConnectDB(&config)
	initalizers.Migration()
}

func main() {
	r := gin.Default()
	routes.UserRoutes(r)

	r.Run()
}
