import React from 'react';
import EventGrid from '../components/EventGrid';
function Events() {
  return (
    <div className="mx-auto max-w-[1920px] ">
      <div className="flex flex-col items-center">
        <EventGrid />
      </div>
    </div>
  );
}

export default Events;
