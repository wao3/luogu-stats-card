const { fetchStats, renderSVG } = require("../src/stats-card");
const { renderError } = require("../src/common.js")

module.exports = async (req, res) => {
  const { id } = req.query;

  res.setHeader("Content-Type", "image/svg+xml");
  res.setHeader("Cache-Control", `public, max-age=7200`); // 7200秒（2小时） 缓存

  const validId = /^[1-9]\d*$/;
  if(!validId.test(id)) {
    return res.send(renderError(`"${id}"不是一个合法uid`));
  }

  const stats = await fetchStats(id);
  return res.send(renderSVG(stats));
};
