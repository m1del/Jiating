import React, { useState, useEffect } from 'react';
import { v4 as uuidv4 } from 'uuid';
import { EventData } from './EventModel';
import GetEvents from './GetEvents';
import { styles } from '../../styles';

export default function Event() {
  const [eventData, setEventData] = useState<EventData>({
    id: '',
    created_at: null,
    updated_at: null,
    admin_id: 1,
    event_name: '',
    date: '',
    description: '',
    content: '',
    is_draft: false,
    published_at: null,
    image_id: 0,
  });

  useEffect(() => {
    GetEvents(setEventData);
  }, [setEventData]);

  const editPost = async () => {
    const editUrl = `/admin/eventform?id=${eventData.id}`;
    window.location.href = editUrl;
  };

  return (
    <div className="mt-8 flex h-screen flex-col items-center">
      {eventData.is_draft && (
        <button className={`${styles.button}`} onClick={editPost}>
          Edit Post
        </button>
      )}
      <img src="/trad.jpg" className="mt-5 h-[50%] w-[80%] object-cover" />
      <div className=" flex w-[60%] flex-col text-center">
        <h3 className="mt-5 text-3xl font-semibold">{eventData.event_name}</h3>
        <p className="mt-3 text-lg leading-relaxed">{eventData.content}</p>
      </div>
    </div>
  );
}
