package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Survey : the survey struct definition
type Survey struct {
	gorm.Model
	UUID      string     `gorm:"unique_index"`
	Question  string     `json:"question"`
	Responses []Response `gorm:"PRELOAD:true"`
}

// TableName : Gorm
func (s *Survey) TableName() string {
	return "surveys"
}

// BeforeCreate : Gorm hook
func (s *Survey) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return
}

//AfterFind : Gorm hook
func (s *Survey) AfterFind() (err error) {
	s.ID = 0
	return
}
