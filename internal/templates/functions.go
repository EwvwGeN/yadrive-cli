package templates

import "text/template"

var functions = template.FuncMap{
	"percent": func(a,b int64) float64 {
		return float64(a)/float64(b) * 100
	},
}