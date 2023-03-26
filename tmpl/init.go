package tmpl

import (
	"text/template"

	"github.com/wao3/luogu-stats-card/common"
)

type TmplStruct struct {
	Tmpl string
	Name string
}

var Tmpl, _ = template.New("ROOT").Parse("")

func init() {
	initBox()
	registerTmpls(CardTmpl, BoxTmpl, NameTitleTmpl, BarTmpl, BarChartTmpl, ErrTmpl)
}

func registerTmpls(tmpl ...*TmplStruct) {
	for _, t := range tmpl {
		registerTmpl(t)
	}
}

func registerTmpl(tmpl *TmplStruct) {
	var err error
	Tmpl, err = Tmpl.New(tmpl.Name).Parse(tmpl.Tmpl)
	common.FastFail(err)
}
