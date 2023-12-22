import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';

const ProtectedRoute = ({ children }) => {
    const [isLoading, setIsLoading] = useState(true);
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        fetch('http://localhost:3000/api/session-info', { credentials: 'include' })
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

    return isAuthenticated ? children : <Navigate to="/login" />;
};

export default ProtectedRoute;
