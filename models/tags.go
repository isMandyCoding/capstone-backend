package types

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	EventID uint
	TagName string
}
