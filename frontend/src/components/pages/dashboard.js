import React, { useEffect, useState } from 'react';
import { getProductReport } from '../../api/transaction';

const Dashboard = () => {
  const [reports, setReports] = useState([]);
  useEffect(() => {
    document.title = 'Warehouse Dashboard';

    async function getReport() {
      try {
        const { data } = await getProductReport();
        setReports(data || []);
      } catch (error) {
        console.log(error);
      }
    }

    getReport();
  }, []);

  return (
    <div className="w-3/4">
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

      <h2 className="text-3xl font-bold mb-4 mt-10">Product Stock</h2>
      <div className="overflow-x-auto mt-5">
        <table className="w-full text-sm text-left text-gray-500">
          <thead className="text-xs text-gray-700 uppercase bg-gray-300">
            <tr>
              <th scope="col" className="px-6 py-3">
                Id
              </th>
              <th scope="col" className="px-6 py-3">
                Product Name
              </th>
              <th scope="col" className="px-6 py-3">
                Current Stock in Dus
              </th>
              <th scope="col" className="px-6 py-3">
                Current Stock in Pcs
              </th>
            </tr>
          </thead>
          <tbody>
            {reports.map((supp) => (
              <tr className="bg-white border-b">
                <th
                  scope="row"
                  className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                >
                  {supp.productPK}
                </th>
                <td className="px-6 py-4">{supp.productName}</td>
                <td className="px-6 py-4">{supp.stockInDus}</td>
                <td className="px-6 py-4">{supp.stockInPcs}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Dashboard;
