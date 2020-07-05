package Controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Conn struct {
	*gorm.DB
}

type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

func GetDb() Conn {
	db, err := gorm.Open("mysql", "zhang:z1234567@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	db.SingularTable(true)
	conn := Conn{db}

	return conn
}

func Response(c *gin.Context, stu interface{}) {
	var resp Resp
	resp.Data = stu

	ret, _ := json.Marshal(resp)
	c.Data(200, "application/json; charset=utf-8", ret)
}
