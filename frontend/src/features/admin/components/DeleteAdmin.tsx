import React, { useState } from 'react';
import { Button, Modal } from '../../../components';

function DeleteAdmin() {
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  const [showModal, setShowModal] = useState(false); // State for modal visibility
  const [modalMessage, setModalMessage] = useState(''); // State for modal message

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
    if (error) {
      setError('');
    }
  };

  const handleModalClose = () => {
    setShowModal(false);
    setModalMessage('');
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email) {
      setError('Email is required');
      return;
    }

    const encodedEmail = encodeURIComponent(email);

    try {
      const response = await fetch(`http://localhost:3000/api/admins/${encodedEmail}`, {
        method: 'DELETE',
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      console.log('Admin deleted successfully');
      setEmail(''); // Clear the form
      setModalMessage('Admin deleted successfully'); // Set success message
      setShowModal(true); // Show the modal
    } catch (error) {
      console.error('Error deleting admin:', error);
      setError('Failed to delete admin. Please try again.');
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit} className="m-8 flex flex-col space-y-4">
        <input
          type="email"
          name="email"
          placeholder="Admin's Email"
          value={email}
          onChange={handleChange}
          className={`rounded border p-2 ${error ? 'border-red-500' : ''}`}
        />
        {error && <p className="text-red-500">{error}</p>}

        <Button buttonText="Delete Admin" type="submit" additionalClasses="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" />
      </form>
      {showModal && (
        <Modal
          msg={modalMessage}
          onClose={handleModalClose}
        />
      )}
    </div>
  );
}

export default DeleteAdmin;
