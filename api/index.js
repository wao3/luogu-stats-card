const { fetchStats, renderSVG } = require("../src/stats-card");
const { renderError } = require("../src/common.js")

module.exports = async (req, res) => {
  const { id, hide_title, dark_mode } = req.query;

  res.setHeader("Content-Type", "image/svg+xml");
  res.setHeader("Cache-Control", `public, max-age=43200`); // 43200s（12h） cache

  const validId = /^[1-9]\d*$/;
  if(!validId.test(id)) {
    return res.send(renderError(`"${id}"不是一个合法uid`));
  }

  const stats = await fetchStats(id);
  return res.send(renderSVG(stats, {
    hideTitle: hide_title,
    darkMode: dark_mode,
  }));
};
