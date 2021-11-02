import React, { useEffect, useState } from 'react'
import socket from '../libs/socket'
import { useParams } from "react-router-dom";

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import TableCell from '@material-ui/core/TableCell';
import { withStyles, makeStyles } from '@material-ui/core/styles';


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


export const PlayerDetail = () => {
    const params = useParams()
    const classes = useStyles()
    const [playerData, setPlayerData] = useState([])

    useEffect(() => {
        socket.emit('conectado')
    }, [])

    useEffect(() => {
        socket.on('redis', (data) => {
            setPlayerData(data.redis)
        })
    })

    return (
        <div style={{ padding: '2%' }}>
            <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="customized table">
                    <TableHead>
                        <TableRow>
                            <StyledTableCell align="center">Game</StyledTableCell>
                            <StyledTableCell align="center">Game name</StyledTableCell>
                            <StyledTableCell align="center">Winner</StyledTableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {playerData.map((row, index) => {
                            if ((JSON.parse(row)).ganador === params.id) {
                                return (
                                    <StyledTableRow key={index}>
                                        <StyledTableCell align="center">{(JSON.parse(row)).gamenumber}</StyledTableCell>
                                        <StyledTableCell align="center">{(JSON.parse(row)).gamename}</StyledTableCell>
                                        <StyledTableCell align="center">{(JSON.parse(row)).ganador}</StyledTableCell>
                                    </StyledTableRow>)
                            } else {
                                return null
                            }
                        }
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    )
}