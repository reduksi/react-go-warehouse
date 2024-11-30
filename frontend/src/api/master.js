import { apiService } from './apiService';

export const getAllProducts = () => {
  return apiService({
    url: `/products`,
    method: 'GET',
    data: null,
  });
};
export const getAllSuppliers = () => {
  return apiService({
    url: `/suppliers`,
    method: 'GET',
    data: null,
  });
};
export const getAllCustomers = () => {
  return apiService({
    url: `/customers`,
    method: 'GET',
    data: null,
  });
};
export const getAllWarehouses = () => {
  return apiService({
    url: `/warehouses`,
    method: 'GET',
    data: null,
  });
};

export const createProducts = (data) => {
  return apiService({
    url: `/products`,
    method: 'POST',
    data: data,
  });
};
export const createSuppliers = (data) => {
  return apiService({
    url: `/suppliers`,
    method: 'POST',
    data: data,
  });
};
export const createCustomers = (data) => {
  return apiService({
    url: `/customers`,
    method: 'POST',
    data: data,
  });
};
export const createWarehouses = (data) => {
  return apiService({
    url: `/warehouses`,
    method: 'POST',
    data: data,
  });
};
