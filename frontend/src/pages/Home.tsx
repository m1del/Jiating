import React from 'react';
import FloatBar from '../components/home/FloatBar';
import Hero from '../components/home/Hero';
import HomeHeading from '../components/home/HomeHeading';
import { GoogleLoginButton } from '../features/authentication';

function Home() {
  return (
    <div className="mx-auto max-w-[1920px]">
      <div className="flex flex-col items-center">
        <HomeHeading />
        <FloatBar />

        <Hero />
      </div>
    </div>
  );
}

export default Home;
