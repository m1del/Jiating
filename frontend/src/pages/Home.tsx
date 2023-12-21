import React from 'react';
import FloatBar from '../components/FloatBar';
import HomeHeading from '../components/HomeHeading';
import Hero from '../components/Hero';

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
