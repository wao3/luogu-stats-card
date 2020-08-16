const axios = require("axios");
const { renderError } = require("./common.js")

async function fetchStats(id) {
  const res = await axios.get(`https://www.luogu.com.cn/user/${id}?_contentOnly`)

  const stats = {
    name: "NULL",
    color: "Gray",
    ccfLevel: 0,
    passed: [0,0,0,0,0,0,0,0],
    hideInfo: false
  }
  if(res.data.code !== 200) {
    return stats;
  }

  const user = res.data.currentData.user;
  const passed = res.data.currentData.passedProblems;

  stats.name = user.name;
  stats.color = user.color;
  stats.ccfLevel = user.ccfLevel;

  if(!passed) {
    stats.hideInfo = true;
    return stats;
  }

  for(let i of passed) {
    stats.passed[i.difficulty]++;
  }

  return stats;
} 

const renderSVG = (stats, options) => {
  const {
    name,
    color,
    ccfLevel,
    passed,
    hideInfo
  } = stats;
  const { hideTitle } = options || {};

  if(hideInfo) {
    return renderError("用户开启了“完全隐私保护”，获取数据失败");
  }

  let max = 0;
  for(let i of passed) {
    if(max < i) {
      max = i;
    }
  }
  max = (parseInt(max / 100) + 1) * 100;

  const DIFFICULTY = [
    {label: "未评定", color:"#bfbfbf"},
    {label: "入门", color:"#fe4c61"},
    {label: "普及-", color:"#f39c11"},
    {label: "普及/提高-", color:"#ffc116"},
    {label: "普及+/提高", color:"#52c41a"},
    {label: "提高+/省选-", color:"#3498db"},
    {label: "省选/NOI-", color:"#9d3dcf"},
    {label: "NOI/NOI+/CTSC", color:"#0e1d69"}
  ]

  const NAMECOLOR = {
    "Gray": "#bbbbbb",
    "Blue": "#0e90d2",
    "Green": "#5eb95e",
    "Orange": "#e67e22",
    "Red": "#e74c3c",
    "Purple": "#9d3dcf"
  }

  const renderItem = (label, color, num, height) => {
    const width = (num+1) / (max+1) * 300;
    return `
    <g transform="translate(0, ${height})">
      <text data-testid="s-text" x="2" y="15" class="s-text">${label}</text>
      <text x="${width + 100}" y="15" class="s-text">${num}题</text>
      <rect height="11" fill="${color}" rx="5" ry="5" x="90" y="5" width="${width}"></rect>
    </g>
    `
  }

  let items = "";
  for(let i = 0; i < 8; ++i) {
    items += renderItem(DIFFICULTY[i].label, DIFFICULTY[i].color, passed[i], i*30);
  }

  return `
  <svg xmlns="http://www.w3.org/2000/svg" width="500" height="330" viewBox="0 0 500 330" fill="none">
    <style>
      .header { font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif; fill: ${NAMECOLOR[color]};} 
      .s-text { font: 400 11px 'Segoe UI', Ubuntu, Sans-Serif; fill: #333333;}
    </style>
    <rect data-testid="card-bg" x="0.5" y="0.5" rx="4.5" height="99%" stroke="#E4E2E2" width="99%" fill="#fffefe" stroke-opacity="1" />
    <g data-testid="card-title" transform="translate(25, ${hideTitle ? 0 : 35})">
      <g transform="translate(0, 0)">
        <text x="0" y="0" class="header">${name} <tspan fill="#333333" style="font-weight: 500;"> 的练习情况</tspan></text></g>
    </g>
    <g data-testid="main-card-body" transform="translate(0, ${hideTitle ? 20 : 55})">
      <svg data-testid="lang-items" x="25">

        <line x1="90"  y1="0" x2="90" y2="240" style="stroke:#ccc; stroke-width:1"/>
        <line x1="165" y1="0" x2="165" y2="240" style="stroke:#ccc; stroke-width:1"/>
        <line x1="240" y1="0" x2="240" y2="240" style="stroke:#ccc; stroke-width:1"/>
        <line x1="315" y1="0" x2="315" y2="240" style="stroke:#ccc; stroke-width:1"/>
        <line x1="390" y1="0" x2="390" y2="240" style="stroke:#ccc; stroke-width:1"/>

        <text x="87" y="250" class="s-text">0</text>
        <text x="160" y="250" class="s-text">${max*0.25}</text>
        <text x="235" y="250" class="s-text">${max*0.5}</text>
        <text x="310" y="250" class="s-text">${max*0.75}</text>
        <text x="385" y="250" class="s-text">${max}</text>

        ${items}
      </svg>
    </g>
  </svg>
  `

}

module.exports = { fetchStats, renderSVG }