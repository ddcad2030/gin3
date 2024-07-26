package controllers

import (
	"net/http"
	"strings"

	"github.com/ddcad2030/gin3/initalizers"
	"github.com/ddcad2030/gin3/models"
	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	users := []models.User{}
	result := initalizers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": http.StatusText(http.StatusNotFound),
		})
		return
	}
	c.JSON(http.StatusOK, users)

}

func UserCreate(c *gin.Context) {
	var payload *models.UserCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	newUser := models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Password: payload.Password,
	}
	result := initalizers.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "fail",
				"message": "Email already exists",
			})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"stauts": "success",
		"data":   newUser,
	})
}
