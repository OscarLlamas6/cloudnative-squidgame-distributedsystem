const app = require('./app')
const redisData = require('./redisDb')
const mongoData = require('./mongoDb')

const http = require('http')
const server = http.createServer(app)

const io = require('socket.io')(server, { cors: { origin: '*' } })

io.on('connection', (socket) => {

    socket.on('conectado', () => {

        redisData.getAllGames().then(res => {
            io.emit('redis', { redis: res })
        })

        mongoData.allData().then(res => {
            io.emit('mongo', { mongo: res })
        })

        redisData.top10Jugadores().then(res => {
            io.emit('top', { top: res })
        })

    })

})

server.listen(process.env.NODE_API_PORT || 8080)
console.log('Server on port', process.env.NODE_API_PORT || 8080)

module.exports = io