package model

import (
  "github.com/jinzhu/gorm"
  "github.com/satori/go.uuid"
  "golang.org/x/crypto/bcrypt"
  "time"
)

// User : the user struct definition
type User struct {
  Abstract              `sql:"embedded;prefix:-"`
  Email       string    `valid:"email"gorm:"type:varchar(100);unique_index"json:"email"`
  Firstname   string    `valid:"stringlength(2|15)"json:"firstname"`
  Lastname    string    `valid:"stringlength(2|15)"json:"lastname"`
  Accesslevel int       `valid:"required,range(1|2)"json:"access_level"`
  Dateofbirth time.Time `valid:"required"json:"date_of_birth"`
  Password    string    `json:"password,omitempty"`
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

// AfterFind : Remove password from the user to avoid security issues
func (u *User) AfterFind() (err error) {
  u.Password = ""
  return
}

// AfterCreate : Remove password from the user to avoid security issues
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
  u.Password = ""
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
