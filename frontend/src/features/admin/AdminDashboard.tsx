// components/AdminDashboard.tsx
import React, { useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';

const AdminDashboard = () => {
  const { user, setUser, isAuthenticated } = useAuth();

  useEffect(() => {
    fetch('http://localhost:3000/api/session-info', {
        credentials: 'include', // include cookies
    })
        .then(res => {
            if (res.ok) {
                return res.json();
            } else {
                throw new Error('Not Authenticated');
            }
        })
        .then(data => {
            if (data.authenticated) {
                setUser({
                    name: data.name,
                    email: data.email,
                })
            }
        })
        .catch(error => {
            console.log("Authentication error: ", error);
        })
  }, [setUser]);

  useEffect(() => {
    if (!isAuthenticated) {
        // redirect to login page if not authenticated
      window.location.href = '/auth/google';
    }
  }, [isAuthenticated]);

  return (
    <div>
        <h1>Admin Dashboard</h1>
        {user && <p>Welcome, {user.name}</p>}
    </div>
  );
};

export default AdminDashboard;
