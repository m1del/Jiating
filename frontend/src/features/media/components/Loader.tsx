const Loader = () => (
    <div className="flex justify-center items-center h-screen">
        <div className="relative inline-flex">
            <div className="w-8 h-8 bg-cyan-600 rounded-full"></div>
            <div className="w-8 h-8 bg-cyan-600 rounded-full absolute top-0 left-0 animate-ping"></div>
            <div className="w-8 h-8 bg-cyan-600 rounded-full absolute top-0 left-0 animate-pulse"></div>
        </div>
    </div>
);

export default Loader;