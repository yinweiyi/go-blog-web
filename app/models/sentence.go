package models

import (
	"gorm.io/gorm"
)

type Sentence struct {
	gorm.Model
	Author      string
	Content     string
	Translation string
}
