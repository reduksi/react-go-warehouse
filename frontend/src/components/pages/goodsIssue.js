import React, { useEffect, useState } from 'react';
import InputText from '../global/Input';
import DateInput from '../global/dateInput';
import Select from '../global/select';
import NestedTable from '../global/nestedTable';
import { FaTrash } from 'react-icons/fa';
import {
  getAllProducts,
  getAllSuppliers,
  getAllWarehouses,
} from '../../api/master';
import { createGoodsIssue, getAllGoodsIssue } from '../../api/transaction';
import Swal from 'sweetalert2';

const Goods = () => {
  const [form, setForm] = useState({
    details: [{ product: '', qty_dus: '', qty_pcs: '' }],
  });
  const [products, setProducts] = useState([]);
  const [suppliers, setSuppliers] = useState([]);
  const [warehouses, setWarehouses] = useState([]);
  const [transactions, setTransactions] = useState([]);

  async function getTransactions() {
    try {
      const { data } = await getAllGoodsIssue();
      setTransactions(data.reverse() || []);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    document.title = 'Goods Receipt';

    async function getProducts() {
      try {
        const { data } = await getAllProducts();
        setProducts(data || []);
      } catch (error) {
        console.log(error);
      }
    }

    async function getSuppliers() {
      try {
        const { data } = await getAllSuppliers();
        setSuppliers(data || []);
      } catch (error) {
        console.log(error);
      }
    }

    async function getWarehouses() {
      try {
        const { data } = await getAllWarehouses();
        setWarehouses(data || []);
      } catch (error) {
        console.log(error);
      }
    }
    getTransactions();
    getProducts();
    getSuppliers();
    getWarehouses();
  }, []);

  function onChange(e) {
    const { name, value } = e.target;

    if (name.includes('-')) {
      const [field, index] = name.split('-');
      setForm((prevState) => {
        const updatedDetails = [...prevState.details];
        updatedDetails[index][field] = value;
        return {
          ...prevState,
          details: updatedDetails,
        };
      });
    } else {
      setForm((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }
  }

  function addDetail() {
    setForm((prevState) => ({
      ...prevState,
      details: [
        ...prevState.details,
        { product: '', qty_dus: '', qty_pcs: '' },
      ],
    }));
  }

  function deleteDetail(index) {
    setForm((prevState) => {
      const updatedDetails = prevState.details.filter((_, i) => i !== index);
      return {
        ...prevState,
        details: updatedDetails,
      };
    });
  }

  async function onSubmit(e) {
    e.preventDefault();

    const payload = {
      Header: {
        TrxOutNo: form.transactionNo,
        TrxOutNotes: form.notes,
        TrxOutDate: new Date(form.transactionDate),
        WhsIdf: warehouses.find((x) => x.whsName === form.warehouse)?.whsPK,
        TrxOutSuppIdf: suppliers.find((x) => x.supplierName === form.supplier)
          ?.supplierPK,
      },
      Details: form.details.map((detail) => ({
        TrxOutDQtyDus: Number(detail.qty_dus),
        TrxOutDQtyPcs: Number(detail.qty_pcs),
        TrxOutDProductIdf: products.find(
          (x) => x.productName === detail.product
        )?.productPK,
      })),
    };

    try {
      await createGoodsIssue(payload);
      Swal.fire({
        icon: 'success',
        title: 'Out Transaction has been saved',
        showConfirmButton: false,
      });
      getTransactions();
      setForm({
        transactionNo: '',
        transactionDate: '',
        product: '',
        supplier: '',
        warehouse: '',
        qty_dus: '',
        qty_pcs: '',
        notes: '',
        details: [{ product: '', qty_dus: '', qty_pcs: '' }],
      });
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <div className="w-3/4">
      <h2 className="text-3xl font-bold mb-4">Goods Issue</h2>
      <p>Goods Out Transaction</p>
      <form className="mt-6" onSubmit={onSubmit}>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-4 my-4">
          <InputText
            label="Transaction in No"
            name="transactionNo"
            onChange={onChange}
            value={form.transactionNo}
          />
          <DateInput
            label="Transaction in Date"
            name="transactionDate"
            onChange={onChange}
            value={form.transactionDate}
          />
        </div>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-4 my-4">
          <Select
            label="Supplier"
            name="supplier"
            options={suppliers.map((x) => x.supplierName)}
            onChange={onChange}
            value={form.supplier}
          />
          <Select
            label="Warehouse"
            name="warehouse"
            options={warehouses.map((x) => x.whsName)}
            onChange={onChange}
            value={form.warehouse}
          />
        </div>
        <div className="grid grid-cols-1 gap-4 my-4">
          <InputText
            label="Transaction Notes"
            name="notes"
            onChange={onChange}
            value={form.notes}
          />
        </div>

        <div className="my-4">
          <h3 className="font-bold">Products</h3>
          {form.details.map((detail, index) => (
            <div key={index} className="relative my-4 pb-4">
              <div className="grid grid-cols-1 gap-4 my-4">
                <Select
                  label="Product"
                  name={`product-${index}`}
                  options={products.map((x) => x.productName)}
                  onChange={onChange}
                  value={detail.product}
                />
              </div>
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-4 my-4">
                <InputText
                  label="Quantity (box)"
                  type="number"
                  name={`qty_dus-${index}`}
                  onChange={onChange}
                  value={detail.qty_dus}
                />
                <InputText
                  label="Quantity (pcs)"
                  type="number"
                  name={`qty_pcs-${index}`}
                  onChange={onChange}
                  value={detail.qty_pcs}
                />
              </div>
              {form.details.length !== 1 && (
                <div className="absolute -top-3 -right-3">
                  <button
                    type="button"
                    onClick={() => deleteDetail(index)}
                    className="bg-transparent border-2 bg-white border-red-500 text-red-500 w-7 h-7 flex items-center justify-center rounded-full text-md hover:bg-red-500 hover:text-white"
                  >
                    <FaTrash />
                  </button>
                </div>
              )}
            </div>
          ))}

          <div className="flex justify-end -mt-8 -mr-3">
            <button
              type="button"
              onClick={addDetail}
              className="bg-transparent border-2 bg-white border-indigo-500 text-indigo-500 w-7 h-7 flex items-center justify-center rounded-full text-lg hover:bg-blue-500 hover:text-white"
            >
              +
            </button>
          </div>
        </div>
        <div className="flex justify-center mt-4">
          <button
            type="submit"
            className="w-32 bg-white tracking-wide text-gray-800 font-bold rounded border-b-2 border-indigo-500 hover:border-indigo-600 hover:bg-indigo-500 hover:text-white shadow-md py-2 px-6 inline-flex items-center"
          >
            <span className="mx-auto">Submit</span>
          </button>
        </div>
      </form>
      <NestedTable
        data={transactions}
        products={products}
        suppliers={suppliers}
        warehouses={warehouses}
      />
    </div>
  );
};

export default Goods;
