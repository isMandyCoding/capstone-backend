package types

import (
	"github.com/jinzhu/gorm"
)

type Opportunity struct {
	gorm.Model
	StartTime       string `gorm:"type:varchar(255);column:start_time"`
	EndTime         string `gorm:"type:varchar(255);column:end_time"`
	Label           string `gorm:"type:varchar(255);column:label"`
	Tags            string `gorm:"type:varchar(255);column:tags"`
	Description     string `gorm:"type:varchar(1000);column:label"`
	Location        string `gorm:"type:varchar(255);column:location"`
	NumOfVolunteers int    `gorm:"type:integer;column:num_of_volunteers"`
}
