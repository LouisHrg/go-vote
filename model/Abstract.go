package model

import (
	"time"
)

// Abstract : gathering all the common column for model definition
type Abstract struct {
	UUID      string 			`gorm:"unique_index"json:"uuid"`
	ID        uint 				`gorm:"primary_key"json:"-"`
	CreatedAt time.Time 	`json:"createdAt"`
	UpdatedAt time.Time 	`json:"updatedAt"`
	DeletedAt *time.Time 	`sql:"index"json:"-"`
}

