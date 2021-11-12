const io = require('./socket')

const redis = require("redis");
const dataMongo = require('./mongoDb');

const dataRedis = {}

const client = redis.createClient({
    host: process.env.REDIS_HOST,
    port: process.env.REDIS_PORT,
    password: process.env.REDIS_PASS,
});

dataRedis.getAllGames = () => {
    return new Promise((resolve, reject) => {
        client.keys('*', function (err, keys) {

            client.mget(keys, function (error, val) {
                resolve(val)
            })
        })
    })
}


dataRedis.top10Jugadores = () => {
    return new Promise((resolve, reject) => {

        let topTen = []

        dataRedis.getAllGames().then(res => {

            if (res !== undefined) {

                for (const iterator of res) {
                    const game = JSON.parse(iterator)

                    if (topTen.length !== 0) {
                        for (let index = 0; index < topTen.length; index++) {
                            const jugador = topTen[index]

                            if (jugador.jugador === game['ganador']) {
                                topTen[index].ganados++
                                index = topTen.length + 1
                            } else if (index === (topTen.length - 1)) {
                                topTen.push({ jugador: game['ganador'], ganados: 1 })
                                index = topTen.length + 1
                            }

                        }
                    } else {
                        topTen.push({ jugador: game['ganador'], ganados: 1 })
                    }

                }

            }

            const sortable = topTen.sort(function (a, b) {
                return b.ganados - a.ganados;
            })

            const slicedArray = sortable.slice(0, 10)

            resolve(slicedArray)
        })

    })
}

// dataRedis.top10Jugadores().then(res => {
//     console.log(res)
// })

module.exports = dataRedis