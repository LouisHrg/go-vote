package model

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

// Survey : the survey struct definition
type Survey struct {
	Abstract 							`sql:"embedded;prefix:-"`
	Title  string    	`json:"title"`
	Desc  string    	`json:"desc"`
	StartDate time.Time    			`json:"startDate"`
	EndDate time.Time    			`json:"endDate"`
	Responses []Response 	`gorm:"PRELOAD:true"`
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

// AfterDelete : Gorm Hook
func (s *Survey) AfterDelete(tx *gorm.DB) (err error) {
 	tx.Model(&Response{}).Where("survey_id = ?", s.ID).Update("deleted_at", time.Now())
  return
}
