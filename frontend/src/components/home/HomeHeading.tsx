import React from 'react';

function HomeHeading() {
  return (
    <div className="h-full w-full">
      <div className='flex h-[44rem] items-center bg-[url("../header.png")] bg-contain bg-no-repeat'>
        <h1 className=" mb-28 ml-72 text-8xl uppercase tracking-wide text-white">
          Dance <br /> with <br />
          <span className="text-cyan">Jiating</span>
        </h1>
      </div>
    </div>
  );
}

export default HomeHeading;
