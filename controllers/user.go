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

func UserGetById(c *gin.Context) {
	var user *models.User

	result := initalizers.DB.First(&user, "id = ?", c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.GetOperationErrorResponse(result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, models.GetOperationSuccessResponse(user))
}

func UserUpdate(c *gin.Context) {
	var user *models.User
	resultFind := initalizers.DB.First(&user, "id = ?", c.Param("id"))
	if resultFind.Error != nil {
		c.JSON(http.StatusNotFound, models.GetOperationErrorResponse(resultFind.Error.Error()))
		return
	}
	var payload *models.UserUpdate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadGateway, models.GetOperationFailureResponse(err.Error()))
		return
	}

	updateUser := models.UserUpdate{
		Name:      payload.Name,
		Email:     payload.Email,
		UpdatedAt: time.Now(),
	}

	result := initalizers.DB.Model(&user).Updates(updateUser)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, models.GetOperationFailureResponse(result.Error.Error()))
		return
	}

	userResponse := &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, models.GetOperationSuccessResponse(userResponse))

}

func UserDelete(c *gin.Context) {
	var user *models.User
	resultFind := initalizers.DB.First(&user, "id = ?", c.Param("id"))
	if resultFind.Error != nil {
		c.JSON(http.StatusNotFound, models.GetOperationErrorResponse(resultFind.Error.Error()))
		return
	}

	result := initalizers.DB.Where("id = ?", c.Param("id")).Delete(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, models.GetOperationFailureResponse(result.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, models.GetOperationSuccessResponse(user))
}
