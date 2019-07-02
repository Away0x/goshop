<script lang="ts">
import { Component, Emit, Prop, Vue, Watch } from 'vue-property-decorator';


@Component({
  name: 'user-addresses-create-and-edit',
})
export default class UserAddressesCreateAndEdit extends Vue {
  /* ------------------------ INPUT & OUTPUT ------------------------ */

  /* ------------------------ VUEX (vuex getter & vuex action) ------------------------ */

  /* ------------------------ LIFECYCLE HOOKS (created & mounted & ...) ------------------------ */

  /* ------------------------ COMPONENT STATE (data & computed & model) ------------------------ */
  private province = ''; // 省
  private city = ''; // 市
  private district = ''; // 区

  /* ------------------------ WATCH ------------------------ */

  /* ------------------------ METHODS ------------------------ */
  private onDistrictChanged(val: string[]) {
    if (val.length === 3) {
      this.province = val[0];
      this.city = val[1];
      this.district = val[2];
    }
  }

}
</script>

<template>
<form class="user-addresses-create-and-edit-module form-horizontal" role="form">
  <!-- {{ csrf_field() }} -->

    <select-district @change="onDistrictChanged" />

  <input type="hidden" name="province" v-model="province">
  <input type="hidden" name="city" v-model="city">
  <input type="hidden" name="district" v-model="district">

  <div class="form-group row">
    <label class="col-form-label text-md-right col-sm-2">详细地址</label>
    <div class="col-sm-9">
      <input type="text" class="form-control" name="address" value="">
    </div>
  </div>
  <div class="form-group row">
    <label class="col-form-label text-md-right col-sm-2">邮编</label>
    <div class="col-sm-9">
      <input type="text" class="form-control" name="zip" value="">
    </div>
  </div>
  <div class="form-group row">
    <label class="col-form-label text-md-right col-sm-2">姓名</label>
    <div class="col-sm-9">
      <input type="text" class="form-control" name="contact_name" value="">
    </div>
  </div>
  <div class="form-group row">
    <label class="col-form-label text-md-right col-sm-2">电话</label>
    <div class="col-sm-9">
      <input type="text" class="form-control" name="contact_phone" value="">
    </div>
  </div>
  <div class="form-group row text-center">
    <div class="col-12">
      <button type="submit" class="btn btn-primary">提交</button>
    </div>
  </div>
</form>
</template>
