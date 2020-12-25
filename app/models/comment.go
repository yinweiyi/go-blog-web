package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	ParentId        int
	Content         string
	Avatar          string
	Nickname        string
	Email           string
	CommentableId   int
	CommentableType string
	Ip              string
	IsAudited       int
	IsRead          int
	IsAdminReply    int
	TopId           int

	Children []Comment `gorm:"-"`
}

func (u *Comment) AfterFind(tx *gorm.DB) (err error) {
	u.Avatar = "/assets/" + u.Avatar
	return
}
