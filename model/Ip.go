package model

import (
	"github.com/jinzhu/gorm"
)

// IP : the ip struct definition
type IP struct {
	gorm.Model
	Address  string    	`json:"ip"`
}

// TableName : Gorm
func (i *IP) TableName() string {
	return "ips"
}
