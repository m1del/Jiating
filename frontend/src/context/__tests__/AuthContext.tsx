import { render } from '@testing-library/react';
import React from 'react';
import AuthContext from '../AuthContext';


// Example of a test for a React Context
describe('AuthContext', () => {
    it('renders without crashing', () => {
        render(
            <AuthContext.Provider value={{ isAuthenticated: false }}>
                <div>Test Component</div>
            </AuthContext.Provider>
        );
    });
});
