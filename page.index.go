package main

import (
	"html/template"

	"github.com/yuriizinets/kyoto"
)

type PageIndex struct {
	Rand kyoto.Component
}

func (*PageIndex) Template() *template.Template {
	return template.Must(template.New("page.index.html").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func (p *PageIndex) Init() {
	p.Rand = kyoto.RegC(p, &ComponentRand{})
}
