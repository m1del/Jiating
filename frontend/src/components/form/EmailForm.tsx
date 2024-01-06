import React, { FormEvent, useState } from 'react';
import { styles } from '../../styles';

type FormData = {
  name: string;
  email: string;
  subject: string;
  message: string;
};

const EmailForm = () => {
  const [formData, setFormData] = useState<FormData>({ name: '', email: '', subject: '', message: '' });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

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
      //TODO: handle sucess -> clear form & show sucess message
    } catch (err) {
      console.error(err);
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
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="flex flex-col md:flex-row md:space-x-4">
          <div className="flex-1">
            <label className="block text-sm font-medium text-gray-700">
              Name
              <input
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                required
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </label>
          </div>
          <div className="flex-1">
            <label className="block text-sm font-medium text-gray-700">
              Email
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                required
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
                focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </label>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">
            Subject
            <input
              type="text"
              name="subject"
              value={formData.subject}
              onChange={handleChange}
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
              focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
          </label>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">
            Message
            <textarea
              name="message"
              value={formData.message}
              onChange={handleChange}
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm 
              focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              rows={4}
            />
          </label>
        </div>

        <button
          type="submit"
          className={`${styles.button}`}
        >
          Send
        </button>
      </form>
    </div>
  );
};

export default EmailForm;
