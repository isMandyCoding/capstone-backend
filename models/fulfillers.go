package types

import (
	"github.com/jinzhu/gorm"
)

type Fulfiller struct {
	gorm.Model
	OpportunityID uint
	VolunteerID   uint
	Approved      bool   `gorm:"type:boolean"`
	PlannedStart  string `gorm:"type:varchar(255)"`
	PlannedEnd    string `gorm:"type:varchar(255)"`
	ActualStart   string `gorm:"type:varchar(255)"`
	ActualEnd     string `gorm:"type:varchar(255)"`
}
