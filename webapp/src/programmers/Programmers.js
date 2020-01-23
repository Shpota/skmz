import React, {Component} from "react";
import {Programmer} from "./Programmer";
import "./Progremmer.css";

export class Programmers extends Component {
    render() {
        const programmers = this.props.programmers.map(
            (pr, i) => <Programmer programmer={pr} key={i}/>
        );
        return <div>{programmers}</div>
    }
}