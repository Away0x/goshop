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
export const AUTH_API_ROOT: string = MAIN_CONFIG.AUTH_API_ROOT;
export const WORK_API_ROOT: string = MAIN_CONFIG.WORK_API_ROOT;
export const QUESTION_API_ROOT: string = MAIN_CONFIG.QUESTION_API_ROOT;
export const COACH_API_ROOT: string = MAIN_CONFIG.COACH_API_ROOT;
export const OFFICE_API_ROOT: string = MAIN_CONFIG.OFFICE_API_ROOT;
export const GUIDE_API_ROOT: string = MAIN_CONFIG.GUIDE_API_ROOT;
export const RESOURCE_API_ROOT: string = MAIN_CONFIG.RESOURCE_API_ROOT;
export const DICTIONARY_API_ROOT: string = MAIN_CONFIG.DICTIONARY_API_ROOT;

// ---------------------------- page url ----------------------------
export const ADMIN_URL: string = MAIN_CONFIG.ADMIN_URL; // 管理后台地址
export const RECORDING_PLATFORM: string = MAIN_CONFIG.RECORDING_PLATFORM; // 录播平台登录页面

// ---------------------------- ewebeditor config ----------------------------
export const EWEBEDITOR_URL = '/ewebeditor/ewebeditor.htm';
export const EWEBEDITOR_STYLE = 'expand600';
export const EWEBEDITOR_SKIN = 'flat10';
export function create_ewebeditor_src(editor_key: string) {
  return `${EWEBEDITOR_URL}?id=content${editor_key}&style=${EWEBEDITOR_STYLE}&skin=${EWEBEDITOR_SKIN}`;
}

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
export const WK_STORAGE_EXPIRE = MAIN_CONFIG.WK_STORAGE_EXPIRE || DISABLED_EXPIRE;
export const DISABLED_CHAPTER_TREE_CACHE = MAIN_CONFIG.DISABLED_CHAPTER_TREE_CACHE || false;
export const DISABLED_KNOW_TREE_CACHE = MAIN_CONFIG.DISABLED_KNOW_TREE_CACHE || false;

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
const apiConfigArr = [API_ROOT, AUTH_API_ROOT, WORK_API_ROOT, QUESTION_API_ROOT, COACH_API_ROOT, OFFICE_API_ROOT, GUIDE_API_ROOT, RESOURCE_API_ROOT, DICTIONARY_API_ROOT, ADMIN_URL];
const otherConfigArr = [LOCALSTROAGE_PREFIX, TOKEN_KEY];
if (!apiConfigArr.every((i) => !!i)) {
  console.error('【MAIN_CONFIG】 API PREFIX 加载错误，请检查配置文件', apiConfigArr);
}
if (!otherConfigArr.every((i) => !!i)) {
  console.error('【MAIN_CONFIG】 OTHER CONFIG 加载错误，请检查配置文件', otherConfigArr);
}
