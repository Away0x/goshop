<script lang="ts">
import { Component, Emit, Prop, Vue, Watch } from 'vue-property-decorator';

@Component({
  name: 'user-addresses-create-and-edit',
  components: {},
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
<div class="module_UserAddressesCreateAndEdit">

</div>
</template>
