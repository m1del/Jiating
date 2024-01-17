import React, { useEffect, useState } from 'react';
import { Button } from '../../../components';

type Admin = {
  id: string;
  name: string;
  email: string;
  position: string;
  status: string;
}

const AdminList: React.FC = () => {
  const [admins, setAdmins] = useState<Admin[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const pageSize = 5; // idk putting a default here, may reduce to fit

  const fetchAdmins = async(page: number) => {
    setIsLoading(true);
    try {
      const resp = await fetch(`http://localhost:3000/api/admins?page=${page}&pageSize=${pageSize}`, {
        credentials: 'include',
      });
      if(!resp.ok) {
        throw new Error(`http error! status: ${resp.status}`)
      }
      const data = await resp.json();
      setAdmins(data.admins);
      const totalAdmins = data.totalCount;
      const totalPages = Math.ceil(totalAdmins/pageSize);
      setTotalPages(totalPages);
    } catch (err) {
      setError('oops! try again in a few seconds')
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchAdmins(currentPage);
  }, [currentPage]);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const refreshData = () => {
    fetchAdmins(currentPage);
  }

  return(
    <div className='container m-auto p-6 rounded-lg bg-white shadow-md'>
      <div className='text-center'>
        <h2 className='mb-4 text-2xl font-semibold text-gray-700'>
          Admin Management
        </h2>
      <Button buttonText='Refresh' type='submit'onClick={refreshData}/>
      </div>

      {error && (
        <p className='mt-3 font-semibold text-red-500'>Error: {error}</p>
      )}

      {isLoading && <p>loading...</p>}
      
      {!isLoading && (
        <div className='mt-6 overflow-x-auto'>
          <table className='w-full text-left text-sm text-gray-500'>
            <thead className='bg-gray-50 text-xs uppercase text-gray-700'>
              <tr>
                <th className='px-6 py-3'>Name</th>
                <th className='px-6 py-3'>Email</th>
                <th className='px-6 py-3'>Position</th>
                <th className='px-6 py-3'>Status</th>
              </tr>
            </thead>
            <tbody>
              {admins.map((admin) => (
                <tr
                  className='border-b bg-white hover:bg-gray-50'
                  key={admin.id}
                >
                  <td className='px-6 py-4'>{admin.name}</td>
                  <td className='px-6 py-4'>{admin.email}</td>
                  <td className='px-6 py-4'>{admin.position}</td>
                  <td className='px-6 py-4'>{admin.status}</td>
                </tr>
              ))}
            </tbody>
          </table>
          </div>
      )}

      <div className="mt-4 flex justify-between items-center">
        <button
          className="px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-md"
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 1}
        >
          Previous
        </button>
        <span>
          Page {currentPage} of {totalPages}
        </span>
        <button
          className="px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-md"
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage === totalPages}
        >
          Next
        </button>
      </div>

    </div>
  );
};

export default AdminList;
