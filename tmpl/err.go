package tmpl

var ErrTmpl = &TmplStruct{
	Name: "ErrTmpl",
	Tmpl: `
<svg
    xmlns="http://www.w3.org/2000/svg" width="500" height="63" viewBox="0 0 500 63" fill="none">
    <rect x="0.5" y="0.5" rx="4.5" height="99%" stroke="#E4E2E2" width="99%" fill="#fffefe" stroke-opacity="1"/>
    <g transform="translate(250, 20)" style="text-anchor: middle">
        <text class="t" dominant-baseline="text-before-edge" style="font-size: 18px; fill: #e74c3c; font-weight: bold">{{.}}</text>
    </g>
</svg>
`,
}
