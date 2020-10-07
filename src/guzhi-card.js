const {
  Card,
  renderError,
  renderChart,
  renderNameTitle,
} = require("./common.js");

const renderGuzhiCard = (userInfo, scores, options) => {
  const regNum = /^\d*$/;
  if(!scores || typeof scores !== 'string') {
    return renderError('咕值信息不能为空', {width: 400});
  }
  let sp = ',';
  if(scores.indexOf('，') >= 0) {
    sp = '，';
  }
  const scoreArray = scores.split(sp).filter(x => regNum.test(x)).map(x => parseInt(x)).filter(x => x >= 0 && x <= 100);
  if(scoreArray.length != 5) {
    return renderError(`咕值信息"${scores}"不合法`, {width: 400});
  }
  const scoreSum = scoreArray.reduce((a, b) => a+b);

  const {
    name,
    color,
    ccfLevel,
  } = userInfo || {};

  const { 
    hideTitle, 
    darkMode,
    cardWidth = 500, 
  } = options || {};

  const paddingX = 25;
  const labelWidth = 90;  //柱状图头部文字长度
  const progressWidth = cardWidth - 2*paddingX - labelWidth - 60; //500 - 25*2(padding) - 90(头部文字长度) - 60(预留尾部文字长度)，暂时固定，后序提供自定义选项;

  const getScoreColor = (score) => {
    if (score >= 80) return "#52c41a";
    if (score >= 60) return "#fadb14";
    if (score >= 30) return "#f39c11";
    return "#e74c3c";
  }
  const datas = [
    {label: "基础信用", data: scoreArray[0], color: getScoreColor(scoreArray[0])},
    {label: "练习情况", data: scoreArray[1], color: getScoreColor(scoreArray[1])},
    {label: "社区贡献", data: scoreArray[2], color: getScoreColor(scoreArray[2])},
    {label: "比赛情况", data: scoreArray[3], color: getScoreColor(scoreArray[3])},
    {label: "获得成就", data: scoreArray[4], color: getScoreColor(scoreArray[4])},
  ]

  const title = userInfo != null ? renderNameTitle(name, color, ccfLevel, "的咕值信息", cardWidth, `总咕值: ${scoreSum}分`) : "";

  const body = renderChart(datas, labelWidth, progressWidth, "分");

  return new Card({
    width: cardWidth - 2*paddingX,
    height: datas.length*30 + 10,
    hideTitle,
    darkMode,
    title,
    body,
  }).render();
}

module.exports = { renderGuzhiCard }