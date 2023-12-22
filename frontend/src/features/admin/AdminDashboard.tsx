// components/AdminDashboard.tsx
import React, { useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';
import { GoogleLogoutButton } from '../authentication';
import { UserInfo } from './components';

const AdminDashboard = () => {
    const {authUser, setAuthUser, setIsLoggedin } = useAuth();

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
              setAuthUser({ 
                id: data.id, 
                email: data.email, 
                name: data.name, 
                avatar_url: data.avatar_url 
              });
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
      <div className='flex flex-col items-center justify-center bg-gray-100 p-10 rounded-lg shadow-md'>
          <h1 className='text-4xl font-bold text-gray-800 mb-6'>Admin Dashboard</h1>

          {authUser && <UserInfo user={authUser} />}

          <div className='mt-5'>
              <GoogleLogoutButton />
          </div>
      </div>
  );
};

export default AdminDashboard;
