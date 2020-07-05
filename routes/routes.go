package routes

import (
	"api-demo/app/Http/Controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/", Controllers.UserIndex)
		user.GET("/:id", Controllers.UserShow)
		// user.POST("/", Controllers.UserStore)
		// user.PUT("/:id", Controllers.UserUpdate)
		// user.DELETE("/:id", Controllers.UserDestroy)
	}

	return router
}
