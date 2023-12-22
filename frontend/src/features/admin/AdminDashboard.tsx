// components/AdminDashboard.tsx
import React, { useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';
import GoogleLogoutButton from '../authentication/components/GoogleLogoutButton';

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
        <div className='flex flex-col items-center justify-center'>
            <h1 className='text-3xl text-bold mb-4'>Admin Dashboard</h1>
            {authUser && (
                <div className='flex justify-center items-center'>
                    <p className='text-lg'>hi, <span>{authUser.email}</span></p>
                    <img className='rounded-full m-5'
                    src={authUser.avatar_url} alt={authUser.name} />
                </div>
            )}
            <GoogleLogoutButton />
        </div>
    );
};

export default AdminDashboard;
