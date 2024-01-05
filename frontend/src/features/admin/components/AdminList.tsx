import React, { useState } from 'react';
import { styles } from '../../../styles';

const AdminList = () => {
    const [admins, setAdmins] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);

    const fetchAdmins = () => {
        setIsLoading(true);
        fetch('http://localhost:3000/admin/list', {
            credentials: 'include',
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                setAdmins(data);
                setIsLoaded(true);
            })
            .catch(err => {
                setError(err.message);
            })
            .finally(() => {
                setIsLoading(false);
            });
    };

    return (
        <div className="container mx-auto p-6 bg-white rounded-lg shadow-md max-w-4xl">
            <div className='text-center'>
                <h2 className="text-2xl font-semibold text-gray-700 mb-4">Admin Management</h2>
            
                <button 
                    className={`${styles.button} ${isLoading ? 'opacity-50 cursor-not-allowed' : ''}`}
                    onClick={fetchAdmins}
                    disabled={isLoading}
                >
                    {isLoading ? 'Loading...' : 'View Admins'}
                </button>
            </div>
            
            {error && <p className="mt-3 text-red-500 font-semibold">Error: {error}</p>}

            {isLoaded && (
                <div className="overflow-x-auto mt-6">
                    <table className="w-full text-sm text-left text-gray-500">
                        <thead className="text-xs text-gray-700 uppercase bg-gray-50">
                            <tr>
                                <th className="px-6 py-3">Name</th>
                                <th className="px-6 py-3">Email</th>
                                <th className="px-6 py-3">Position</th>
                                <th className="px-6 py-3">Status</th>
                            </tr>
                        </thead>
                        <tbody>
                            {admins.map((admin) => (
                                <tr className="bg-white border-b hover:bg-gray-50" key={admin.id}>
                                    <td className="px-6 py-4">{admin.name}</td>
                                    <td className="px-6 py-4">{admin.email}</td>
                                    <td className="px-6 py-4">{admin.position}</td>
                                    <td className="px-6 py-4">{admin.status}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
};

export default AdminList;