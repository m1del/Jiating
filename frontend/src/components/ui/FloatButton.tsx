function FloatButton({ path, text }: { path: string; text: string }) {
  return (
    <div className="flex flex-col items-center space-y-2 text-lg font-semibold uppercase text-cyan-400">
      <img className="h-auto max-w-[4rem]" src={path} />
      <span className="text-cyan">{text}</span>
    </div>
  );
}

export default FloatButton;
