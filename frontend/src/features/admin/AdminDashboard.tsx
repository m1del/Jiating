// components/AdminDashboard.tsx
import { Navigate } from 'react-router-dom';
import { useAuth } from '../../context/useAuth';
import { GoogleLogoutButton } from '../authentication';
import {
  AdminList,
  CreateAdmin,
  CreateEventButton,
  DeleteAdmin,
  UserInfo,
} from './components';

const AdminDashboard = () => {
  const { authUser, isAuthenticated} = useAuth();

  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }

  return (
    <div className="flex flex-col items-center justify-center rounded-lg bg-gray-100 p-10 shadow-md">
      <h1 className="mb-6 text-4xl font-bold text-gray-800">Admin Dashboard</h1>

      {authUser && <UserInfo user={authUser} />}

      <CreateAdmin />
      <DeleteAdmin/>
      <AdminList />

      <CreateEventButton />

      <div className="mt-5">
        <GoogleLogoutButton />
      </div>
    </div>
  );
};

export default AdminDashboard;
