const axios = require('axios');
const { renderError } = require('luogu-stats-card');
const practice = require('./route/practice');
const guzhi = require('./route/guzhi');

module.exports.main = async function (event, context) {
  checkParam(event.queryStringParameters);
  let result = null;
  if (event.path.startsWith("/practice")) {
    result = await practice(event);
  } else if (event.path.startsWith("/guzhi")) {
    result = await guzhi(event);
  } else {
    result = renderError(`路径错误：${event.path}`, { darkMode: dark_mode });
  }
  return {
    statusCode: 200,
    headers: {
      "content-type": "image/svg+xml; charset=utf-8",
      "Cache-Control": "public, max-age=43200",
    },
    body: result
  };
};

function checkParam(queryParam) {
  const { id, dark_mode, card_width = 500 } = queryParam;

  const regNum = /^[1-9]\d*$/;
  const clamp = (min, max, n) => Math.max(min, Math.min(max, n));

  if (!regNum.test(card_width)) {
    return renderError(`卡片宽度"${card_width}"不合法`, { darkMode: dark_mode });
  }

  if(id != undefined && !regNum.test(id)) {
    return renderError(`卡片宽度"${card_width}"不合法`, { darkMode: dark_mode });
  }

  queryParam.card_width = clamp(500, 1920, card_width);
}