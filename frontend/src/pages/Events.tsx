import React from 'react';
import EventGrid from '../components/events/EventGrid';
function Events() {
  return (
    <div className="container mx-auto w-full ">
      <div className="flex flex-col items-center">
        <div className="w-full rounded-md bg-gray-700 p-6 text-white shadow-md">
          <h1 className="mb-2 text-4xl font-bold">Media</h1>
          <p className="text-xl">Explore our collection of photos and events</p>
        </div>
        <EventGrid />
      </div>
    </div>
  );
}

export default Events;
