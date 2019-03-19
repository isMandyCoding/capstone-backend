package types

import (
	"github.com/jinzhu/gorm"
)

type NPO struct {
	gorm.Model
	NPOName      string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:varchar(1000)"`
	Website      string `gorm:"type:varchar(500)"`
	Email        string `gorm:"type:varchar(255);unique;not null"`
	FirstName    string `gorm:"type:varchar(255)"`
	LastName     string `gorm:"type:varchar(255)"`
	Password     string `gorm:"type:varchar(255)"`
	Opportunitys []Opportunity
}
