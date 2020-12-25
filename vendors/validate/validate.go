package validate

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

var msgMap = map[string]string{
	"email.c_email":    "邮箱格式不正确",
	"content.required": "内容不能为空",
}

//TranslateOverride 翻译错误信息
func TranslateOverride(err error) string {
	var errMap = "表单格式错误"
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			key := fmt.Sprintf("%v.%v", e.Field(), e.Tag())
			if _, ok := msgMap[key]; ok {
				if e.Param() != "" {
					return strings.Replace(msgMap[key], "{"+e.Tag()+"}", e.Param(), -1)
				} else {
					return msgMap[key]
				}
			} else {
				return "表单格式错误:" + key
			}
		}
	}
	return errMap
}

func GetValidator() *validator.Validate {
	validate := validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	RegisterCustomValidation(validate)
	return validate
}

var customValidate = map[string]func(validator.FieldLevel) bool{
	"c_mobile": func(fl validator.FieldLevel) bool {
		ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
		return ok
	},
	"c_email": func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		if email == "" {
			return true
		}
		return validator.New().Var(email, "required,email") == nil
	},
}

//RegisterCustomValidation 注册自定义验证
func RegisterCustomValidation(validate *validator.Validate) {
	for tag, fn := range customValidate {
		_ = validate.RegisterValidation(tag, fn)
	}
}
