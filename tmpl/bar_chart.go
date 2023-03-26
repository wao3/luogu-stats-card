package tmpl

import (
	"github.com/wao3/luogu-stats-card/model"
)

const LabelWidth = 90

var BarChartTmpl = &TmplStruct{
	Name: "BarChartTmpl",
	Tmpl: `
<g transform="translate(90, 0)" style="text-anchor: middle">
	{{range .CoordinateLines}}
	<line x1="{{.LineX}}" y1="0" x2="{{.LineX}}" y2="{{.LineHeight}}" class="line"/>
	<text x="{{.LineX}}" y="{{.TextY}}" class="text">{{.Value}}</text>
	{{end}}
</g>
{{range .Bars}} {{template "BarTmpl" .}} {{end}}
`,
}

type CoordinateLine struct {
	LineX      float64
	LineHeight int
	Value      int
}

func (c *CoordinateLine) TextY() int {
	return c.LineHeight + 10
}

type BarChartBox struct {
	ChartWidth      int
	Bars            []Bar
	CoordinateLines []CoordinateLine
	maxValue        int
}

func NewBarChartBoxFromStats(stats *model.Stats, chartWidth int) *BarChartBox {
	b := &BarChartBox{
		ChartWidth: chartWidth,
	}
	b.InitBarsFromStats(stats)
	return b
}

func NewBarChartBoxFromGuzhi(guzhi *model.Guzhi, chartWidth int, colors model.Colors) *BarChartBox {
	b := &BarChartBox{
		ChartWidth: chartWidth,
	}
	b.InitBarsFromGuzhi(guzhi, colors)
	return b
}

func (b BarChartBox) BoxName() string {
	return BarChartTmpl.Name
}

func (b BarChartBox) BoxHeight() int {
	return b.LineHeight() + 10
}

func (b *BarChartBox) initCoordinateLine() {
	b.CoordinateLines = make([]CoordinateLine, 0, 5)
	var step = b.maxValue / 4
	for i := 0; i < 5; i++ {
		b.CoordinateLines = append(b.CoordinateLines, CoordinateLine{
			LineX:      float64(b.ChartWidth-LabelWidth) / 4 * float64(i),
			LineHeight: b.LineHeight(),
			Value:      step * i,
		})
	}
}

func (b *BarChartBox) InitBarsFromStats(stats *model.Stats) {
	if stats == nil {
		return
	}
	b.Bars = make([]Bar, 0, len(stats.Passed))
	var maxValue int
	for _, p := range stats.Passed {
		if p > maxValue {
			maxValue = p
		}
	}
	b.maxValue = (((maxValue - 1) / 100) + 1) * 100

	for i, level := range model.LevelOrder {
		b.Bars = append(b.Bars, Bar{
			Label:       level.GetLevelName(),
			Value:       stats.Passed[level],
			Color:       level.GetLevelColor(model.ColorDefault),
			MaxValue:    b.maxValue,
			Order:       i,
			ChartWidth:  b.ChartWidth,
			ValueSuffix: "题",
		})
	}
	b.initCoordinateLine()
}

func (b *BarChartBox) InitBarsFromGuzhi(guzhi *model.Guzhi, colors model.Colors) {
	if guzhi == nil {
		return
	}
	b.Bars = make([]Bar, 0, len(guzhi.Guzhi))
	b.maxValue = 100
	for i, guzhiLabel := range model.GuzhiOrder {
		b.Bars = append(b.Bars, Bar{
			Label:       string(guzhiLabel),
			Value:       guzhi.Guzhi[guzhiLabel],
			Color:       model.GetGuzhiColor(guzhi.Guzhi[guzhiLabel], colors),
			MaxValue:    b.maxValue,
			Order:       i,
			ChartWidth:  b.ChartWidth,
			ValueSuffix: "分",
		})
	}
	b.initCoordinateLine()
}

func (b *BarChartBox) LineHeight() int {
	return len(b.Bars) * 30
}
