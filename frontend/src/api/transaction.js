import { apiService } from './apiService';

export const createGoodsReceipt = (data) => {
  return apiService({
    url: `/goods-receipt`,
    method: 'POST',
    data: data,
  });
};

export const createGoodsIssue = (data) => {
  return apiService({
    url: `/goods-issue`,
    method: 'POST',
    data: data,
  });
};

export const getAllGoodsReceipt = () => {
  return apiService({
    url: `/goods-receipt`,
    method: 'GET',
    data: null,
  });
};

export const getAllGoodsIssue = () => {
  return apiService({
    url: `/goods-issue`,
    method: 'GET',
    data: null,
  });
};
export const getProductReport = () => {
  return apiService({
    url: `/stock-report`,
    method: 'GET',
    data: null,
  });
};
