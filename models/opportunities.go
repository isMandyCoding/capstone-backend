package types

import (
	"github.com/jinzhu/gorm"
)

type Opportunity struct {
	gorm.Model
	NPOID           uint
	StartTime       string `gorm:"type:varchar(255)"`
	EndTime         string `gorm:"type:varchar(255)"`
	Label           string `gorm:"type:varchar(255)"`
	Tags            string `gorm:"type:varchar(255)"`
	Description     string `gorm:"type:varchar(1000)"`
	Location        string `gorm:"type:varchar(255)"`
	NumOfVolunteers int    `gorm:"type:integer"`
	Fulfillers      []Fulfiller
}
