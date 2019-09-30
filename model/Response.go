package model

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

// Response : the survey struct definition
type Response struct {
	Abstract 					`sql:"embedded;prefix:-"`
	Message  	string 	`json:"message,omitempty"`
	SurveyID 	int 		`json:"-"json:"omitempty"`
	Survey   	Survey 	`gorm:"PRELOAD:false"json:"-"`
	Users 		[]*User `gorm:"many2many:responses_users;association_foreignkey:uuid"json:"users,omitempty"`
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
