import { FC } from 'react';

interface YearNavigationProps {
  years: string[];
  onYearSelect: (year: string) => void;
  selectedYear: string;
}

const YearNavigation: FC<YearNavigationProps> = ({ years, onYearSelect, selectedYear }) => {
  return (
    <div className="flex overflow-x-auto space-x-3 bg-white shadow-md rounded-md py-2 px-2">
      {years.map((year) => (
        <button
          key={year}
          className={`bg-cyan-600 hover:bg-cyan-700 text-white font-medium py-2 px-4 
          rounded focus:outline-none focus:ring-2 focus:ring-cyan-300 transition duration-300
          ${year === selectedYear ? 'bg-gradient-to-r from-cyan-800 to-cyan-600 border border-cyan-300 shadow-lg' : 'border border-transparent'}
          `}
          onClick={() => onYearSelect(year)}
        >
          {year}
        </button>
      ))}
    </div>
  );
};

export default YearNavigation;
