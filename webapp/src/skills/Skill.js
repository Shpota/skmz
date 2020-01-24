import React from "react";
import "./Skill.css"

export function Skill(props) {
    const skill = props.skill;
    const textStyle = skill.importance > 1 ? "boldText" : "";
    return <div className="skill card-panel lighten-5">
        {
            skill.icon ? <i className={`fab fa-lg fa-${skill.icon}`}/> : ""
        } <span className={textStyle}>{skill.name}</span>
    </div>
}