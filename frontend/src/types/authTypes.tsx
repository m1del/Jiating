export type UserType = {
  id: string;
  email: string;
  name: string;
  avatar_url: string;
};

export type AuthContextType = {
  authUser: UserType | null;
  setAuthUser: (user: UserType | null) => void;
  isAuthenticated: boolean;
  checkAuthentication: () => Promise<void>;
};

export type AuthProviderProps = {
  children: React.ReactNode; // any valid React child: JSX, string, array of JSX
};
