package routes

import (
	"net/http"

	"github.com/ddcad2030/gin3/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": http.StatusText(http.StatusOK),
		})
	})

	r.GET("/users", controllers.UserGet)
	r.GET("/users/:id", controllers.UserGetById)
	r.POST("/users", controllers.UserCreate)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.DELETE("/users/:id", controllers.UserDelete)
}
