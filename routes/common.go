package routes

import (
	"blog/vendors/helpers"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	strip "github.com/grokify/html-strip-tags-go"
	"html/template"
	"path/filepath"
	"strings"
)

func RegisterCommonFile(engine *gin.Engine) {
	//加载静态文件
	engine.Static("/assets", "./public")
	//加载公共模板
	engine.HTMLRender = loadTemplates("./resources/views")

}

//加载公共模板
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/*[^layouts]/*.html")
	if err != nil {
		panic(err.Error())
	}
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(fileName(include), template.FuncMap{
			"randomColor":  helpers.RandomColor,
			"randomInt":    helpers.RandomInt,
			"formatAsDate": helpers.FormatAsDate,
			"stripTags":    strip.StripTags,
		}, files...)
	}
	return r
}

//获取文件名
func fileName(path string) string {
	pathArr := strings.Split(path, "\\")
	pathArr = pathArr[len(pathArr)-2:]
	return strings.Join(pathArr, "/")
}
