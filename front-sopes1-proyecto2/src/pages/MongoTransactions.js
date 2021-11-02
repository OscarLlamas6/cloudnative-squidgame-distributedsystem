import React, { useEffect } from 'react'
import socket from '../libs/socket'


export const MongoTransactions = () => {

    useEffect(() => {
        socket.emit('conectado')
    }, [])

    useEffect(() => {
        socket.on('mongo', (data) => {
            console.log(data)
        })
    })

    return (
        <div>
            MongoTransactions works!
        </div>
    )
}