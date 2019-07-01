import 'bootstrap'
import $ from 'jquery'
import 'popper.js'
import Vue from 'vue';

(window as any).$ = $;
(window as any).jQuery = $;

import SelectDistrict from './components/SelectDistrict.vue';
import UserAddressesCreateAndEdit from './components/UserAddressesCreateAndEdit.vue';

Vue.component('select-district', SelectDistrict);
Vue.component('user-addresses-create-and-edit', UserAddressesCreateAndEdit);

new Vue({
  el: '#app'
});
