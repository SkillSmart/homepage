import React from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import Aux from 'react-aux';

import {LandingPage} from './pages';



const App = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={LandingPage}/>
        </Switch>
    </BrowserRouter>
);

export default App;