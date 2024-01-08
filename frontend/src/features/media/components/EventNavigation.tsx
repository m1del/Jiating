import React from 'react';

const EventNavigation = ({ events, onEventSelect }) => {
  return (
    <div className="flex overflow-x-auto py-2 space-x-2 bg-gray-200 p-4 rounded-lg shadow-md">
      {events.map(event => (
        <button
          key={event}
          className="bg-cyan-500 hover:bg-cyan-700 text-white 
          font-semibold py-2 px-4 rounded focus:outline-none focus:ring-2 focus:ring-green-200 
          transition ease-in-out duration-150"
          onClick={() => onEventSelect(event)}
        >
          {event}
        </button>
      ))}
    </div>
  );
};

export default EventNavigation;
