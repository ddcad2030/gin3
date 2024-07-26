package initalizers

import (
	"log"

	"github.com/ddcad2030/gin3/models"
)

func Migration() {
	log.Println("migration")
	DB.Debug().AutoMigrate(&models.User{})
}
