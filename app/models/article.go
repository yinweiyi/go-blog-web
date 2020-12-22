package models

import (
	"blog/vendors/helpers"

	strip "github.com/grokify/html-strip-tags-go"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string
	Slug        string
	Author      string
	ContentType uint
	Markdown    string `gorm:"->:false;<-:create"`
	Html        string
	Description string
	Keywords    string
	IsTop       int
	IsShow      int
	Views       int
	Order       int
	CategoryId  int

	Category Category `gorm:"foreignKey:CategoryId"`
	Tags     []*Tag   `gorm:"many2many:article_tag;"`
}

func (article *Article) ShortHtml() string {
	striped := strip.StripTags(article.Html)
	return helpers.Substr(striped, 0, 200, "")
}

func (article *Article) ShortTitle() string {
	return helpers.Substr(article.Title, 0, 30, "...")
}
