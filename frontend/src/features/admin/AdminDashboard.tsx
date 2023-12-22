// components/AdminDashboard.tsx
import React, { useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';

const AdminDashboard = () => {
    console.log('AdminDashboard');
    const { setAuthUser, setIsLoggedin } = useAuth();

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
            if (data.authenticated) {
              setAuthUser({ email: data.email, name: data.name });
              setIsLoggedin(true);
            } else {
              setIsLoggedin(false);
              window.location.href = 'http://localhost:3000/auth/google';
            }
          })
          .catch(error => {
            console.log("Authentication error: ", error);
            setIsLoggedin(false);
            window.location.href = 'http://localhost:3000/auth/google';
          });
    }, [setAuthUser, setIsLoggedin]);
      
      
      
  
    return (
        <div>
            <h1>Admin Dashboard</h1>
        </div>
    );
};

export default AdminDashboard;
