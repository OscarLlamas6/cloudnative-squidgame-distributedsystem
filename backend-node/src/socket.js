const app = require('./app')

const http = require('http')
const server = http.createServer(app)

const io = require('socket.io')(server, { cors: { origin: '*' } })

io.on('connection', (socket) => {

    console.log('conectado')

    socket.on('conectado', () => {
        console.log('nueva conexion', socket.id)
    })

})

server.listen(process.env.NODE_API_PORT || 8080)
console.log('Server on port', 8080)

module.exports = io