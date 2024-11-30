import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
// import GoodsIssueForm from './components/forms/GoodsIssueForm';
// import SupplierList from './components/lists/SupplierList';
// import ProductList from './components/lists/ProductList';
import Sidebar from './components/ui/sidebar';
import Navbar from './components/ui/navbar';
import Dashboard from './components/pages/dashboard';
import GoodsReceiptPage from './components/pages/goodsReceipt';
import GoodsIssuePage from './components/pages/goodsIssue';
import SuppliersPage from './components/pages/suppliers';
import ProductsPage from './components/pages/products';
import WarehousePage from './components/pages/warehouse';
import CustomerPage from './components/pages/customers';

const App = () => {
  return (
    <Router>
      <div className="flex bg-gray-100">
        {/* Sidebar */}
        <Sidebar />
        <div className="flex-1 p-5 ml-4 bg-white rounded-l-3xl">
          <Navbar />
          <div className="p-10 flex justify-center">
            <Routes>
              <Route path="/" element={<Dashboard />} />
              <Route path="/goods-receipt" element={<GoodsReceiptPage />} />
              <Route path="/goods-issue" element={<GoodsIssuePage />} />
              <Route path="/suppliers" element={<SuppliersPage />} />
              <Route path="/products" element={<ProductsPage />} />
              <Route path="/warehouses" element={<WarehousePage />} />
              <Route path="/customers" element={<CustomerPage />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
};

export default App;
