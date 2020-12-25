package config

import (
	"blog/vendors/config"
	"time"
)

func init() {
	config.Add("captcha", config.StrMap{
		"length": 3,
		"width":  100,
		"height": 32,
		"expire": 10 * time.Minute,
	})
}
