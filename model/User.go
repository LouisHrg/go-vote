package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// User : the user struct definition
type User struct {
	gorm.Model
	UUID      string `gorm:"unique_index"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int
	Password  string
}

// TableName : Gorm related
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate : Gorm hook
func (u *User) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return
}

// AfterFind : Gorm hook
func (u *User) AfterFind() (err error) {
	u.Password = ""
	u.ID = 0
	return
}
