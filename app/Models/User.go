package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID         int     `gorm:"primary_key" json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Status     int     `json:"status"`
	StatusName string  `json:"status_name"`
	CreatedAt  string  `json:"create_time"`
	RoleId     int     `json:"role_id"`
	Role       Role    `gorm:"ForeignKey:RoleId"`
	Profile    Profile `gorm:"ForeignKey:UserId"`
}

type Users []User

// 查询后触发事件
func (user *User) AfterFind(scope *gorm.Scope) error {
	if user.Status == 1 {
		user.StatusName = "启用"
	} else if user.Status == 0 {
		user.StatusName = "禁用"
	}

	return nil
}
