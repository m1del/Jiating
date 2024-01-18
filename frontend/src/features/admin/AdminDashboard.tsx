// components/AdminDashboard.tsx
import { useState } from 'react';
import { Navigate } from 'react-router-dom';
import { Button } from '../../components';
import { useAuth } from '../../context/useAuth';
import { logoutGoogleUser } from '../../services/authService';
import { Drafts } from '../event';
import {
  AdminList,
  CreateAdmin,
  CreateEventButton,
  DeleteAdmin,
  UserInfo,
} from './components';
import DropdownMenu from './form/Dropdown';
import Position from './form/Position';

const AdminDashboard = () => {
  const { authUser, isAuthenticated} = useAuth();
  const [position, setPosition] = useState('');

  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }

  const handlePositionChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPosition(e.target.value);
  }

  return (
    <div className="container mx-auto p-4">
      <div className="flex flex-wrap">

        <div className="w-full md:w-1/3 px-2 rounded-lg bg-white">
          <div className='flex flex-col justify-center items-center'>
            {authUser && <UserInfo user={authUser}/>}

            <DropdownMenu/>
            <div className='m-4'>
              <Position value={position} onChange={handlePositionChange}/>
            </div>

            <Button 
              buttonText="Logout"
              onClick={() => logoutGoogleUser()}
              additionalClasses="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
            />
          </div>
        </div>

        <div className="w-full md:w-2/3 px-2 min-w-48">
          <AdminList/>
          <div className='bg-white flex m-4 mx-auto rounded-lg p-4 justify-center w-full'>
            <CreateAdmin/>
            <DeleteAdmin/>
          </div>
        </div>

      </div>
      <div className="container m-4 mx-auto p-6 w-full min-h-80 rounded-lg bg-white shadow">
          <Drafts/>
          <CreateEventButton/>
      </div>
    </div>
  );
};

export default AdminDashboard;
