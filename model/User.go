package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User : the user struct definition
type User struct {
	Abstract 							`sql:"embedded;prefix:-"`
	Email     string 			`gorm:"type:varchar(100);unique_index"json:"email"`
	Firstname string 			`json:"firstname"`
	Lastname  string 			`json:"lastname"`
	Accesslevel       int    			`json:"accesslevel"`
	Dateofbirth time.Time    			`json:"dateofbirth"`
	Password  string 			`json:"-"`
}

// TableName : Gorm related
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate : Gorm hook
func (u *User) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("Password", hashPassword(u.Password))
	scope.SetColumn("UUID", uuid.NewV4().String())
	return
}

// hashPassword : simple password hashing method
func hashPassword(password string) (string) {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}

// CheckPasswordHash : Compare password with a hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
