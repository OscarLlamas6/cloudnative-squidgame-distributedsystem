//const URI = process.env.COSMOSDB_CONNECTION

const URI = `mongodb://${process.env.MONGO_USER}:${process.env.MONGO_PASS}@${process.env.MONGO_HOST}:${process.env.MONGO_PORT}`

const { MongoClient } = require('mongodb');

const client = new MongoClient(URI)

const dataMongo = {}

dataMongo.allData = () => {

    return new Promise((resolve, reject) => {
        client.connect(function () {
            const db = client.db('SQUIDGAMES')
            const collection = db.collection('LOGS')
            collection.find({}).toArray(function (err, result) {
                resolve(result);
            })
        })

    })

}





module.exports = dataMongo

// export const allTweets =  () => {
//     //.db("SOPES1").collection("TWEET")
//     const tweets =  collection()
//     const allTweets =  tweets.find({}).toArray()
//     //console.log(allTweets)
//     return allTweets
// }

//allTweets()

// export const countTweets = async () => {
//     const tweets = await collection()
//     const count = await tweets.count({})
//     //console.log(count)
//     return count
// }

// //countTweets()
// export const countHashTags = async () => {
//     let arra = []
//     const tweets = await collection()
//     const allTweets = await tweets.find({}).toArray()
//     for (const tweet of allTweets) {
//         if (tweet.hashtags !== null && tweet.hashtags !== undefined && Array.isArray(tweet.hashtags)) {

//             for (const iterator of tweet.hashtags) {
//                 arra.push(iterator)
//             }
//         }
//     }
//     let arrb = new Set(arra)
//     //console.log(arrb.size)
//     return arrb.size
// }

// //countHashTags()

// export const countUpvotes = async () => {
//     let count = 0
//     const tweets = await collection()
//     const allTweets = await tweets.find({}).toArray()
//     for (const element of allTweets) {
//         if (element.upvotes !== null && element.upvotes !== undefined && !isNaN(element.upvotes)) {
//             //console.log(parseInt(element.upvotes))
//             count += parseInt(element.upvotes)
//         }
//     }
//     //console.log('Conteo de votos', count)
//     return count
// }

// //countUpvotes()

// export const upvotesVSdownvotes = async () => {
//     let up_array = {}
//     let down_array = {}
//     const tweets = await collection()
//     const allTweets = await tweets.find({}).toArray()
//     for (const element of allTweets) {
//         if (element.upvotes !== null && element.upvotes !== undefined && !isNaN(element.upvotes) && element.fecha !== '') {
//             if (up_array[element.fecha] === undefined || up_array[element.fecha] === null) up_array[element.fecha] = parseInt(element.upvotes)
//             else up_array[element.fecha] += parseInt(element.upvotes)
//         }
//         if (element.downvotes !== null && element.downvotes !== undefined && !isNaN(element.downvotes) && element.fecha !== '') {
//             if (down_array[element.fecha] === undefined || down_array[element.fecha] === null) down_array[element.fecha] = parseInt(element.downvotes)
//             else down_array[element.fecha] += parseInt(element.downvotes)
//         }
//     }

//     let arreglo = [up_array, down_array]
//     //console.log(arreglo)
//     return arreglo
// }

// //upvotesVSdownvotes()

// export const topHashtags = async () => {
//     let arr = {}
//     const tweets = await collection()
//     const allTweets = await tweets.find({}).toArray()
//     for (const tweet of allTweets) {
//         if (tweet.hashtags !== null && tweet.hashtags !== undefined && Array.isArray(tweet.hashtags)) {

//             for (const iterator of tweet.hashtags) {
//                 if (tweet.upvotes !== null && tweet.upvotes !== undefined && !isNaN(tweet.upvotes)) {
//                     if (arr[iterator] === undefined || arr[iterator] === null) arr[iterator] = parseInt(tweet.upvotes)
//                     else arr[iterator] += parseInt(tweet.upvotes)
//                 }
//             }
//         }
//     }

//     const sortable = Object.fromEntries(
//         Object.entries(arr).sort(([, a], [, b]) => b - a)
//     );
//     //console.log(sortable);
//     return sortable
// }

// export const recentPosts = async () => {
//     const tweets = await collection()
//     const count = await tweets.count({})
//     const allTweets = await tweets.find({}).skip((count > 5 ? count - 5 : 0)).limit(5).toArray()
//     //console.log(allTweets)
//     return allTweets
// }