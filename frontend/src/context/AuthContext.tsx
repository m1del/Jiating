// AuthContext.tsx
import React, { createContext, useEffect, useState } from 'react';
import { AuthContextType, AuthProviderProps, UserType } from '../types/authTypes';

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [authUser, setAuthUser] = useState<UserType | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const checkAuthentication = async () => {
    try {
      const response = await fetch('http://localhost:3000/auth/session-info', { credentials: 'include' });
      if (!response.ok) throw new Error('Not Authenticated');
      const data = await response.json();
      if (data.authenticated) {
        setAuthUser({
          id: data.userID,
          email: data.email,
          name: data.name,
          avatar_url: data.avatar_url,
        });
        setIsAuthenticated(true);
      } else {
        setIsAuthenticated(false);
        // Handle redirection or display login prompt
      }
    } catch (error) {
      console.error('Authentication error: ', error);
      setIsAuthenticated(false);
      // Handle error appropriately
    }
  };

  useEffect(() => {
    checkAuthentication();
  }, []);

  const value = {
    authUser,
    setAuthUser,
    isAuthenticated,
    setIsAuthenticated,
    checkAuthentication,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export { AuthContext, AuthProvider };
