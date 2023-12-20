function FloatButton({ path, text }: { path: string; text: string }) {
  return (
    <div className="flex flex-col items-center space-y-2 text-cyan-400 uppercase font-semibold text-s">
      <img className="h-auto max-w-[4rem]" src={path} />
      <span>{text}</span>
    </div>
  );
}

export default FloatButton;
