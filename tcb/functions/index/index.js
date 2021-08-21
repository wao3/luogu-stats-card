const axios = require('axios');

module.exports.main = async function (event, context) {
  const baseUrl = "https://wao3.cn";
  const res = await axios.get(baseUrl + event.path);
  return {
    statusCode: res.status,
    headers: res.headers,
    body: res.data,
  }
};