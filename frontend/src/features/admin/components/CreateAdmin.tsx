import React, { useState } from 'react';
import { Button } from '../../../components';
import { styles } from '../../../styles';

type Admin = {
  name: string;
  email: string;
  position: string;
  status: string;
};

function CreateAdmin() {
  const [admin, setAdmin] = useState<Admin>({
    name: '',
    email: '',
    position: '',
    status: 'Active',
  });
  const [errors, setErrors] = useState<Admin>({
    name: '',
    email: '',
    position: '',
    status: '',
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    setAdmin({ ...admin, [e.target.name]: e.target.value });

    // clear warning message
    if (errors[e.target.name]) {
      setErrors({ ...errors, [e.target.name]: '' });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const newErrors = { name: '', email: '', position: '', status: '' };
    let isValid = true;
    // validate form
    Object.keys(admin).forEach((key) => {
      if (!admin[key]) {
        newErrors[key] = `${key} is required`;
        isValid = false;
      }
    });

    setErrors(newErrors);
    if (!isValid) {
      return;
    }

    try {
      const response = await fetch('http://localhost:3000/api/admins', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(admin),
        credentials: 'include',
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      console.log('Admin created successfully');
    } catch (error) {
      console.error('Error creating admin:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="m-8 flex flex-col space-y-4">
      <input
        type="text"
        name="name"
        placeholder="Name"
        value={admin.name}
        onChange={handleChange}
        className={`rounded border p-2 ${errors.name ? 'border-red-500' : ''}`}
      />
      {errors.name && <p className="text-red-500">{errors.name}</p>}

      <input
        type="email"
        name="email"
        placeholder="Email"
        value={admin.email}
        onChange={handleChange}
        className={`rounded border p-2 ${errors.email ? 'border-red-500' : ''}`}
      />
      {errors.email && <p className="text-red-500">{errors.email}</p>}

      <input
        type="text"
        name="position"
        placeholder="Position"
        value={admin.position}
        onChange={handleChange}
        className={`rounded border p-2 ${
          errors.position ? 'border-red-500' : ''
        }`}
      />
      {errors.position && <p className="text-red-500">{errors.position}</p>}
      <div className="relative inline-block w-full text-gray-700">
        <select
          name="status"
          value={admin.status}
          onChange={handleChange}
          className="focus:shadow-outline h-10 w-full appearance-none rounded-lg border pl-3 pr-6 text-base placeholder-gray-600"
        >
          <option value="Active">Active</option>
          <option value="Inactive">Inactive</option>
          <option value="Hiatus">Hiatus</option>
        </select>
        <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2">
          <svg
            className="h-4 w-4 fill-current"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          >
            <path d="M5.516 7.548c0.436-0.446 1.043-0.481 1.576 0l3.908 3.747 3.908-3.747c0.533-0.481 1.141-0.446 1.576 0 0.436 0.445 0.408 1.197 0 1.615l-4.695 4.502c-0.217 0.223-0.502 0.335-0.789 0.335s-0.571-0.112-0.789-0.335l-4.695-4.502c-0.408-0.418-0.436-1.17 0-1.615z" />
          </svg>
        </div>
      </div>

      
      <Button buttonText="Create Admin" type="submit" additionalClasses={styles.button} />
    </form>
  );
}

export default CreateAdmin;
