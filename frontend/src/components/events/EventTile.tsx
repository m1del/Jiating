import React from 'react';
import { motion } from 'framer-motion';

function EventTile({ size, gridPos }) {
  const imgPath = '../lion_1.png';
  const convertPath = `'${imgPath}'`;
  let spefClass = '';

  const bounceVariants = {
    initial: { y: 0 }, // Initial position
    animate: {
      y: [-10, 10, -10],
      transition: { ease: 'linear', repeat: Infinity, duration: 3 },
    }, // Animation values and settings
  };

  if (size == 'large') {
    spefClass = `overflow-hidden row-span-2 col-span-2 ${gridPos} hover:-translate-y-4 duration-500 w-full`;
  } else if (size == 'long') {
    spefClass = `col-span-2 w-8/12 rounded-r-none ${gridPos}`;
  } else {
    spefClass = `overflow-hidden relative group w-full ${gridPos}`;
  }
  const eventClass = ` rounded-2xl h-full bg-[url('../lion_1.png')] bg-cover bg-no-repeat ${spefClass} text-white hover:cursor-pointer`;

  if (size == 'large') {
    return (
      <div className={eventClass}>
        <div className="flex h-full flex-col justify-end">
          <div className="bg-black-rgba px-4 py-6">
            <h2 className="pb-5 text-4xl">CNY 2023</h2>
            <p className="pb-2 text-lg">
              Lorem ipsum dolor, sit amet consectetur adipisicing elit.
              Consequatur eaque accusantium tempore deleniti odit cum obcaecati.
            </p>
            <p>Written by: Michael Shi</p>
            <p>05/15/2002</p>
          </div>
        </div>
      </div>
    );
  } else if (size == 'medium') {
    return (
      <div className={eventClass}>
        <div className="absolute h-full w-full bg-gradient-to-t from-black to-transparent opacity-70">
          <div className="bg-black-rgba absolute bottom-0 left-0 right-0 h-0 w-full overflow-hidden transition-all duration-500 ease-in group-hover:h-full">
            <div className="flex h-full flex-col items-center justify-center px-5">
              <h2 className="pb-4 text-2xl">CNY 2023</h2>
              <p className="pb-3 text-center text-lg">
                Lorem ipsum dolor sit amet consectetur adipisicing elit.
                Delectus deleniti, quo eius similique ipsa ab labore.
              </p>
              <p>Written by: Michael Shi</p>
              <p>05/15/2002</p>
            </div>
          </div>
        </div>
        <motion.div
          variants={bounceVariants}
          initial="initial"
          animate="animate"
          style={{ top: '86%', position: 'absolute' }}
          className="arrow bottom-5 group-hover:hidden"
        >
          <span></span>
        </motion.div>
      </div>
    );
  } else if (size == 'long') {
    return (
      <div className={eventClass}>
        <div className="relative flex h-full w-full items-center justify-center">
          <div className="h-1/2 w-1/2"></div>
          <div className="bg-black-rgba absolute left-[22rem] ml-9 flex h-full w-[52.5%] flex-col items-center justify-center rounded-2xl rounded-l-none px-4 text-center">
            <h2 className="pb-4 text-2xl">CNY 2023</h2>
            <p className="pb-3 text-base">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus
              deleniti, quo eius similique ipsa ab labore.
            </p>
            <p>Written by: Michael Shi</p>
            <p>05/15/2002</p>
          </div>
        </div>
      </div>
    );
  }
}

export default EventTile;

// header, imgPath, author, desc, size
