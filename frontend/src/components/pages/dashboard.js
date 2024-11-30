import React, { useEffect } from 'react';

const Dashboard = () => {
  useEffect(() => {
    document.title = "Warehouse Dashboard";
  }, []);

  return (
    <div className='w-3/4'>
      <h2 className="text-3xl font-bold mb-4">Dashboard</h2>
      <p>Welcome to the warehouse management system!</p>
      <div className="mt-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          <div className="p-4 border shadow-md rounded-lg bg-gray-100">
            <h3 className="font-bold text-xl">Goods Receipt</h3>
            <p>Manage incoming products</p>
          </div>
          <div className="p-4 border shadow-md rounded-lg bg-gray-100">
            <h3 className="font-bold text-xl">Goods Issue</h3>
            <p>Manage outgoing products</p>
          </div>
          <div className="p-4 border shadow-md rounded-lg bg-gray-100">
            <h3 className="font-bold text-xl">Inventory</h3>
            <p>Manage products and suppliers</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;