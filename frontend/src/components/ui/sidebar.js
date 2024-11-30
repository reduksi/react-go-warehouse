import React from 'react';
import { Link } from 'react-router-dom';
import {
  FaTachometerAlt,
  FaBox,
  FaArrowCircleUp,
  FaArrowCircleDown,
  FaUserTie,
  FaWarehouse,
  FaUsers,
} from 'react-icons/fa';

const Sidebar = () => {
  return (
    <div className="flex flex-col w-64 bg-white rounded-r-3xl overflow-hidden h-[100vh]">
      <div className="flex items-center justify-center h-20 shadow-md">
        <h1 className="text-3xl uppercase text-indigo-500">SAMB</h1>
      </div>
      <ul className="flex flex-col py-4">
        <li>
          <Link
            to="/"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaTachometerAlt />
            </span>
            <span className="text-sm font-medium">Dashboard</span>
          </Link>
        </li>
        <li>
          <Link
            to="/goods-receipt"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaArrowCircleUp />
            </span>
            <span className="text-sm font-medium">Goods Receipt</span>
          </Link>
        </li>
        <li>
          <Link
            to="/goods-issue"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaArrowCircleDown />
            </span>
            <span className="text-sm font-medium">Goods Issue</span>
          </Link>
        </li>
        <li>
          <Link
            to="/suppliers"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaUserTie />
            </span>
            <span className="text-sm font-medium">Suppliers</span>
          </Link>
        </li>
        <li>
          <Link
            to="/products"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaBox />
            </span>
            <span className="text-sm font-medium">Products</span>
          </Link>
        </li>
        {/* Add the Warehouses Link */}
        <li>
          <Link
            to="/warehouses"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaWarehouse />
            </span>
            <span className="text-sm font-medium">Warehouses</span>
          </Link>
        </li>
        {/* Add the Customers Link */}
        <li>
          <Link
            to="/customers"
            className="flex flex-row items-center h-12 transform hover:translate-x-4 transition-transform ease-in duration-200 text-gray-500 hover:text-gray-800"
          >
            <span className="inline-flex items-center justify-center h-12 w-12 text-lg text-gray-400">
              <FaUsers />
            </span>
            <span className="text-sm font-medium">Customers</span>
          </Link>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
