package tmpl

import "fmt"

var BarTmpl = &TmplStruct{
	Name: "BarTmpl",
	Tmpl: `
<g transform="translate(0, {{.Y}})">
	<text x="0" y="15" class="text">{{.Label}}</text>
	<text x="{{.ValueStrX}}" y="15" class="text">{{.ValueStr}}</text>
	<rect height="11" fill="{{.Color}}" rx="5" ry="5" x="90" y="5" width="{{.ValueWidth}}"/>
</g>
`,
}

type Bar struct {
	Label       string
	Value       int
	Color       string
	MaxValue    int
	Order       int
	ChartWidth  int
	ValueSuffix string
}

func NewBar(label string, value int, color string, maxValue int, order int, chartWidth int, valueSuffix string) *Bar {
	return &Bar{
		Label:       label,
		Value:       value,
		Color:       color,
		MaxValue:    maxValue,
		Order:       order,
		ChartWidth:  chartWidth,
		ValueSuffix: valueSuffix,
	}
}

func (h *Bar) Y() int {
	return h.Order * 30
}

func (h *Bar) ValueWidth() float64 {
	// Value+minValue 防止柱状图宽度为0
	minValue := h.MaxValue / 100
	return float64(h.ChartWidth-LabelWidth) * (float64(h.Value+minValue) / float64(h.MaxValue+minValue))
}

func (h *Bar) ValueStr() string {
	return fmt.Sprintf("%d%s", h.Value, h.ValueSuffix)
}

func (h *Bar) ValueStrX() float64 {
	return h.ValueWidth() + 100
}
