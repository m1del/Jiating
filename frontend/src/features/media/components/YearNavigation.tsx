const YearNavigation = ({ years, onYearSelect }) => {
  return (
    <div className="flex overflow-x-auto py-2 space-x-1 bg-white shadow-md rounded-md">
      {years.map(year => (
        <button
          key={year}
          className="bg-cyan-500 hover:bg-cyan-700 text-white font-medium py-2 px-4 
          rounded focus:outline-none focus:bg-gray-400 transition duration-200"
          onClick={() => onYearSelect(year)}
        >
          {year}
        </button>
      ))}
    </div>
  );
};

export default YearNavigation;
