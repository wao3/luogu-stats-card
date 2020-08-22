const axios = require("axios");
const anafanafo = require('anafanafo')
const {
  NAMECOLOR,
  Card,
  renderError,
  renderCCFBadge,
  renderChart,
} = require("./common.js")

/**
 * 
 * @param {number} id 用户id
 * @returns {Object} 获取的用户数据 {name, color, ccfLevel, passed, hideInfo}
 */
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

  const { 
    hideTitle, 
    darkMode,
    cardWidth = 500, 
  } = options || {};

  if(hideInfo) {
    return renderError("用户开启了“完全隐私保护”，获取数据失败");
  }
  
  const paddingX = 25;
  const labelWidth = 90;  //头部文字长度
  const progressWidth = cardWidth - 2*paddingX - labelWidth - 60; //500 - 25*2(padding) - 90(头部文字长度) - 60(预留尾部文字长度)，暂时固定，后序提供自定义选项;
  const nameLength = anafanafo(name)/10*1.8; //计算字体大小为18pt的文本长度

  const datas = [
    {label: "未评定", color:"#bfbfbf", data: passed[0]},
    {label: "入门", color:"#fe4c61", data: passed[1]},
    {label: "普及-", color:"#f39c11", data: passed[2]},
    {label: "普及/提高-", color:"#ffc116", data: passed[3]},
    {label: "普及+/提高", color:"#52c41a", data: passed[4]},
    {label: "提高+/省选-", color:"#3498db", data: passed[5]},
    {label: "省选/NOI-", color:"#9d3dcf", data: passed[6]},
    {label: "NOI/NOI+/CTSC", color:"#0e1d69", data: passed[7]}
  ]
  const charts = renderChart(datas, labelWidth, progressWidth, "题");

  const textColor = darkMode?"#fffefe":"#333333";

  const title = `
  <g transform="translate(0, 0)" font-family="Verdana, Microsoft Yahei" text-rendering="geometricPrecision" font-size="18">
    <text x="0" y="0" fill="${NAMECOLOR[color]}" font-weight="bold" textLength="${nameLength}">
      ${name}
    </text>
    ${ccfLevel < 3 ? "" : renderCCFBadge(ccfLevel, nameLength + 5)}
    <text x="${nameLength + (ccfLevel < 3 ? 10 : 28)}" y="0" fill="${textColor}" font-weight="normal">
      的练习情况
    </text>
  </g>`;

  const body = `<g>${charts}</g>`

  return new Card({
    width: cardWidth - 2*paddingX,
    height: datas.length*30 + 10,
    hideTitle,
    darkMode,
    title,
    body,
  }).render();
}

module.exports = { fetchStats, renderSVG }