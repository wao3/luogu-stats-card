const anafanafo = require('anafanafo');

const NAMECOLOR = {
  "Gray": "#bbbbbb",
  "Blue": "#0e90d2",
  "Green": "#5eb95e",
  "Orange": "#e67e22",
  "Red": "#e74c3c",
  "Purple": "#9d3dcf",
  "Cheater": "#ad8b00"
}

class Card {
  constructor({
    width = 450,
    height = 250,
    title = "",
    body = "",
    titleHeight = 25,
    hideTitle = false,
    css = "",
    darkMode = "",
    paddingX = 25,
    paddingY = 35,
    hideBorder = false,
  }) {
    this.width = width;
    this.height = height;
    this.titleHeight = titleHeight;
    this.title = title;
    this.body = body;
    this.hideTitle = hideTitle;
    this.css = css;
    this.darkMode = darkMode;
    this.paddingX = paddingX;
    this.paddingY = paddingY;
    this.hideBorder = hideBorder;
  }

  render() {
    const cardSize = {
      width: this.width + 2*this.paddingX,
      height: this.height + 2*this.paddingY,
    };
    if(!this.hideTitle) cardSize.height += this.titleHeight;

    const bgColor = this.darkMode?"#444444":"#fffefe";
    let borderColor = "";
    if(!this.hideBorder) borderColor = this.darkMode?"#444444":"#E4E2E2";

    return `
      <svg xmlns="http://www.w3.org/2000/svg" width="${cardSize.width}" height="${cardSize.height}" viewBox="0 0 ${cardSize.width} ${cardSize.height}" fill="none">
        <style>
          .text { font: 400 11px 'Segoe UI', Ubuntu, Sans-Serif; fill: ${this.darkMode?"#fffefe":"#333333"} }
          .title {fill: ${this.darkMode?"#fffefe":"#333333"}}
          .line { stroke:${this.darkMode?"#666666":"#dddddd"}; stroke-width:1 }
          ${this.css}
        </style>
        <rect x="0.5" y="0.5" rx="4.5" height="99%" stroke="${borderColor}" width="99%" fill="${bgColor}" stroke-opacity="1" />
        
        ${this.hideTitle ? `` : `
        <g transform="translate(${this.paddingX}, ${this.paddingY})">
          ${this.title}
        </g>`}

        <g transform="translate(${this.paddingX}, ${this.hideTitle ? this.paddingY : this.paddingY + this.titleHeight})">
          ${this.body}
        </g>
      </svg>`;
  }
}

/**
 * 渲染错误卡片
 * @param {string} e 描述错误的文本
 * @param {Object} option 其余选项
 */
const renderError = (e, option) => {
  const css = `.t {font: 600 18px 'Microsoft Yahei UI'; fill: #e74c3c;}`
  const text = `<text class="t" dominant-baseline="text-before-edge">${e}</text>`
  return new Card({
    width: 300,
    height: 23,
    hideTitle: true,
    css,
    body: text,
    paddingY: 20,
    paddingX: 20,
    ...option,
  }).render();
};

/**
 * 渲染 ccf badge
 * @param {number} level CCF等级
 * @param {number} x badge的x坐标
 * @returns {string} ccf badge的svg字符串
 */
