package model

import (
	"github.com/jinzhu/gorm"
)

// Response : the survey struct definition
type Response struct {
	gorm.Model
	Message  string `json:"message"`
	SurveyID int
	Survey   Survey `gorm:"PRELOAD:false"`
}

// TableName : Gorm related
func (s *Response) TableName() string {
	return "responses"
}
