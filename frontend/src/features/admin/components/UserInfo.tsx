import React from 'react';

type UserType = {
  name: string;
  email: string;
  avatar_url: string;
};

type UserInfoProps = {
  user: UserType;
};

const UserInfo: React.FC<UserInfoProps> = ({ user }) => {
  return (
    <div className='flex flex-col justify-center items-center bg-white p-4 rounded-lg shadow'>
      <img className='w-24 h-24 rounded-full border-2 border-gray-300 m-3'
           src={user.avatar_url} alt={user.name} />
      <p className='text-xl text-gray-700'>
          Hi, <span className='font-semibold'>{user.name}</span>
      </p>
      <p className='text-md text-gray-600 mt-2'>
          Email: <span className='font-semibold'>{user.email}</span>
      </p>
    </div>
  );
}

export default UserInfo;
