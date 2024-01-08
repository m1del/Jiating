import React from 'react';
import { EventData } from '../events/EventModel';

type SetEventData = (eventData: EventData) => void;

const GetEvents = (setEventData: SetEventData) => {
  const search = window.location.search;
  const params = new URLSearchParams(search);
  const eventID = params.get('id');
  if (eventID === null) return;

  const fetchData = async () => {
    const getterUrl = `http://localhost:3000/admin/get-event?id=${eventID}`;
    try {
      const resp = await fetch(getterUrl, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });
      if (!resp.ok) {
        throw new Error('Failed to get event');
      } else {
        const data = await resp.json();
        setEventData(data);
      }
    } catch (err) {
      console.error(err);
    }
  };

  fetchData();
};

export default GetEvents;
