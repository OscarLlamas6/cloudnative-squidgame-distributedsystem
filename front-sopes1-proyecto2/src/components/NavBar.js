import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton'
import Typography from '@material-ui/core/Typography'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

export default function NavBar() {
  const classes = useStyles();


  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton href="/" edge="start" color="inherit" aria-label="menu">
          </IconButton>
          <Button className={classes.title} href="/" color="inherit">
            <Typography variant="h6">
              Games
            </Typography>
          </Button>
          <Button className={classes.title} href="/last-games" color="inherit">
            <Typography variant="h6">
              Last 10 games
            </Typography>
          </Button>
          <Button className={classes.title} href="/top-players" color="inherit">
            <Typography variant="h6">
              Top 10 players
            </Typography>
          </Button>
          <Button className={classes.title} href="/mongo-transactions" color="inherit">
            <Typography variant="h6">
              MongoDb transactions
            </Typography>
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}