const io = require('./socket')

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
    console.log('Redis cambio', key)
    io.emit('test', { test: 'test' })
});

client.set("key", "value", redis.print);
client.get("key", redis.print);