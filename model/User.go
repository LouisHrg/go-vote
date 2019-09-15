package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Content string `json:"content"`
}


func (u *User) TableName() string {
	return "users"
}
