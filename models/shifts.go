package types

import (
	"github.com/jinzhu/gorm"
)

type Shift struct {
	gorm.Model
	VolunteerID     uint
	EventID         uint
	WasWorked       bool `gorm:"type:boolean"`
	ActualStartTime int64
	ActualEndTime   int64
}
