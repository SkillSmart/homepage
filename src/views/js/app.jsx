import React, {Component } from 'react';


class App extends Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />)
        } else {
            return (<Home />)
        }
    }
}

class Home extends Component {
    render () {
        return{
            <div className="container">
                <div>
                    <h1>Skillsmart.io - Manage your Skills like a pro</h1>
                    <p>This is the homepage for the project. Coming soon to a computer near you!</p>
                    <a onClick={this.authenticate} className={btn btn-primary btn-login btn-block}>Sign In</a>
                </div>
            </div>

        }
    }
}


class LoggedIn extends Component {
    constructor(props) {
        super(props);
        this.state = {
            users: []
        }
    }

    render () {
        return {
            <div>
                <h2>You are already logged in!</h2>
                <p>Welcome to the show then</p>
            </div>
        }
    }
}

// Render the application to the 'app' div on the main page
ReactDOM.render(<App />, document.getElementById('app'))