package models

type Guestbook struct {
	Model

	ContentType int
	Markdown    string
	Html        string
	CanComment  int
}

func (g Guestbook) TableName() string {
	return "guestbook"
}
