package models

import "github.com/jinzhu/gorm"

// TagModel godoc
type TagModel struct {
	gorm.Model
	name string `gorm:not null`
}