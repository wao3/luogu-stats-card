const { renderGuzhiCard } = require('luogu-stats-card');
const fetchStats = require('../fetchStats');

module.exports = async function (event) {
  const {
    id,
    scores,
    hide_title,
    dark_mode,
    card_width = 500,
  } = event.queryStringParameters;

  const stats = await fetchStats(id);

  return renderGuzhiCard(stats, scores, {
    hideTitle: stats === null ? true : hide_title,
    darkMode: dark_mode,
    cardWidth: card_width,
  });
};