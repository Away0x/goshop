import Axios, { AxiosResponse } from 'axios';
import qs from 'qs';
import { message } from 'antd';

import {
  MyAxiosRequestConfig,
  URLENCODED_CONTENT_TYPE,
  JSON_CONTENT_TYPE,
  APIResponseData,
  CommonResponse,
  CommonErrorMessage,
} from './type';
import {
  OK_CODE,
  NO_LOGIN_CODE,
  OTHER_ERROR_CODES,
} from './code';
import Token from '@/auth/token';

const axios = Axios.create({
  headers: URLENCODED_CONTENT_TYPE,
});

// request
axios.interceptors.request.use(
  (config: MyAxiosRequestConfig) => {
    // 处理参数
    if (!config.params) { config.params = {}; }
    if (!config.data) { config.data = {}; }

    const token = Token.get();
    if (token) {
      config.params.token = token; // query 上携带 token
    }

    if (!config.json) {
      config.data = qs.stringify(config.data);
    } else {
      config.headers = Object.assign({}, (config.headers || {}), JSON_CONTENT_TYPE);
    }
    
    return config;
  },
  (error: any) => {
    message.error(error);
    return Promise.reject(error);
  },
);

// response
axios.interceptors.response.use(
  (response: AxiosResponse<APIResponseData>) => {
    // 响应数据具体在 services 中处理
    return response;
  },
  (error: any) => {
    message.error(error);
    return Promise.reject(error);
  },
);

class Services {
  private static _instance: Services | null = null;

  public static instance(): Services {
    if (this._instance === null) {
      this._instance = new this();
    }

    return this._instance!;
  }

  private errorResponseData(config: MyAxiosRequestConfig, result?: APIResponseData): CommonResponse<any> {
    const msg = (result && result.msg) ? result.msg : CommonErrorMessage;

    if (config.showErrorMessage) {
      message.error(msg);
    }

    return {
      status: false,
      message: msg,
      data: result ? result.data : null,
    } 
  }

  private responseData<T>(config: MyAxiosRequestConfig, result: APIResponseData<T>): CommonResponse<T> {
    const statusCode = Number(result.code);

    // 如有携带 token，存储
    if (result.token) {
      Token.set(result.token);
    }

    // 请求成功
    if (statusCode === OK_CODE) {
      return {
        status: true,
        message: result.msg || '成功',
        data: result.data,
      };
    }

    // 请求失败 (code 处理)
    if (statusCode === NO_LOGIN_CODE) {
      Token.logout({
        triggerGlobalCallback: true,
        reloadPage: true,
      });

      return this.errorResponseData(config);
    }

    if (OTHER_ERROR_CODES.indexOf(statusCode) !== -1) {
      return this.errorResponseData(config, result);
    }

    return this.errorResponseData(config, result);
  }

  private async request<T>(config: MyAxiosRequestConfig): Promise<CommonResponse<T | null>> {
    try {
      const resp: AxiosResponse<APIResponseData<T>> = await axios(config);

      return this.responseData<T>(config, resp.data);
    } catch(err) {
      console.warn(err);
      return this.errorResponseData(config);
    }
  }

  public async get<T>(config: MyAxiosRequestConfig): Promise<CommonResponse<T | null>> {
    return this.request<T>(
      Object.assign({}, config, {method: 'GET'})
    );
  }

  public async post<T>(config: MyAxiosRequestConfig): Promise<CommonResponse<T | null>> {
    return this.request<T>(
      Object.assign({json: false}, config, {method: 'POST' })
    );
  }
}

export default Services.instance();

(window as any).__Services__ = Services.instance();
