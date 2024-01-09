import React from 'react';
import FloatBar from './FloatBar';

function HomeHeading() {
  //Aspect ratio of image is 1440 x 520.5
  // height should be 520.5 / 1440 * 100 = 36.14583333333333
  return (
    <div className="h-full w-full">
      <div className='relative flex w-full items-center bg-[url("../header.png")] bg-contain bg-no-repeat pb-[36.146%]'>
        <h1 className="absolute bottom-[40%] left-[10%] text-2xl uppercase tracking-wide text-white sm:bottom-[38%] sm:text-3xl md:bottom-[40%] md:text-5xl lg:bottom-[40%] lg:text-6xl xl:bottom-[40%] xl:translate-y-[10%] xl:text-7xl 2xl:translate-y-[25%] 2xl:text-8xl">
          Dance <br /> with <br />
          <span className="text-cyan">Jiating</span>
        </h1>
        <FloatBar />
      </div>
    </div>
  );
}

export default HomeHeading;
