package types

import (
	"github.com/jinzhu/gorm"
)

type Volunteer struct {
	gorm.Model
	Username  string `gorm:"type:varchar(255)"`
	Bio       string `gorm:type:varchar(1000)`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(255)"`
}
