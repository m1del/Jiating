import FloatButton from "./ui/FloatButton";
function FloatBar() {
  return (
    <>
      <div className="bg-gray-50 rounded inline-flex px-48 py-4 space-x-24 border-solid border-2 border-cyan-400 shadow-[0_4px_6px_-1px_rgba(0,255,255,0.3)]">
        <FloatButton path="icons/chinese_head.png" text="Book" />
        <FloatButton path="icons/lantern.png" text="Join" />
        <FloatButton path="icons/news.png" text="News" />
      </div>
    </>
  );
}

export default FloatBar;
