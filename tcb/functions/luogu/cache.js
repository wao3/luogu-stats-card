const cloudbase = require('@cloudbase/node-sdk')
const app = cloudbase.init({
  env: 'luogu-7gv9ryjh2760929a'
})
const db = app.database();

// 12 hour
const EXPIRE_TIME_MILLISECOND = 1000 * 60 * 60 * 12;
const CACHE_NAME = 'cache';

const cache = {
  get: async (key) => {
    const res = await db.collection(CACHE_NAME)
      .doc(key)
      .get();
    if (res.data === null || res.data.length === 0) {
      return null;
    }
    const data = res.data[0];
    if (!data || !data.updateTime) {
      return null;
    }
    const updateTime = new Date(data.updateTime);
    const now = new Date();
    if (now.valueOf() - updateTime.valueOf() > EXPIRE_TIME_MILLISECOND) {
      return null;
    }
    return data.value;
  },
  put: async (key, value) => {
    if (key == null) {
      return null;
    }
    const res = await db
      .collection(CACHE_NAME)
      .doc(key)
      .set({
        value,
        updateTime: new Date(),
      });
    if (!res.updated || !res.upsertedId) {
      return value;
    }
    return null;
  }
}

module.exports = cache;