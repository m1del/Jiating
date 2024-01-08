import React from 'react';
import { styles } from '../../../styles';
export default function CreateEventButton() {
  const redirect = () => {
    window.location.href = '/admin/eventform';
  };
  return (
    <button onClick={redirect} className={`${styles.button}`}>
      {' '}
      Add Event{' '}
    </button>
  );
}
