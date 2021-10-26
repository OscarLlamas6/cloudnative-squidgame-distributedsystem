require('dotenv').config()
const app = require('./app');

app.listen(app.get('port'));

console.log('Server on port', app.get('port'))


// const schedule = require('node-schedule');

// const job = schedule.scheduleJob('10 * * * * *', function(){
//   console.log('The answer to life, the universe, and everything!');
// });
