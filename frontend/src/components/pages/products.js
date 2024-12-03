import React, { useEffect, useState } from 'react';
import InputText from '../global/Input';
import { getAllProducts, createProducts } from '../../api/master';
import Swal from 'sweetalert2';

const ProductPage = () => {
  const [products, setProducts] = useState([]);
  const [newValue, setNewValue] = useState('');
  async function getProducts() {
    try {
      const { data } = await getAllProducts();
      setProducts(data || []);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    document.title = 'Product Table';
    getProducts();
  }, []);

  async function onSubmit(e) {
    e.preventDefault();

    try {
      await createProducts({ productName: newValue });
      setNewValue('');
      Swal.fire({
        icon: 'success',
        title: 'New Product has been added',
        showConfirmButton: false,
      });
      getProducts();
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <div className="w-3/4">
      <h2 className="text-3xl font-bold mb-4">Product</h2>
      <p className="mb-4">List of Products!</p>

      <form className="mt-6" onSubmit={onSubmit}>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-4 my-4">
          <InputText
            label="Product"
            value={newValue}
            onChange={(e) => {
              setNewValue(e.target.value);
            }}
            required
          />
          <button
            type="submit"
            className="w-24 bg-white tracking-wide text-gray-800 items-center font-bold rounded border-b-2 border-indigo-500 hover:border-indigo-600 hover:bg-indigo-500 hover:text-white shadow-md py-2 px-6"
          >
            <span className="mx-auto">Add</span>
          </button>
        </div>
      </form>

      <div className="overflow-x-auto">
        <table className="w-full text-sm text-left text-gray-500">
          <thead className="text-xs text-gray-700 uppercase bg-gray-300">
            <tr>
              <th scope="col" className="px-6 py-3">
                Id
              </th>
              <th scope="col" className="px-6 py-3">
                Product Name
              </th>
            </tr>
          </thead>
          <tbody>
            {products.map((supp) => (
              <tr className="bg-white border-b">
                <th
                  scope="row"
                  className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                >
                  {supp.productPK}
                </th>
                <td className="px-6 py-4">{supp.productName}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ProductPage;
