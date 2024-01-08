const EventNavigation = ({ events, onEventSelect, selectedEvent }) => {
  return (
    <div className="flex overflow-x-auto space-x-4 bg-gray-200 px-4 py-2 rounded-lg shadow-md">
      {events.map(event => (
        <button
          key={event}
          className={`bg-cyan-600 hover:bg-cyan-800 text-white font-medium py-2 px-4 rounded
          focus:outline-none focus:ring-2 focus:ring-cyan-300 transition duration-300
          ${event === selectedEvent ? 'bg-gradient-to-r from-cyan-800 to-cyan-600 shadow-lg' : ''}`}
          onClick={() => onEventSelect(event)}
        >
          {event}
        </button>
      ))}
    </div>
  );
};

export default EventNavigation;
