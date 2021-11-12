const io = require('./socket')
const redisData = require('./redisDb')
const mongoData = require('./mongoDb')

const redis = require("redis");
const client = redis.createClient({
    host: process.env.REDIS_HOST,
    port: process.env.REDIS_PORT,
    password: process.env.REDIS_PASS,
});

client.on("error", function (error) {
    console.error(error);
});

// enable notify-keyspace-events for all kind of events (can be refined)
client.config('set', 'notify-keyspace-events', 'KEA');

client.subscribe('__keyevent@0__:set');
// you can target a specific key with a second parameter
// example, client_redis.subscribe('__keyevent@0__:set', 'mykey')

client.on('message', function (channel, key) {

    redisData.getAllGames().then(res => {
        io.emit('redis', { redis: res })
    })

    mongoData.allData().then(res => {
        io.emit('mongo', { mongo: res })
    })

    redisData.top10Jugadores().then(res => {
        io.emit('top', { top: res })
    })

    mongoData.topGames().then(res => {
        io.emit('topGames', { topGames: res })
    })

    mongoData.topServicios().then(res => {
        io.emit('topServicios', { topServicios: res })
    })

});

//client.set("key", "value", redis.print);
//client.get("key", redis.print);