import React from 'react';
import { bwLogo } from '../../assets';

function Hero() {
  return (
    <div className="mt-16 flex max-w-[1350px] flex-col items-center justify-around gap-8 sm:my-16 sm:gap-16 lg:my-32 lg:flex-row lg:gap-8 xl:gap-24">
      <img className="w-[40%] lg:ml-10" src={bwLogo}></img>
      <div className=" flex flex-col items-center sm:ml-10 lg:inline xl:w-[50%]">
        <div className="flex w-full justify-center sm:flex-none">
          <h2 className="mb-5 flex h-16 w-[63%] justify-center leading-9 sm:h-20 sm:justify-normal md:h-20 lg:w-full">
            <div className="border-cyan border-l-4 pl-4">
              <span className="text-cyan text-3xl font-semibold uppercase md:text-4xl">
                Jiating
              </span>
              <br />
              <p className="text-lg font-semibold uppercase sm:text-xl md:text-2xl">
                How we got started
              </p>
            </div>
          </h2>
        </div>
        <div className="mb-10 mt-0 flex w-[80%] flex-col text-center text-base text-gray-600 sm:w-[66%] sm:text-left sm:text-2xl lg:mt-5 lg:w-5/6  xl:text-2xl">
          <p className="leading-8 lg:leading-9">
            JiaTing is an all-inclusive cultural and leadership based
            organization that aims to promote traditional/modern lion and dragon
            dance through training and performance.
          </p>
          <br />
          <p className="leading-8 lg:leading-9">
            Originating from the University of Florida's Chinese American
            Student Association (CASA), JiaTing seeks to support a greater
            understanding of Asian heritage in the Gainesville community and is
            open to all members regardless of previous experience or background.
          </p>
        </div>
      </div>
    </div>
  );
}
export default Hero;
