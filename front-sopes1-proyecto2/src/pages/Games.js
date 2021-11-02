import React, { useEffect, useState } from 'react'

import { withStyles, makeStyles } from '@material-ui/core/styles';

import socket from '../libs/socket'

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

const useStyles = makeStyles({
    table: {
        minWidth: 700,
    },
})

export const Games = () => {
    const classes = useStyles()
    const [dataMongo, setDataMongo] = useState([])

    useEffect(() => {
        socket.emit('conectado')
    }, [])

    useEffect(() => {
        socket.on('mongo', (data) => {
            console.log(data)
            setDataMongo(data.mongo)
        })
    },[dataMongo])

    return (
        <div style={{ padding: '2%' }}>
            <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="customized table">
                    <TableHead>
                        <TableRow>
                            <StyledTableCell align="center">Index</StyledTableCell>
                            <StyledTableCell align="center">Id Game</StyledTableCell>
                            <StyledTableCell align="center">Game name</StyledTableCell>
                            <StyledTableCell align="center">Winner</StyledTableCell>
                            <StyledTableCell align="center">Players</StyledTableCell>
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
                            </StyledTableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    )
}
