import React from 'react';
import {render} from '@testing-library/react';
import {App} from './App';

test('renders sanity check', () => {
    const {getByText} = render(<App/>);
    const linkElement = getByText(/Sanity check/i);
    expect(linkElement).toBeInTheDocument();
});