const renderCCFBadge = (level, x) => {
  const ccfColor = (ccf) => {
    if(ccf >= 3 && ccf <= 5) return "#5eb95e";
    if(ccf >= 6 && ccf <= 7) return "#3498db";
    if(ccf >= 8) return "#f1c40f";
    return null;
  }
  return `
  <svg xmlns="http://www.w3.org/2000/svg" x="${x}" y="-14" width="18" height="18" viewBox="0 0 18 18" fill="${ccfColor(level)}" style="margin-bottom: -3px;">
    <path d="M16 8C16 6.84375 15.25 5.84375 14.1875 5.4375C14.6562 4.4375 14.4688 3.1875 13.6562 2.34375C12.8125 1.53125 11.5625 1.34375 10.5625 1.8125C10.1562 0.75 9.15625 0 8 0C6.8125 0 5.8125 0.75 5.40625 1.8125C4.40625 1.34375 3.15625 1.53125 2.34375 2.34375C1.5 3.1875 1.3125 4.4375 1.78125 5.4375C0.71875 5.84375 0 6.84375 0 8C0 9.1875 0.71875 10.1875 1.78125 10.5938C1.3125 11.5938 1.5 12.8438 2.34375 13.6562C3.15625 14.5 4.40625 14.6875 5.40625 14.2188C5.8125 15.2812 6.8125 16 8 16C9.15625 16 10.1562 15.2812 10.5625 14.2188C11.5938 14.6875 12.8125 14.5 13.6562 13.6562C14.4688 12.8438 14.6562 11.5938 14.1875 10.5938C15.25 10.1875 16 9.1875 16 8ZM11.4688 6.625L7.375 10.6875C7.21875 10.8438 7 10.8125 6.875 10.6875L4.5 8.3125C4.375 8.1875 4.375 7.96875 4.5 7.8125L5.3125 7C5.46875 6.875 5.6875 6.875 5.8125 7.03125L7.125 8.34375L10.1562 5.34375C10.3125 5.1875 10.5312 5.1875 10.6562 5.34375L11.4688 6.15625C11.5938 6.28125 11.5938 6.5 11.4688 6.625Z">
    </path>
  </svg>`
}

/**
 * 渲染柱状图
 * @param {Object[]} datas 柱状图的数据数组
 * @param {string} datas.label 一条数据的标签
 * @param {string} datas.color 一条数据的颜色
 * @param {number} datas.data 一条数据的数值
 * @param {number} labelWidth 标签宽度
 * @param {number} progressWidth 柱状图的长度
 * @param {string} [unit] 数据单位
 */
const renderChart = (datas, labelWidth, progressWidth, unit) => { //(label, color, height, num, unit) => {
  let chart = "";
  let maxNum = datas.reduce((a, b) => Math.max(a, b.data), 0);
  maxNum = (parseInt((maxNum-1) / 100) + 1) * 100;

  for(let i = 0; i < datas.length; ++i) {
    const width = (datas[i].data+1) / (maxNum+1) * progressWidth;
    chart += `
    <g transform="translate(0, ${i*30})">
      <text x="0" y="15" class="text">${datas[i].label}</text>
      <text x="${width + labelWidth + 10}" y="15" class="text">${datas[i].data + unit}</text>
      <rect height="11" fill="${datas[i].color}" rx="5" ry="5" x="${labelWidth}" y="5" width="${width}"></rect>
    </g>
    `
  }

  const bodyHeight = datas.length * 30 + 10;
  const dw = progressWidth / 4;
  let coordinate = "";
  for(let i = 0; i <= 4; ++i) {
    coordinate += `
    <line x1="${labelWidth + dw*i}" y1="0" x2="${labelWidth + dw*i}" y2="${bodyHeight - 10}"  class="line"/>
    <text x="${labelWidth + dw*i - (i==0?3:5) }" y="${bodyHeight}"  class="text">${maxNum*i/4}</text>
    `;
  }
  return coordinate + chart;
}

/**
 * 
 * @param {string} name 用户名
 * @param {string} color 用户颜色
 * @param {number} ccfLevel 用户ccf等级
 * @param {string} title 标题的后缀
 * @param {string} rightTop 右上角的标签（展示总数）
 */
const renderNameTitle = (name, color, ccfLevel, title, cardWidth, rightTop) => {
  const nameLength = anafanafo(name)/10*1.8;
  const nameColor = NAMECOLOR[color];

  return `
  <g transform="translate(0, 0)" font-family="Verdana, Microsoft Yahei" text-rendering="geometricPrecision" font-size="18">
    <text x="0" y="0" fill="${nameColor}" font-weight="bold" textLength="${nameLength}">
      ${name}
    </text>
    ${ccfLevel < 3 ? "" : renderCCFBadge(ccfLevel, nameLength + 5)}
    <text x="${nameLength + (ccfLevel < 3 ? 10 : 28)}" y="0" class="title" font-weight="normal">
      ${title}
    </text>
    <text x="${cardWidth - 160}" y="0" class="title" font-weight="normal" font-size="13px">
      ${rightTop}
    </text>
  </g>`;
}

module.exports = { 
  NAMECOLOR,
  Card,
  renderError,
  renderCCFBadge,
  renderChart,
  renderNameTitle,
};
