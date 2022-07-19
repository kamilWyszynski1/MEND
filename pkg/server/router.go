package server

import (
	"github.com/gin-gonic/gin"
)

// NewRouter mounts routes and returns new gin router.
func NewRouter(userController *UserController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			userGroup.POST("/", userController.Create)
			userGroup.GET("/:id", userController.Get)
			userGroup.PUT("/:id", userController.Update)
			userGroup.DELETE("/:id", userController.Delete)
		}
	}
	return router
}
