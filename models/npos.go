package types

import (
	"github.com/jinzhu/gorm"
)

type NPO struct {
	gorm.Model
	NPOName      string `gorm:"type:varchar(255);column:npo_name"`
	Description  string `gorm:"type:varchar(1000);column:description"`
	Website      string `gorm:"type:varchar(500);column:website"`
	Email        string `gorm:"type:varchar(255);unique;not null"`
	FirstName    string `gorm:"type:varchar(255)"`
	LastName     string `gorm:"type:varchar(255)"`
	Password     string `gorm:"type:varchar(255)"`
	Opportunitys []Opportunity
}
