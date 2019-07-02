import {
  Component,
  Vue,
  Prop,
  Watch,
  Emit,
} from 'vue-property-decorator';
import { filter, findKey } from 'lodash';

const addressData: AddressDataType = require('china-area-data/v3/data');

interface KeyValueStringMap {
  [key: string]: string
}
interface AddressDataType {
  [key: string]: KeyValueStringMap
}

@Component({
  name: 'select-district',
})
export default class SelectDistrict extends Vue {
  /* ------------------------ INPUT & OUTPUT ------------------------ */
  // 用来初始化省市区的值，在编辑时会用到
  @Prop({type: Array, default () { return [];}}) private initValue!: string[];
  @Emit('change') private changedEvent(vals: string[]) {}

  /* ------------------------ LIFECYCLE HOOKS (created & mounted & ...) ------------------------ */
  // created hook
  private created() {
    this.setFromValue(this.initValue);
  }

  /* ------------------------ COMPONENT STATE (data & computed & model) ------------------------ */
  // 省列表
  private provinces = addressData['86'];
  // 城市列表
  private cities: KeyValueStringMap = {};
  // 地区列表
  private districts: KeyValueStringMap = {};
  // 当前选中的省
  private provinceId = '';
  // 当前选中的市
  private cityId = '';
  // 当前选中的区
  private districtId = '';

  /* ------------------------ WATCH ------------------------ */
  // 当选择的省发生改变时触发
  @Watch('provinceId') private provinceIdChanged(newVal: string) {
    if (!newVal) {
      this.cities = {};
      this.cityId = '';
      return;
    }

    // 将城市列表设为当前省下的城市
    this.cities = addressData[newVal];
    // 如果当前选中的城市不在当前省下，则将选中城市清空
    if (!this.cities[this.cityId]) {
      this.cityId = '';
    }
  }

  // 当选择的市发生改变时触发
  @Watch('cityId') private cityIdChanged(newVal: string) {
    if (!newVal) {
      this.districts = {};
      this.districtId = '';
      return;
    }

    // 将地区列表设为当前城市下的地区
    this.districts = addressData[newVal];
    // 如果当前选中的地区不在当前城市下，则将选中地区清空
    if (!this.districts[this.districtId]) {
      this.districtId = '';
    }
  }

  // 当选择的区发生改变时触发
  @Watch('districtId') private districtIdChanged(newVal: string) {
    this.changedEvent([
      this.provinces[this.provinceId],
      this.cities[this.cityId],
      this.districts[this.districtId]
    ]);
  }

  /* ------------------------ METHODS ------------------------ */
  private setFromValue(value: string[]) {
    value = filter(value)
    // 如果数组长度为 0，则将省清空（由于我们定义了观察器，会联动触发将城市和地区清空）
    if (value.length === 0) {
      this.provinceId = '';
      return;
    }

    // 从当前省列表中找到与数组第一个元素同名的项的索引
    const provinceId = findKey(this.provinces, o => o === value[0]);
    // 没找到，清空省的值
    if (!provinceId) {
      this.provinceId = '';
      return;
    }
    // 找到了，将当前省设置成对应的ID
    this.provinceId = provinceId;

    // 由于观察器的作用，这个时候城市列表已经变成了对应省的城市列表
    // 从当前城市列表找到与数组第二个元素同名的项的索引
    const cityId = findKey(addressData[provinceId], o => o === value[1]);
    // 没找到，清空城市的值
    if (!cityId) {
      this.cityId = '';
      return;
    }
    // 找到了，将当前城市设置成对应的ID
    this.cityId = cityId;

    // 由于观察器的作用，这个时候地区列表已经变成了对应城市的地区列表
    // 从当前地区列表找到与数组第三个元素同名的项的索引
    const districtId = findKey(addressData[cityId], o => o === value[2]);
    // 没找到，清空地区的值
    if (!districtId) {
      this.districtId = '';
      return;
    }
    // 找到了，将当前地区设置成对应的ID
    this.districtId = districtId;
  }
}
