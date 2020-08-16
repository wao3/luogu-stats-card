const { renderSVG } = require("../src/stats-card");

module.exports = async (req, res) => {
  const { id } = req.query;

  res.setHeader("Content-Type", "image/svg+xml");
  res.setHeader("Cache-Control", `public, max-age=7200`); // 7200秒（2小时） 缓存

  const stats = await fetchStats(id);
  return renderSVG(stats);
};
