import React from 'react';
import {render} from '@testing-library/react';
import {Skill} from './Skill';

test('renders a skill with low importance', () => {
    const skill = {
        name: "Java",
        importance: 1
    };
    const {getByText} = render(<Skill skill={skill}/>);
    const element = getByText(/Java/i);
    expect(element).toBeInTheDocument();
    expect(element.classList.contains('boldText')).toBe(false)
});

test('renders a skill with high importance', () => {
    const skill = {
        name: "Java",
        importance: 2
    };
    const {getByText} = render(<Skill skill={skill}/>);
    const skillName = getByText(/Java/i);
    expect(skillName).toBeInTheDocument();
    expect(skillName.classList.contains('boldText')).toBe(true)
});
