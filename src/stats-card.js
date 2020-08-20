const axios = require("axios");
const anafanafo = require('anafanafo')
const { Card, renderError } = require("./common.js")

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
    hideInfo,
  } = stats;
  const { hideTitle, darkMode } = options || {};
  
  const bodyHeight = 250;
  const cardWidth = 500;  //卡片宽度
  const labelWidth = 90;  //头部文字长度
  const progressWidth = 300; //500 - 25*2(padding) - 90(头部文字长度) - 60(预留尾部文字长度)，暂时固定，后序提供自定义选项;
  const nameLength = anafanafo(name)/10*1.8; //计算字体大小为18pt的文本长度

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

  const ccfColor = (ccf) => {
    if(ccf >= 3 && ccf <= 5) return "#5eb95e";
    if(ccf >= 6 && ccf <= 7) return "#3498db";
    if(ccf >= 8) return "#f1c40f";
    return null;
  }

  const ccfBadge = `
    <svg xmlns="http://www.w3.org/2000/svg" x="${nameLength + 5}" y="-14" width="18" height="18" viewBox="0 0 18 18" fill="${ccfColor(ccfLevel)}" style="margin-bottom: -3px;">
      <path d="M16 8C16 6.84375 15.25 5.84375 14.1875 5.4375C14.6562 4.4375 14.4688 3.1875 13.6562 2.34375C12.8125 1.53125 11.5625 1.34375 10.5625 1.8125C10.1562 0.75 9.15625 0 8 0C6.8125 0 5.8125 0.75 5.40625 1.8125C4.40625 1.34375 3.15625 1.53125 2.34375 2.34375C1.5 3.1875 1.3125 4.4375 1.78125 5.4375C0.71875 5.84375 0 6.84375 0 8C0 9.1875 0.71875 10.1875 1.78125 10.5938C1.3125 11.5938 1.5 12.8438 2.34375 13.6562C3.15625 14.5 4.40625 14.6875 5.40625 14.2188C5.8125 15.2812 6.8125 16 8 16C9.15625 16 10.1562 15.2812 10.5625 14.2188C11.5938 14.6875 12.8125 14.5 13.6562 13.6562C14.4688 12.8438 14.6562 11.5938 14.1875 10.5938C15.25 10.1875 16 9.1875 16 8ZM11.4688 6.625L7.375 10.6875C7.21875 10.8438 7 10.8125 6.875 10.6875L4.5 8.3125C4.375 8.1875 4.375 7.96875 4.5 7.8125L5.3125 7C5.46875 6.875 5.6875 6.875 5.8125 7.03125L7.125 8.34375L10.1562 5.34375C10.3125 5.1875 10.5312 5.1875 10.6562 5.34375L11.4688 6.15625C11.5938 6.28125 11.5938 6.5 11.4688 6.625Z">
      </path>
    </svg>`


  const NAMECOLOR = {
    "Gray": "#bbbbbb",
    "Blue": "#0e90d2",
    "Green": "#5eb95e",
    "Orange": "#e67e22",
    "Red": "#e74c3c",
    "Purple": "#9d3dcf",
    "Cheater": "#ad8b00"
  }

  const renderCoordinate = () => {// (maxNum, labelWidth, progressWidth, bodyHeight) => {
    const dw = progressWidth / 4;
    let res = "";
    for(let i = 0; i <= 4; ++i) {
      res += `
      <line x1="${labelWidth + dw*i}" y1="0" x2="${labelWidth + dw*i}" y2="${bodyHeight - 10}"  class="line"/>
      <text x="${labelWidth + dw*i - (i==0?3:5) }" y="${bodyHeight}"  class="text">${max*i/4}</text>
      `;
    }
    return res;
  }

  const renderLine = () => { //(label, color, height, num, unit) => {
    let items = "";
    for(let i = 0; i < 8; ++i) {
      const width = (passed[i]+1) / (max+1) * progressWidth;
      items += `
      <g transform="translate(0, ${i*30})">
        <text x="2" y="15" class="text">${DIFFICULTY[i].label}</text>
        <text x="${width + 100}" y="15" class="text">${passed[i]}题</text>
        <rect height="11" fill="${DIFFICULTY[i].color}" rx="5" ry="5" x="${labelWidth}" y="5" width="${width}"></rect>
      </g>
      `
    }
    return items;
  }

  const textColor = darkMode?"#fffefe":"#333333";

  const title = `
  <g transform="translate(0, 0)" font-family="Verdana, Microsoft Yahei" text-rendering="geometricPrecision" font-size="18">
    <text x="0" y="0" fill="${NAMECOLOR[color]}" font-weight="bold" textLength="${nameLength}">
      ${name}
    </text>
    ${ccfLevel < 3 ? "" : ccfBadge}
    <text x="${nameLength + (ccfLevel < 3 ? 10 : 28)}" y="0" fill="${textColor}" font-weight="normal">
      的练习情况
    </text>
  </g>`;

  const body = `
  <g>
    ${renderCoordinate()}
    ${renderLine()}
  </g>
  `

  return new Card({
    hideTitle,
    darkMode,
    title,
    body,
  }).render();
}

module.exports = { fetchStats, renderSVG }