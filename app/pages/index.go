package pages

import (
	_ "embed"
	"github.com/arthurvaverko/simple-budget/app/components"
)

//go:embed index.gohtml
var tmpl string

func Index() *components.Page[any] {
	return components.NewPage[any]("/", nil, tmpl)
}
