package Controllers

import (
	"api-demo/app/ModelFilters"
	"api-demo/app/Models"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var conn = GetDb()
	defer conn.Close()

	var users Models.Users
	useFilter := ModelFilters.UserFilter(c)

	conn.Scopes(useFilter...).Preload("Profile").Preload("Role").Find(&users)

	Response(c, users)
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	var conn = GetDb()
	defer conn.Close()

	var user Models.User

	conn.Preload("Profile").Preload("Role").Find(&user, id)

	Response(c, user)
}
