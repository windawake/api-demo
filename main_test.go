package main

import (
	"api-demo/app/Http/Controllers"
	"api-demo/app/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Active(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", 1)
}

func TestIndex(t *testing.T) {
	db, err := gorm.Open("mysql", "zhang:z1234567@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	db.SingularTable(true)
	defer db.Close()

	var users Models.Users
	db.Scopes(Active).Preload("Profile").Preload("Role").Find(&users)

	jUser, _ := json.Marshal(users)

	fmt.Printf("%v\n", string(jUser))
}

func TestShow(t *testing.T) {
	db, err := gorm.Open("mysql", "zhang:z1234567@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	db.SingularTable(true)
	defer db.Close()

	// var users Users
	// db.Find(&users)

	var user Models.User
	// var profile Models.Profile
	// var role Models.Role

	// db.Find(&user, 1)
	// db.Find(&profile, 2)
	// db.Model(&user).Related(&profile)

	// db.Model(&user).Related(&role)

	// db.Preload("Role").First(&user, 2)
	// db.Preload("Profile").Preload("Role").Find(&user, 1)
	db.Preload("Profile").Preload("Role").Find(&user, 2)
	// user.StatusName = "hello"
	// fmt.Printf("%+v\n", user)

	jUser, _ := json.Marshal(user)
	// j, _ := json.Marshal(role)
	// jProfile, _ := json.Marshal(profile)

	// fmt.Printf("%v %v\n", string(jUser), string(jProfile))

	fmt.Printf("%v\n", string(jUser))
}

func TestDemo01(t *testing.T) {
	router := gin.Default()
	router.GET("/user", Controllers.UserIndex)

	request, _ := http.NewRequest(http.MethodGet, "/user", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	got := response.Body.String()

	fmt.Println(got)

}

func modify(ages *[]int){
	*ages = append(*ages, 1);
}

func TestSlice(t *testing.T) {
	ages:=make([]int, 0)

	modify(&ages)
	fmt.Println(ages)
}
