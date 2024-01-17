import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../../../context/useAuth';

type ProtectedRouteProps = {
    children: React.ReactNode;
};
const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const { isAuthenticated } = useAuth();

  if (isAuthenticated === undefined) {
    // auth status is unknown, show loading TODO: add loading spinner
    return <div>Loading...</div>;
  }

  if (!isAuthenticated) {
    // not authenticated, redirect to login //TODO nothing exists on this route, (how do we want to handle this?)
    return <Navigate to="/login" />;
  }
  // authenticated, render the protected content
  return <>{children}</>;
};


export default ProtectedRoute;
