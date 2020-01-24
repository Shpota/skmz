import React from "react";
import "./Skill.css"

export function Skill(props) {
    const skill = props.skill;
    const skillClass = props.highlighted ? "greenSkill" : "greySkill";
    const cl = `skill ${skillClass} card-panel lighten-5`;
    const textStyle = skill.importance > 1 ? "boldText" : "";
    return <div className={cl} onClick={() => props.updateSearch(skill.name)}>
        {
            skill.icon ? <i className={`fab fa-lg fa-${skill.icon}`}/> : ""
        } <span className={textStyle}>{skill.name}</span>
    </div>
}