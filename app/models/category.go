package models

type Category struct {
	Model
	Name        string
	Slug        string
	Description string
	Order       int
}
