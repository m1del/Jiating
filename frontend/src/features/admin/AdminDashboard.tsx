// components/AdminDashboard.tsx
import { useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';
import { GoogleLogoutButton } from '../authentication';
import CheckAuth from './CheckAuth';
import {
  AdminForm,
  AdminList,
  CreateEventButton,
  UserInfo,
} from './components';

const AdminDashboard = () => {
  const { authUser, setAuthUser, setIsLoggedin } = useAuth();

  useEffect(() => {
    CheckAuth(setAuthUser, setIsLoggedin);
  }, [setAuthUser, setIsLoggedin]);

  return (
    <div className="flex flex-col items-center justify-center rounded-lg bg-gray-100 p-10 shadow-md">
      <h1 className="mb-6 text-4xl font-bold text-gray-800">Admin Dashboard</h1>

      {authUser && <UserInfo user={authUser} />}

      <AdminForm />
      <AdminList />

      <CreateEventButton />

      <div className="mt-5">
        <GoogleLogoutButton />
      </div>
    </div>
  );
};

export default AdminDashboard;
