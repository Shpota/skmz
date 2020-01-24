import React from "react";
import {Skill} from "./Skill";
import "./Skill.css";

export function Skills(props) {
    const skills = props.skills
        .sort((a, b) => b.importance - a.importance)
        .map((sk, i) => <Skill skill={sk} key={i}/>);
    return <div className="skills col row valign-wrapper">{skills}</div>
}