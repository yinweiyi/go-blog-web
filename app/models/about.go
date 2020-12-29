package models

type About struct {
	Model

	Title       string
	ContentType int
	Markdown    string
	Html        string
	Order       int
	IsEnable    int
}
