// AdminForm.tsx
import React, { useState } from 'react';
import { styles } from '../../../styles';

type Admin = {
    name: string;
    email: string;
    position: string;
    status: string;
};

function AdminForm() {
    const [admin, setAdmin] = useState<Admin>({ name: '', email: '', position: '', status: '' });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setAdmin({ ...admin, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:3000/admin/create', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(admin),
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
        <form onSubmit={handleSubmit} 
        className="flex flex-col m-8 space-y-4">
            <input
                type="text"
                name="name"
                placeholder="Name"
                value={admin.name}
                onChange={handleChange}
                className="p-2 border rounded"
            />
            <input
                type="email"
                name="email"
                placeholder="Email"
                value={admin.email}
                onChange={handleChange}
                className="p-2 border rounded"
            />
            <input
                type="text"
                name="position"
                placeholder="Position"
                value={admin.position}
                onChange={handleChange}
                className="p-2 border rounded"
            />
            <input
                type="text"
                name="status"
                placeholder="Status"
                value={admin.status}
                onChange={handleChange}
                className="p-2 border rounded"
            />
            <button type="submit"
            className={`${styles.button}`}> 
                Create Admin
            </button>
        </form>
    );
}

export default AdminForm