import 'bootstrap'
import $ from 'jquery'
import 'popper.js'
import Vue from 'vue';

(window as any).$ = $;
(window as any).jQuery = $;

import SelectDistrict from './components/SelectDistrict';
import UserAddressesCreateAndEdit from './components/UserAddressesCreateAndEdit';

Vue.component('select-district', SelectDistrict);
Vue.component('user-addresses-create-and-edit', UserAddressesCreateAndEdit);

new Vue({
  el: '#app',
  delimiters: ['@[', ']'],
});
