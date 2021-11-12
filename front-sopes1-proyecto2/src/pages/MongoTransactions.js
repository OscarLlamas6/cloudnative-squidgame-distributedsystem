import React, { useEffect, useState } from 'react'
import { Pie } from '../../node_modules/react-chartjs-2';
import socket from '../libs/socket'
import { makeStyles, withStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Link from '@material-ui/core/Link';
import TableCell from '@material-ui/core/TableCell';

const StyledTableCell = withStyles((theme) => ({
    head: {
        backgroundColor: '#5788CA',
        color: theme.palette.common.white,
    },
    body: {
        fontSize: 14,
    },
}))(TableCell)

const StyledTableRow = withStyles((theme) => ({
    root: {
        '&:nth-of-type(odd)': {
            backgroundColor: '#E9F1FD',
        },
    },
}))(TableRow)

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },
}));

export const MongoTransactions = () => {
    const classes = useStyles();

    const [dataMongo, setDataMongo] = useState([])

    const [juegos, setJuegos] = useState([])
    const [numeroJuegos, setNumeroJuegos] = useState([])

    const [servicios, setServicios] = useState([])
    const [numeroServicios, setNumeroServicios] = useState([])

    useEffect(() => {
        socket.emit('conectado')
    }, [])

    useEffect(() => {
        socket.on('mongo', (data) => {
            console.log(data)
            setDataMongo(data.mongo)
        })
    }, [dataMongo])

    useEffect(() => {
        socket.on('topGames', (data) => {
            console.log(data)
            setJuegos(data.topGames.juegos)
            setNumeroJuegos(data.topGames.numeroJuegos)
        })
    }, [juegos])


    useEffect(() => {
        socket.on('topServicios', (data) => {
            console.log(data)
            setServicios(data.topServicios.servicios)
            setNumeroServicios(data.topServicios.numeroServicios)
        })
    }, [servicios])



    return (
        <div className={classes.root}>
            <Grid container spacing={3}>
                <Grid item xs={6} style={{ padding: '5%' }}>

                    <Pie data={

                        {
                            labels: juegos,
                            datasets: [
                                {
                                    label: '# of Votes',
                                    data: numeroJuegos,
                                    backgroundColor: [
                                        'rgba(255, 99, 132, 0.2)',
                                        'rgba(54, 162, 235, 0.2)',
                                        'rgba(255, 206, 86, 0.2)',
                                    ],
                                    borderColor: [
                                        'rgba(255, 99, 132, 1)',
                                        'rgba(54, 162, 235, 1)',
                                        'rgba(255, 206, 86, 1)',
                                    ],
                                    borderWidth: 1,
                                },
                            ],
                        }

                    } />


                </Grid>
                <Grid item xs={6} style={{ padding: '5%' }}>


                    <Pie data={
                        {
                            labels: servicios,
                            datasets: [
                                {
                                    label: '# of Votes',
                                    data: numeroServicios,
                                    backgroundColor: [
                                        'rgba(75, 192, 192, 0.2)',
                                        'rgba(153, 102, 255, 0.2)',
                                        'rgba(255, 159, 64, 0.2)',
                                    ],
                                    borderColor: [
                                        'rgba(75, 192, 192, 1)',
                                        'rgba(153, 102, 255, 1)',
                                        'rgba(255, 159, 64, 1)',
                                    ],
                                    borderWidth: 1,
                                },
                            ],
                        }
                    } />


                </Grid>

                <Grid item xs={12}>

                    <TableContainer component={Paper}>
                        <Table className={classes.table} aria-label="customized table">
                            <TableHead>
                                <TableRow>
                                    <StyledTableCell align="center">Index</StyledTableCell>
                                    <StyledTableCell align="center">Id Game</StyledTableCell>
                                    <StyledTableCell align="center">Game name</StyledTableCell>
                                    <StyledTableCell align="center">Winner</StyledTableCell>
                                    <StyledTableCell align="center">Players</StyledTableCell>
                                    <StyledTableCell align="center">Request</StyledTableCell>
                                    <StyledTableCell align="center">Service</StyledTableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {dataMongo.map((row, index) => (
                                    <StyledTableRow key={row._id}>
                                        <StyledTableCell align="center">{index}</StyledTableCell>
                                        <StyledTableCell align="center">{row.gamenumber}</StyledTableCell>
                                        <StyledTableCell align="center">{row.gamename}</StyledTableCell>
                                        <StyledTableCell align="center">
                                            <Link href={'/player-detail/' + row.ganador} color="inherit">
                                                {row.ganador}
                                            </Link>
                                        </StyledTableCell>
                                        <StyledTableCell align="center">{row.players}</StyledTableCell>
                                        <StyledTableCell align="center">{row.request}</StyledTableCell>
                                        <StyledTableCell align="center">{row.service}</StyledTableCell>
                                    </StyledTableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>

                </Grid>

            </Grid>

        </div>
    )
}


/*
const groupBy = (key, arr) => arr.reduce((cache, item) =>
    ({
        ...cache, [item[key]]:
            item[key] in cache
                ? cache[item[key]].concat(item)
                : [item]
    }),
        {}
    )
*/