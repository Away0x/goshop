import 'bootstrap'
import $ from 'jquery'
import 'popper.js'
import Vue from 'vue';
import 'sweetalert';
import Axios, { AxiosRequestConfig } from 'axios';
import qs from 'qs';

const axios = Axios.create({
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
    'X-Requested-With': 'XMLHttpRequest',
  },
});

axios.interceptors.request.use((config: AxiosRequestConfig) => {
    if (!config.data) {
      config.data = {};
    }

    const csrf_meta = document.querySelector('meta[name="X-CSRF-Token"]');
    if (csrf_meta) {
      config.data['X-CSRF-Token'] = csrf_meta.getAttribute('content');
    }
    config.data = qs.stringify(config.data);

    return config;
  },
  (error: any) => {
    console.warn(error)
  },
);

(window as any).$ = $;
(window as any).jQuery = $;
(window as any).axios = axios;

import SelectDistrict from './components/SelectDistrict';
import UserAddressesCreateAndEdit from './components/UserAddressesCreateAndEdit';

Vue.component('select-district', SelectDistrict);
Vue.component('user-addresses-create-and-edit', UserAddressesCreateAndEdit);

new Vue({
  el: '#app',
  delimiters: ['@[', ']'],
});
