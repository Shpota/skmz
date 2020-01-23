import React, {Component} from 'react';
import './App.css';
import ApolloClient, {gql} from 'apollo-boost';
import {Programmers} from "./programmers/Programmers";

export class App extends Component {

    state = {
        programmers: []
    };

    constructor(props) {
        super(props);
        this.client = new ApolloClient({
            uri: '/query',
        });
    }

    componentDidMount() {
        this.client.query(
            {
                query: gql`
                    {
                        programmers {
                            name,
                            picture,
                            company,
                            skills{
                                name,
                                icon,
                                importance
                            }
                        }
                    }
                `,
            })
        .then(result => this.setState({
            programmers: result.data.programmers
        }));
    }

    render() {
        return <div className="container collection">
            <Programmers programmers={this.state.programmers}/>
        </div>;
    }
}
