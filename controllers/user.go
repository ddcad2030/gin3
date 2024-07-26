package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/ddcad2030/gin3/initalizers"
	"github.com/ddcad2030/gin3/models"
	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	users := []models.User{}
	result := initalizers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.GetOperationSuccessResponse(users))
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
		Email:     payload.Email,
		Name:      payload.Name,
		Password:  payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := initalizers.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			c.JSON(http.StatusConflict, models.GetOperationFailureResponse("Email already exists"))
			return
		}
		c.JSON(http.StatusBadGateway, models.GetOperationErrorResponse(result.Error.Error()))
		return
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Password:  newUser.Password,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	c.JSON(http.StatusOK, models.GetOperationSuccessResponse(userResponse))
}
