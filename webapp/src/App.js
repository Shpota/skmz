import React, {Component} from 'react';
import ApolloClient, {gql} from 'apollo-boost';
import {Programmers} from "./programmers/Programmers";
import {SearchBox} from "./search/SearchBox";

export class App extends Component {
    client = new ApolloClient({
        uri: `${process.env.REACT_APP_API_URL}/query`
    });
    state = {
        programmers: [],
        search: ""
    };
    updateSearch = (search) => {
        this.setState({search: search.trim()});
        this.requestProgrammers(search.trim())
    };

    componentDidMount() {
        this.requestProgrammers(this.state.search);
    }

    requestProgrammers(search) {
        this.client.query({
            query: gql`
                {
                    programmers(skill: "${search}") {
                        name,
                        title,
                        picture,
                        company,
                        skills{
                            name,
                            icon,
                            importance
                        }
                    }
                }
            `
        })
        .then(result => this.setState({
            programmers: result.data.programmers
        }));
    }

    render() {
        return <div className="container collection">
            <SearchBox search={this.state.search} updateSearch={this.updateSearch}/>
            <Programmers programmers={this.state.programmers}
                         search={this.state.search} updateSearch={this.updateSearch}/>
        </div>;
    }
}
