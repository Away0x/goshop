import { AxiosRequestConfig } from 'axios';

export const URLENCODED_CONTENT_TYPE = {
  'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
};
export const JSON_CONTENT_TYPE = {
  'Content-Type': 'application/json;charset=UTF-8',
};
export const CommonErrorMessage = '网络错误';

export interface MyAxiosRequestConfig extends AxiosRequestConfig {
  json?: boolean     // 设置请求头为 JSON_CONTENT_TYPE
  showErrorMessage?: boolean // 是否弹窗错误 message
}

/** 后端响应的数据格式 */
export interface APIResponseData<T = any> {
  code: number
  data: T
  msg: string
  token: string | null
}

/** services 整理后的数据格式 */
export interface CommonResponse<T = any> {
  status: boolean
  message: string
  data: T
}