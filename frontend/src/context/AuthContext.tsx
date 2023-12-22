// context/AuthContext.tsx
import React, { createContext, useContext, useState } from 'react';

type UserType = {
  name: string;
  email: string;
};

// define a type for the context
type AuthContextType = {
  isAuthenticated: boolean;
  user: UserType | null;
  setUser: (user: UserType | null) => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState<UserType | null>(null);

  return (
    <AuthContext.Provider value={{ isAuthenticated: !!user, user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
