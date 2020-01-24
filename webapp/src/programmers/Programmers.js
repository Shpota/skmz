import React from "react";
import {Programmer} from "./Programmer";
import "./Progremmer.css";

export function Programmers(props) {
    const programmers = props.programmers.map(
        (pr, i) => <Programmer programmer={pr} key={i}
                               search={props.search}
                               updateSearch={props.updateSearch}/>
    );
    return <div>{programmers}</div>
}