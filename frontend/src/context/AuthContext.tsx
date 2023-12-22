import React, { createContext, useContext, useState } from 'react';

type UserType = {
  id: string;
  email: string;
  name: string;
  avatar_url: string;
};

type AuthContextType = {
  authUser: UserType | null;
  setAuthUser: (user: UserType | null) => void;
  isLoggedin: boolean;
  setIsLoggedin: (loggedIn: boolean) => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: React.PropsWithChildren<{}>) {
  const [authUser, setAuthUser] = useState<UserType | null>(null);
  const [isLoggedin, setIsLoggedin] = useState(false);

  const value = {
    authUser,
    setAuthUser,
    isLoggedin,
    setIsLoggedin,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth(): AuthContextType {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
