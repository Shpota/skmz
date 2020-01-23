import React, {Component} from "react";
import "./Skill.css"

export class Skill extends Component {
    render() {
        const skill = this.props.skill;
        const textStyle = skill.importance > 1 ? "boldText" : "";
        return <div className="skill card-panel lighten-5">
            {
                skill.icon ? <i className={`fab fa-lg fa-${skill.icon}`}/> : ""
            } <span className={textStyle}>{skill.name}</span>
        </div>
    }
}