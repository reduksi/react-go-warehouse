import axios from 'axios'
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

const instance = axios.create({
  baseURL: API_URL,
  timeout: 1000,
})


export const apiService = async (Params) => {
  const { url, method, data, params, multipart = false, headers } = Params
  const service = await instance({
    url: url,
    method: method,
    data: data,
    params: params,
    timeout: 60000,
    multipart,
    headers,
  })
  return service
}
