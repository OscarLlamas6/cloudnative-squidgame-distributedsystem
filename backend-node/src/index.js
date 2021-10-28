require('dotenv').config()
const app = require('./app');
const redis = require('./redisDb')

app.listen(app.get('port'));

console.log('Server on port', app.get('port'))