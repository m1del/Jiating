import React from 'react';
import FloatButton from './FloatButton';
function FloatBar() {
  return (
    <>
      <div className="sm: absolute bottom-0 left-[50%] inline-flex w-[45%] -translate-x-[50%] translate-y-[50%] items-center justify-center space-x-16 rounded-xl border border-solid border-gray-300 bg-gray-50 px-36 py-2 shadow-[3px_4px_5px_2px_rgba(0,255,255,0.3)] md:translate-y-0 md:px-48 md:py-4">
        <FloatButton path="icons/chinese_head.png" text="Book" />
        <FloatButton path="icons/lantern.png" text="Join" />
        <FloatButton path="icons/news.png" text="News" />
      </div>
    </>
  );
}

export default FloatBar;
