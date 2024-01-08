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
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => {
        setAdmins(data);
        setIsLoaded(true);
      })
      .catch((err) => {
        setError(err.message);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <div className="container mx-auto mb-5 max-w-4xl rounded-lg bg-white p-6 shadow-md">
      <div className="text-center">
        <h2 className="mb-4 text-2xl font-semibold text-gray-700">
          Admin Management
        </h2>

        <button
          className={`${styles.button} ${
            isLoading ? 'cursor-not-allowed opacity-50' : ''
          }`}
          onClick={fetchAdmins}
          disabled={isLoading}
        >
          {isLoading ? 'Loading...' : 'View Admins'}
        </button>
      </div>

      {error && (
        <p className="mt-3 font-semibold text-red-500">Error: {error}</p>
      )}

      {isLoaded && (
        <div className="mt-6 overflow-x-auto">
          <table className="w-full text-left text-sm text-gray-500">
            <thead className="bg-gray-50 text-xs uppercase text-gray-700">
              <tr>
                <th className="px-6 py-3">Name</th>
                <th className="px-6 py-3">Email</th>
                <th className="px-6 py-3">Position</th>
                <th className="px-6 py-3">Status</th>
              </tr>
            </thead>
            <tbody>
              {admins.map((admin) => (
                <tr
                  className="border-b bg-white hover:bg-gray-50"
                  key={admin.id}
                >
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
