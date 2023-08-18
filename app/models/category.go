package models

import (
	"time"

	"gorm.io/gorm"

)


type Category struct {

	ID string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentId string `gorm:"size:36;"`
	Section Section
	SectionId string `gorm:"size:36;index"`
	Products []Product `gorm:"many2many:product_categories;"`
	Name string `gorm:"size:100;"`
	Slug string `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	
}	