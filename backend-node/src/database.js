// get the client
const mysql = require('mysql2');

// create the connection to database
const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASS,
    port: process.env.DB_PORT,
    multipleStatements: true
})

let sql_initial = `

CREATE SCHEMA IF NOT EXISTS INTERMEDIAS;

USE INTERMEDIAS;

CREATE TABLE IF NOT EXISTS INTERMEDIAS.ADMIN(
    cui int not null,
    nombre varchar(50) not null,
    apellido varchar(50) not null,
    correo varchar(50) not null,
    contra varchar(50) not null,
    primary key (cui));

`
connection.query(sql_initial, function (err, result) {
    if (err) throw err;
    console.log("Database created and connected");
});


module.exports = connection;