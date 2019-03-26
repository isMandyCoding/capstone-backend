package types

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	NPOID           uint
	Name            string `gorm:"type:varchar(255)"`
	StartTime       int64
	EndTime         int64
	Tags            []*Tag `gorm:"many2many:user_languages;"`
	Description     string `gorm:"type:varchar(1000)"`
	Location        string `gorm:"type:varchar(255)"`
	NumOfVolunteers int    `gorm:"type:integer"`
	Shifts          []Shift
}
