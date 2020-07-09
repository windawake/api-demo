package ModelFilters

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddFilter(useFunc *[]func(db *gorm.DB) *gorm.DB, handler func(req *gin.Context) func(db *gorm.DB) *gorm.DB, c *gin.Context) {
	ret := handler(c)
	if ret != nil {
		*useFunc = append(*useFunc, ret)
	}

}

func UserFilter(c *gin.Context) []func(db *gorm.DB) *gorm.DB {
	useFunc := make([]func(db *gorm.DB) *gorm.DB, 0)

	AddFilter(&useFunc, Status, c)
	AddFilter(&useFunc, CreateTime, c)

	return useFunc
}

// 搜索条件 where status = ?
func Status(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	status := c.Query("status")
	if status == "" {
		return nil
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", c.Query("status"))
	}
	
}

// 搜索条件 created_at >= ? and created_at < ?
func CreateTime(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	createTime := c.QueryArray("create_time")
	if len(createTime) != 2 || createTime[0] >= createTime[1] {
		return nil
	}
	
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("created_at >= ? and created_at < ?", createTime[0], createTime[1])
	}
}
