import React from "react";
import {Skill} from "./Skill";
import "./Skill.css";

export function Skills(props) {
    const skills = props.skills
        .sort((a, b) => b.importance - a.importance)
        .map((sk, i) => <Skill
            skill={sk} highlighted={highlighted(sk, props.search)}
            key={i} updateSearch={props.updateSearch}/>);
    return <div className="skills col row valign-wrapper">{skills}</div>
}

function highlighted(skill, search) {
    return search && skill.name.toLowerCase().startsWith(search.toLowerCase())
}