const W: any = window;
const MAIN_CONFIG: { [key: string]: any } = W.MAIN_CONFIG;
if (!MAIN_CONFIG) { console.error('配置文件 MAIN_CONFIG.js 不存在，请检查!'); }

// ---------------------------- 开发环境判断 ----------------------------
const ENV_DEVELOPMENT = 'development';
// const ENV_PRODUCTION = 'production';

/** 是否为 development 环境 */
export function isDev() {
  return process.env.NODE_ENV === ENV_DEVELOPMENT;
}

// ---------------------------- 项目配置相关 ----------------------------
/** PUBLIC_URL */
export const PUBLIC_URL = process.env.PUBLIC_URL;

// ---------------------------- 接口根地址 ----------------------------
export const API_ROOT: string = MAIN_CONFIG.API_ROOT;

// ---------------------------- key ----------------------------
export const LOCALSTROAGE_PREFIX: string= MAIN_CONFIG.LOCALSTROAGE_PREFIX; // localstroage prefix
export const TOKEN_KEY: string = MAIN_CONFIG.TOKEN_KEY; // token key

// ---------------------------- cookie ----------------------------
interface CookiesType {
  set(key: string, val: string, attr?: { [key: string]: any }): string | void;
  get(key: string, json?: boolean): any;
  remove(key: string, attr?: { [key: string]: any }): void;
}

const Cookies: CookiesType = (window as any).Cookies;
export function getCookie() {
  if (!Cookies) {
    console.log('js-cookie.js 未加载');
    return null;
  }
  return Cookies;
}

// ---------------------------- cache config ----------------------------
export const DISABLED_EXPIRE = 666666; // 不设置缓存

// ---------------------------- 其他 ----------------------------
// 补丁函数
type PatchCallbackType = (params: any) => void;
export const PATCH_CALLBACK: PatchCallbackType | null = MAIN_CONFIG.PATCH_CALLBACK || null;
export const PATCH_CALLBACK_ROUTE_MIDDLEWARE = (params: any) => {
  if (PATCH_CALLBACK) {
    PATCH_CALLBACK(params);
  }
  return null;
}

// ---------------------------- 必有配置的加载检查 ----------------------------
const apiConfigArr = [API_ROOT];
const otherConfigArr = [LOCALSTROAGE_PREFIX, TOKEN_KEY];
if (!apiConfigArr.every((i) => !!i)) {
  console.error('【MAIN_CONFIG】 API PREFIX 加载错误，请检查配置文件', apiConfigArr);
}
if (!otherConfigArr.every((i) => !!i)) {
  console.error('【MAIN_CONFIG】 OTHER CONFIG 加载错误，请检查配置文件', otherConfigArr);
}
