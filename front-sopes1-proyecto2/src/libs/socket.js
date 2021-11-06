import env from "react-dotenv"
import io from 'socket.io-client'
require('dotenv').config()



let socket = io({ path: "/socket.io" });
//let socket = io('http://localhost:4000')



export default socket;