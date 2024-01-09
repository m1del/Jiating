import React from 'react';
function FloatButton({ path, text }: { path: string; text: string }) {
  return (
    <div className="flex flex-col items-center text-sm font-semibold uppercase text-cyan-400 sm:text-base lg:text-lg">
      <img
        className="h-auto max-w-[1.3rem] sm:max-w-[2rem] lg:max-w-[4rem]"
        src={path}
      />
      <span className="text-cyan">{text}</span>
    </div>
  );
}

export default FloatButton;
