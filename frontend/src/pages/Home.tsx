import React from 'react';
import FloatBar from '../components/FloatBar';
import Hero from '../components/Hero';
import HomeHeading from '../components/HomeHeading';
import AdminForm from '../features/admin/components/AdminForm';
import AdminList from '../features/admin/components/AdminList';
import { GoogleLoginButton } from '../features/authentication';

function Home() {
  return (
    <div className="mx-auto max-w-[1920px]">
      <div className="flex flex-col items-center">
        <HomeHeading />
        <FloatBar />

        <GoogleLoginButton className='m-8'/>
        <AdminForm/>
        <AdminList/>

        <Hero />
      </div>
    </div>
  );
}

export default Home;
