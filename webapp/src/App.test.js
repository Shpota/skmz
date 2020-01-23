import React from 'react';
import {render} from '@testing-library/react';
import {App} from './App';

test('renders sanity check', () => {
    const {container} = render(<App/>);
    expect(container).toBeInTheDocument();
});
