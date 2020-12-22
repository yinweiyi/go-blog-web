package models

import (
	"math/rand"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name        string
	Slug        string
	Description string
	Order       int

	Articles []*Article `gorm:"many2many:article_tag;"`
}

type Min struct {
	Name string
	Slug string
}

//Shuffle 随机
func Shuffle(mins []Min) []Min {
	values := make([]Min, len(mins))
	buf := make([]Min, len(mins))
	copy(buf, mins)
	for i := range values {
		j := rand.Intn(len(buf))
		values[i] = buf[j]
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return values
}
