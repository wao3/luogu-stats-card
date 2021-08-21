const { fetchStats } = require('luogu-stats-card');
const cache = require('./cache');

module.exports = async id => {
  const cacheKey = 'uid:' + id;
  let stats = await cache.get(cacheKey);
  if (!stats) {
    stats = await fetchStats(id);
    if (stats) {
      await cache.put(cacheKey, stats);
    }
  }
  return stats;
}