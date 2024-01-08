import React, { FormEvent, useState } from 'react';
import { Modal } from '../../components';
import { styles } from '../../styles';

type FormData = {
  name: string;
  email: string;
  subject: string;
  message: string;
};


const EmailForm = () => {
  const [formData, setFormData] = useState<FormData>({ name: '', email: '', subject: '', message: '' });
  const [modalMsg, setModalMsg] = useState<string>('');
  const [showModal, setShowModal] = useState<boolean>(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const resetForm = () => {
    setFormData({ name: '', email: '', subject: '', message: '' });
  };

  const closeModal = () => {
    setShowModal(false);
    setModalMsg('');
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setModalMsg('');
    setShowModal(false);

    // email Validation
    if (!validateEmail(formData.email)) {
      alert('Please enter a valid email address.');
      return;
    }

    // text Fields Validation
    if (!sanitizeInput(formData.name) || !sanitizeInput(formData.subject) || !sanitizeInput(formData.message)) {
      alert('Invalid characters in the input fields.');
      return;
    }

    // length Checks
    if (formData.name.length > 100 || formData.subject.length > 150 || formData.message.length > 1000) {
      alert('Input is too long.');
      return;
    }

    try {
      const resp = await fetch('http://localhost:3000/api/send-email', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData),
      });
      if (!resp.ok) {
        throw new Error('Failed to send email');
      }
      setModalMsg('Email sent successfully!');
      setShowModal(true);
      resetForm();
    } catch (err) {
      console.error(err);
      setModalMsg('An error occurred :(');
      setShowModal(true);
    }
  };

  const validateEmail = (email: string) => {
    const re = /\S+@\S+\.\S+/;
    return re.test(email);
  };

  const sanitizeInput = (input: string) => {
    const re = /[<>]/;
    return !re.test(input);
  }; 

  return (
    <div className="container mx-auto px-4">
      <form onSubmit={handleSubmit} 
        className="space-y-6 bg-white p-6 rounded-lg shadow-md text-2xl">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="flex-1">
            <label className="block font-medium text-gray-700">
              Name
              <input
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                required
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                focus:outline-none focus:ring-cyan-700 focus:border-cyan-700"
              />
            </label>
          </div>
          <div className="flex-1">
            <label className="block font-medium text-gray-700">
              Email
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                required
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                focus:outline-none focus:ring-cyan-700-500 focus:border-cyan-700"
              />
            </label>
          </div>
        </div>

        <div>
          <label className="block font-medium text-gray-700">
            Subject
            <input
              type="text"
              name="subject"
              value={formData.subject}
              onChange={handleChange}
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
              focus:outline-none focus:ring-cyan-700 focus:border-cyan-700"
            />
          </label>
        </div>

        <div>
          <label className="block font-medium text-gray-700">
            Message
            <textarea
              name="message"
              value={formData.message}
              onChange={handleChange}
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
              focus:outline-none focus:ring-cyan-700 focus:border-cyan-700"
              rows={4}
            />
          </label>
        </div>

        <button
          type="submit"
          className={`${styles.button}`}
        >
          Get in touch
        </button>
      </form>
      {showModal && <Modal msg={modalMsg} onClose={closeModal} />}
    </div>
  );
};

export default EmailForm;
