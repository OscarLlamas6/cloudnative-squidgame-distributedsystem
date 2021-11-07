import env from "react-dotenv"
import io from 'socket.io-client'
require('dotenv').config()

//let socket = io('http://localhost:4000')

let socket = io({   path: "/socket.io",
                    secure: true, 
                    rejectUnauthorized: false,
                    reconnect: true,
                    transports: ['polling','websocket'] });

// let socket = io('/websockettest', {
//   reconnect: true, 
//   //secure: true, 
//   rejectUnauthorized: false,
//   transports: ['polling','websocket']
// });

// let socket = io('http://localhost:4000', {
//   path: "/socket.io",
//   reconnect: true, 
//   secure: true, 
//   rejectUnauthorized: false,
//   transports: ['polling','websocket']
// });


export default socket;