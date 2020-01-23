import React, {Component} from "react";
import {Skill} from "./Skill";
import "./Skill.css";

export class Skills extends Component {
    render() {
        const skills = this.props.skills.map(
            (sk, i) => <Skill skill={sk} key={i}/>
        );
        return <div className="skills col row valign-wrapper">{skills}</div>
    }
}