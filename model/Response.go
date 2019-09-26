package model

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

// Response : the survey struct definition
type Response struct {
	Abstract 					`sql:"embedded;prefix:-"`
	Message  	string 	`json:"message"`
	SurveyID 	int 		`json:"-"`
	Survey   	Survey 	`gorm:"PRELOAD:false"json:"-"`
	Users 		[]*User `gorm:"many2many:respones_users;"`
}

// TableName : Gorm related
func (r *Response) TableName() string {
	return "responses"
}

// BeforeCreate : Gorm hook
func (r *Response) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return
}

// AfterFind : Gorm hook
func (r *Response) AfterFind() (err error) {
	r.ID = 0
	return
}
