import React from 'react';
import CardContainer from './ProductCards';
import Nav from './Navigation';
import { SignInModalWindow, BuyModalWindow } from './modalwindows';
import About from './About';
import Orders from './orders';

import { BrowserRouter as Router, Route } from "react-router-dom";

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            user: {
                loggedin: false,
                name: ""
            }
        };
    }
    showSignInModalWindow() {
        const state = this.state;
        const newState = Object.assign({}, state, { showSignInModal: true });
        this.setState(newState);
    }
    showBuyModalWindow(id, price) {
        const state = this.state;
        const newState = Object.assign({}, state, { showBuyModal: true, productid: id, price: price });
        this.setState(newState);
    }
    toggleSignInModalWindow() {
        const state = this.state;
        const newState = Object.assign({}, state, { showSignInModal: !state.showSignInModal });
        this.setState(newState);
    }
    toggleBuyModalWindow() {
        const state = this.state;
        const newState = Object.assign({}, state, { showBuyModal: !state.showBuyModal });
        this.setState(newState);
    }
    render() {
        return (
            <div>
                <Router>
                    <div>
                        <Nav user={this.state.user} />
                        <div className='container pt-4 mt-4'>
                            <Route exact path="/" render={() => <CardContainer location='cards.json' />} />
                            <Route path="/promos" render={() => <CardContainer location='promos.json' promo={true} />} />
                            {this.state.user.loggedin ? <Route path="/myorders" render={() => <Orders location='user.json' />} /> : null}
                            <Route path="/about" component={About} />
                        </div>
                        <SignInModalWindow />
                        <BuyModalWindow />
                    </div>
                </Router>
            </div>
        );
    }
}

export default App;
