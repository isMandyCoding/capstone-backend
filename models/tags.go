package types

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	TagName string
	Events  []*Event `gorm:"many2many:event_tags;"`
}
