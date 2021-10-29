import io from 'socket.io-client'

let socket = io('http://localhost:8080')

//let socket = io('https://api-sopes1-py1.ue.r.appspot.com/')

export default socket;