import React, { useEffect, useState } from 'react'

import { withStyles, makeStyles } from '@material-ui/core/styles';

import socket from '../libs/socket'

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
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

const useStyles = makeStyles({
    table: {
        minWidth: 700,
    },
})

export const Top10Players = () => {
    const classes = useStyles()
    const [topTen, setTopTen] = useState([])

    useEffect(() => {
        socket.emit('conectado')
    }, [])

    useEffect(() => {
        socket.on('top', (data) => {
            console.log(data)
            setTopTen(data.top)
        })
    }, [topTen])

    return (
        <div style={{ padding: '2%' }}>
            <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="customized table">
                    <TableHead>
                        <TableRow>
                            <StyledTableCell align="center">Player</StyledTableCell>
                            <StyledTableCell align="center">Wins</StyledTableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {topTen.map((row, index) => (
                            <StyledTableRow key={index}>
                                <StyledTableCell align="center">{row.jugador}</StyledTableCell>
                                <StyledTableCell align="center">{row.ganados}</StyledTableCell>
                            </StyledTableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    )
}


/*


let topTen = []

            for (const key in data.redis) {
                const game = JSON.parse(data.redis[key])
                console.log(key)
                console.log(game)
                if (key !== 0) {

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

            console.log(topTen)


*/