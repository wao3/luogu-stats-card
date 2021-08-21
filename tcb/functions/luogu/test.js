const { main } = require('./index');

main({
  queryStringParameters: {
    id: 313209,
  },
  path: '/practice',
}).then(res => console.log(res));