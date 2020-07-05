package ModelFilters

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserFilter(c *gin.Context) []func(*gorm.DB) *gorm.DB {
	useFunc := make([]func(*gorm.DB) *gorm.DB, 0)

	status := c.Query("status")
	if status != "" {
		useFunc = append(useFunc, Status(status))
	}

	return useFunc
}

func Status(value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", value)
	}
}
