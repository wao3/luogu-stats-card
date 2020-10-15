const { renderGuzhiCard } = require("../src/guzhi-card");
const { fetchStats } = require("../src/stats-card");
const { renderError } = require("../src/common.js")

module.exports = async (req, res) => {
  const { id, scores, hide_title, dark_mode, card_width = 500 } = req.query;

  res.setHeader("Content-Type", "image/svg+xml");
  res.setHeader("Cache-Control", "public, max-age=43200"); // 43200s（12h） cache

  const regNum = /^[1-9]\d*$/;
  const clamp = (min, max, n) => Math.max(min, Math.min(max, n));

  if (!regNum.test(card_width)) {
    return res.send(
      renderError(`卡片宽度"${card_width}"不合法`, { darkMode: dark_mode })
    );
  }
  if(id != undefined && !regNum.test(id)) {
    return res.send(renderError(`"${id}"不是一个合法uid`, {darkMode: dark_mode}));
  }

  let stats = null;

  if(id != undefined) {
    stats = await fetchStats(id);
  }

  return res.send(
    renderGuzhiCard(stats, scores, {
      hideTitle: stats === null ? true : hide_title,
      darkMode: dark_mode,
      cardWidth: clamp(500, 1920, card_width),
    })
  );
};
