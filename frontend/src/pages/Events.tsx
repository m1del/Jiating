import React from 'react';
import EventGrid from '../components/events/EventGrid';
function Events() {
  return (
    <div className="mx-auto max-w-[1920px] ">
      <div className="flex flex-col items-center">
        <h2 className="items-center text-3xl text-cyan-600">Events</h2>
        <EventGrid />
      </div>
    </div>
  );
}

export default Events;
