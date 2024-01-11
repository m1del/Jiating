import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../../../context/useAuth';

type ProtectedRouteProps = { // bruh typescript is so annoying
    children: React.ReactNode;
};

const ProtectedRoute = ({children}: ProtectedRouteProps) => {
    const { isAuthenticated, checkAuthentication } = useAuth();
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const verifyAuth = async () => {
            await checkAuthentication();
            setIsLoading(false);
        };
        verifyAuth();
    }, [checkAuthentication]);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
};

export default ProtectedRoute;
