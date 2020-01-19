import React, {Component} from 'react';
import './App.css';
import {Programmers} from "./content/Programmers";

export class App extends Component {
    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <h1>Sanity check</h1>
                </header>
                <Programmers/>
            </div>
        );
    }
}
