import { apiService } from './apiService'

export const createGoodsReceipt = (data) => {
  return apiService({
    url: `/goods-receipt`,
    method: 'POST',
    data: data,
  })
}

export const createGoodsIssue = (data) => {
  return apiService({
    url: `/goods-issue`,
    method: 'POST',
    data: data,
  })
}

