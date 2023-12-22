import React from 'react';
import { useAuth } from '../../../context/AuthContext';

function GoogleLogoutButton() {
    const { setAuthUser, setIsLoggedin } = useAuth();

    const handleLogout = () => {
        try {
            // redirect to backend logout route
            window.location.href = 'http://localhost:3000/logout/google';
            // update auth context
            setAuthUser(null);
            setIsLoggedin(false);
            
            // clear cookies
            document.cookie.split(";").forEach((c) => {
                document.cookie = c.trim().split("=")[0] + "=;expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
            });

        } catch (err) {
            console.error('Logout failed', err);
        }
    }
  return (
    <button className='bg-cyan hover:bg-gray-500 text-white font-bold py-2 px-4 rounded'
    onClick={handleLogout}>
      Logout
    </button>
  )
}

export default GoogleLogoutButton
