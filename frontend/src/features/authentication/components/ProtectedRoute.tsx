import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';

const ProtectedRoute: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
    const [isLoading, setIsLoading] = useState(true);
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        fetch('http://localhost:3000/auth/session-info', { credentials: 'include' })
            .then(res => {
                if (res.ok) {
                    return res.json();
                } else {
                    throw new Error('Not Authenticated');
                }
            })
            .then(data => {
                setIsAuthenticated(data.authenticated);
                setIsLoading(false);
            })
            .catch(error => {
                console.error("Authentication error: ", error);
                setIsLoading(false);
            });
    }, []);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    // If the user is authenticated, return the protected component
    // Else, redirect the user to the home lol TODO: add need to login message in a page?
    return isAuthenticated ? children : <Navigate to="/" />;
};

export default ProtectedRoute;
