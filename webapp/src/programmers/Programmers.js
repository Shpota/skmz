import React from "react";
import {Programmer} from "./Programmer";
import "./Progremmer.css";

export function Programmers(props) {
    const programmers = props.programmers.map(
        (pr, i) => <Programmer programmer={pr} key={i}/>
    );
    return <div>{programmers}</div>
}