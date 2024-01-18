import React from 'react';

type UserType = {
  id: string;
  name: string;
  email: string;
  avatar_url: string;
};

type UserInfoProps = {
  user: UserType;
};

const UserInfo: React.FC<UserInfoProps> = ({ user }) => {
  return (
    <div className="flex flex-col items-center justify-center rounded-lg bg-white p-4 ">
      <img
        className="m-3 h-24 w-24 rounded-full border-2 border-gray-300"
        src={user.avatar_url}
        alt={user.name}
      />
      <p className="text-xl text-gray-700">
        Hi, <span className="font-semibold">{user.name}</span>
      </p>
      <p className="text-md mt-2 text-gray-600">
        Email: <span className="font-semibold">{user.email}</span>
      </p>
    </div>
  );
};

export default UserInfo;
