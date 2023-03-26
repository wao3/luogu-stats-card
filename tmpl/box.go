package tmpl

import (
	"fmt"
)

var BoxTmpl = &TmplStruct{
	Name: "BoxTmpl",
	Tmpl: "",
}

func initBox() {
	addBoxTmpl(NameTitleTmpl.Name, BarChartTmpl.Name)
}

func addBoxTmpl(templateName ...string) {
	for _, name := range templateName {
		BoxTmpl.Tmpl += fmt.Sprintf(`{{if eq .BoxName "%s"}}{{template "%s" .}}{{end}}`, name, name)
	}
}

type BoxType interface {
	BoxHeight() int
	BoxName() string
	BarChartBox
}
