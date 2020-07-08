package ModelFilters

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func WrapFilter(handler func(db *gorm.DB, req *gin.Context) *gorm.DB, c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return handler(db, c)
	}
}

type Request struct {
	Status     string   `form:"status"`
	CreateTime []string `form:"create_time"`
}

func UserFilter(c *gin.Context) []func(*gorm.DB) *gorm.DB {
	useFunc := make([]func(*gorm.DB) *gorm.DB, 0)

	status := c.Query("status")
	if status != "" {
		useFunc = append(useFunc, WrapFilter(Status, c))
	}

	createTime := c.QueryArray("create_time")
	if len(createTime) == 2 && createTime[0] < createTime[1] {
		useFunc = append(useFunc, WrapFilter(CreateTime, c))
	}

	return useFunc
}

// 搜索条件 where status = ?
func Status(db *gorm.DB, c *gin.Context) *gorm.DB {

	return db.Where("status = ?", c.Query("status"))
}

func CreateTime(db *gorm.DB, c *gin.Context) *gorm.DB {
	createTime := c.QueryArray("create_time")
	return db.Where("created_at >= ? and created_at < ?", createTime[0], createTime[1])
}
