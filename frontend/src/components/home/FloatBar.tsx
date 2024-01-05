import FloatButton from './FloatButton';
function FloatBar() {
  return (
    <>
      <div className="justify-self inline-flex -translate-y-28 space-x-24 rounded-xl border border-solid border-gray-300 bg-gray-50 px-48 py-4 shadow-[3px_4px_5px_2px_rgba(0,255,255,0.3)]">
        <FloatButton path="icons/chinese_head.png" text="Book" />
        <FloatButton path="icons/lantern.png" text="Join" />
        <FloatButton path="icons/news.png" text="News" />
      </div>
    </>
  );
}

export default FloatBar;
