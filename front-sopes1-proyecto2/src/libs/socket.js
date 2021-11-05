import env from "react-dotenv"
import io from 'socket.io-client'
require('dotenv').config()



//let socket = io('http://localhost:8080')

let ruta = `http://${env.BACKEND_HOST}:${env.NODE_API_PORT}`

let socket = io(ruta)

console.log(ruta)

export default socket;