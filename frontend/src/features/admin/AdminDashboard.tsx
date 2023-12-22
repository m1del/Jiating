// components/AdminDashboard.tsx
import React, { useEffect } from 'react';
import { useAuth } from '../context/AuthContext';

const AdminDashboard = () => {
  const { user, setUser, isAuthenticated } = useAuth();

  useEffect(() => {
    // Example: Fetch user session info from backend if not already authenticated
    if (!isAuthenticated) {
      fetch('/api/session-info')
        .then(response => response.json())
        .then(data => {
          if (data.user) {
            setUser(data.user);
          }
        })
        .catch(error => console.error('Failed to fetch user session:', error));
    }
  }, [isAuthenticated, setUser]);

  return (
    <div>
      admin dash
    </div>
  );
};

export default AdminDashboard;
