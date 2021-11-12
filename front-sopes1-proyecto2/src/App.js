import './App.css';
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import React, { useEffect } from 'react'

import NavBar from './components/NavBar'

import { Games } from './pages/Games'
import { Last10Games } from './pages/Last10Games'
import { Top10Players } from './pages/Top10Players'
import { MongoTransactions } from './pages/MongoTransactions'
import { PlayerDetail } from './pages/PlayerDetail'

function App() {

    useEffect(() => {
        
    }, [])

    return (
        <Router>
            <NavBar />
            <Switch>
                <Route exact path="/">
                    <Games />
                </Route>
                <Route exact path="/last-games">
                    <Last10Games />
                </Route>
                <Route exact path="/top-players">
                    <Top10Players />
                </Route>
                <Route exact path="/mongo-transactions">
                    <MongoTransactions />
                </Route>
                <Route exact path="/player-detail/:id">
                    <PlayerDetail />
                </Route>
            </Switch>
        </Router>
    );
}

export default App;
