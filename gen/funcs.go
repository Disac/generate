package gen

import (
	"text/template"
	"unicode"
)

var fns = template.FuncMap{
	"upper": func(str string) string {
		for i, v := range str {
			return string(unicode.ToUpper(v)) + str[i+1:]
		}
		return ""
	},
}
