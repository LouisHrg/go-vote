package model

import (
  "time"

  age "github.com/bearbin/go-age"
  "github.com/jinzhu/gorm"
  "github.com/qor/validations"
  "github.com/satori/go.uuid"
  "golang.org/x/crypto/bcrypt"
)

// User : the user struct definition
type User struct {
  Abstract              `sql:"embedded;prefix:-"`
  Email       string    `valid:"email"gorm:"type:varchar(100);unique_index"json:"email,omitempty"`
  Firstname   string    `valid:"stringlength(2|15),alphanum"json:"firstname,omitempty"`
  Lastname    string    `valid:"stringlength(2|15),alphanum"json:"lastname,omitempty"`
  Accesslevel int       `valid:"required,range(1|2)"json:"access_level,omitempty"`
  Dateofbirth time.Time `valid:"required"json:"date_of_birth,omitempty"`
  Password    string    `json:"password,omitempty"`
}

// TableName : Gorm related
func (u *User) TableName() string {
  return "users"
}

// Validate : custom validation rules
func (u User) Validate(db *gorm.DB) {
  if age.Age(u.Dateofbirth) < 18 {
    db.AddError(validations.NewError(u, "Age", "User needs to be 18+"))
  }

  r := User{}
  _ = db.Where("email = ?", u.Email).First(&r)

  if r.Email == u.Email {
    db.AddError(validations.NewError(u, "Email", "User with this email already exist"))
  }
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
