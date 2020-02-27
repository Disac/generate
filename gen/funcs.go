package gen

import (
	"strings"
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
	"join": func(slice []string) string {
		newSlice := []string{}
		for _, str := range slice {
			newSlice = append(newSlice, "\""+str+"\"")
		}
		return strings.Join(newSlice, ",")
	},
}
