package components

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed page.gohtml
var baseTemplate string

type AppPage interface {
	Route() string
	Render() (*bytes.Buffer, error)
}

type Page[T any] struct {
	template  string
	route     string
	ViewModel T
}

func NewPage[T any](route string, viewModel T, template string) *Page[T] {
	return &Page[T]{
		template:  template,
		route:     route,
		ViewModel: viewModel,
	}
}

func (p *Page[T]) Render() (*bytes.Buffer, error) {
	t := template.New(p.route)
	baseT, err := t.Parse(baseTemplate)
	if err != nil {
		return nil, err
	}

	pageT, err := baseT.Parse(p.template)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = pageT.ExecuteTemplate(buf, p.route, p.ViewModel)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (p *Page[T]) Route() string {
	return p.route
}
