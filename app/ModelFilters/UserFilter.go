package ModelFilters

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserFilter(c *gin.Context) []func(*gorm.DB) *gorm.DB {
	useFunc := make([]func(*gorm.DB) *gorm.DB, 0)
	useFunc = append(useFunc, Active)

	return useFunc
}

func Active(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", 1)
}
