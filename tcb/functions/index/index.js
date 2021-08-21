const axios = require('axios');

module.exports.main = async function (event, context) {
  const baseUrl = "https://www.wao3.cn";
  if (event.path === "/") event.path = "/luogu/index.html"
  const res = await axios.get(baseUrl + event.path);
  return {
    statusCode: res.status,
    headers: res.headers,
    body: res.data,
  }
};