package model

import (
	"github.com/jinzhu/gorm"
)

// IP : the ip struct definition
type IP struct {
	gorm.Model
	Address  string    	`gorm:"primary_key"json:"ip"`
	Attempt  int 				`json:"attempt"`
}

// TableName : Gorm
func (i *IP) TableName() string {
	return "ips"
}
