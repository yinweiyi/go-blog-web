package helpers

import (
	"fmt"
	"html/template"
	"math/rand"
	"strings"
	"time"
)

//随机颜色
func RandomColor() string {
	colors := make([]string, 6)

	for i := 0; i < 6; i++ {
		colors = append(colors, fmt.Sprintf("%x", rand.Intn(16)))
	}
	return "#" + strings.Join(colors, "")
}

//随机数
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

//格式化时间
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

//格式化时间
func FormatAsDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//strip tag
func Substr(str string, start, length int, suffix string) string {
	if length == 0 {
		return ""
	}
	runeStr := []rune(str)
	lenStr := len(runeStr)

	if start < 0 {
		start = lenStr + start
	}
	if start > lenStr {
		start = lenStr
	}
	end := start + length
	if end > lenStr {
		suffix = ""
		end = lenStr
	}
	if length < 0 {
		end = lenStr + length
	}
	if start > end {
		start, end = end, start
	}
	return string(runeStr[start:end]) + suffix
}

func Chuck(size int) [][]int {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	var result [][]int
	for {
		if len(arr) <= size {
			result = append(result, arr)
			break
		}
		result = append(result, arr[:size])
		arr = arr[size:]
	}
	return result
}

//Unescaped
func Unescaped(x string) interface{} {
	return template.HTML(x)
}

func ArrayContain(arr []int, item int) bool {
	for _, d := range arr {
		if d == item {
			return true
		}
	}
	return false
}
