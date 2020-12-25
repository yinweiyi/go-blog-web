package forms

type Comment struct {
	Type      string `form:"type"`
	ID        int    `form:"id"`
	Email     string `form:"email" validate:"c_email"`
	ParentId  int    `form:"parent_id"`
	Nickname  string `form:"nickname"`
	Content   string `form:"content" validate:"required"`
	Captcha   string `form:"captcha" validate:"required"`
	CaptchaId string `form:"captcha_id" validate:"required"`
	IP        string
}
