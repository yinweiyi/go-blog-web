package models

import (
	"gorm.io/gorm"
)

type FriendshipLink struct {
	gorm.Model
	Title       string
	Link        string
	Description string
	IsEnable    int
}
