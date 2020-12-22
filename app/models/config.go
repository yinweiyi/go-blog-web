package models

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model

	Title       string
	SubTitle    string
	Keywords    string
	Icp         string
	Author      string
	Description string
}
