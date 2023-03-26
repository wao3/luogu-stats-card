package tmpl

import (
	"fmt"

	model2 "github.com/wao3/luogu-stats-card/model"
)

var NameTitleTmpl = &TmplStruct{
	Name: "NameTitleTmpl",
	Tmpl: `
<g transform="translate(25, 35)" class="title" text-rendering="geometricPrecision">
	<text>
		<tspan fill="{{.NameColor}}" font-weight="bold">{{.Name}}</tspan>
		{{if .HasCcf}}
		<tspan fill="{{.CcfColor}}" font-size="16" font-family="iconfont">&#xe607;</tspan>
		{{end}}
		<tspan>{{.TitleSuffix}}</tspan>
	</text>
	<text x="{{.SubTitleX}}" y="0" font-size="13">{{.SubTitle}}</text>
</g>
`,
}

type NameTitle struct {
	Name        string
	TitleSuffix string
	SubTitle    string
	CcfColor    *string
	NameColor   string
	cardWidth   int
	colors      model2.Colors
}

func NewNameTitleBoxFromStats(stats *model2.Stats, cardWidth int, colors model2.Colors) *NameTitle {
	if stats == nil {
		return nil
	}
	var sum int
	for _, p := range stats.Passed {
		sum += p
	}
	return &NameTitle{
		Name:        stats.User.Name,
		TitleSuffix: "的练习情况",
		SubTitle:    fmt.Sprintf("已通过：%d题", sum),
		NameColor:   stats.User.GetColor(colors),
		CcfColor:    stats.User.GetCcfColor(colors),
		cardWidth:   cardWidth,
		colors:      colors,
	}
}

func NewNameTitleBoxFromGuzhi(guzhi *model2.Guzhi, cardWidth int, colors model2.Colors) *NameTitle {
	if guzhi == nil {
		return nil
	}
	var sum int
	for _, g := range guzhi.Guzhi {
		sum += g
	}
	return &NameTitle{
		Name:        guzhi.User.Name,
		TitleSuffix: "的咕值信息",
		SubTitle:    fmt.Sprintf("总咕值：%d分", sum),
		NameColor:   guzhi.User.GetColor(colors),
		CcfColor:    guzhi.User.GetCcfColor(colors),
		cardWidth:   cardWidth,
		colors:      colors,
	}
}

func (n *NameTitle) HasCcf() bool {
	return n.CcfColor != nil
}

func (n *NameTitle) SubTitleX() int {
	return n.cardWidth - 50
}
