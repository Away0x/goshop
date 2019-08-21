(window["webpackJsonp"] = window["webpackJsonp"] || []).push([[6],{

/***/ 1:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(170);

/***/ }),

/***/ 101:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(83);

/***/ }),

/***/ 105:
/***/ (function(module, exports) {



/***/ }),

/***/ 106:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* unused harmony export withStoreConsumer */
/* unused harmony export withStoreMapConsumer */
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return withStoreProviders; });
/* harmony import */ var C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(156);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(1);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);
/** 注入单个 store 包包装的组件中 */function withStoreConsumer(Component,store){var WrappedComponent=function WrappedComponent(props){var s=store.useStore();return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(Component,Object.assign({},props,{store:s}));};WrappedComponent.displayName="withStoreConsumer(".concat(Component.displayName||Component.name,")");return WrappedComponent;}/** 注入多个 store 到包装的组件中 */function withStoreMapConsumer(Component,storeMap){var WrappedComponent=function WrappedComponent(props){var sm={};for(var k in storeMap){if(storeMap.hasOwnProperty(k)){sm[k]=storeMap[k].useStore();}}return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(Component,Object.assign({},props,{storeMap:sm}));};WrappedComponent.displayName="withStoreMapConsumer(".concat(Component.displayName||Component.name,")");return WrappedComponent;}/** 组合 Provider */function withStoreProviders(Component,stores)// index 靠前的会包裹 index 靠后的
{if(!stores.length){return Component;}var components=[];stores.forEach(function(c){return components.push(c.Provider);});components.push(Component);var createTree=function createTree(index,children,props){var isLastNode=index===components.length-1;var Com=components[index];if(isLastNode){return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(Com,props,children);}return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(Com,props,createTree(++index,children,props));};var WrappedComponent=function WrappedComponent(_ref){var children=_ref.children,rest=Object(C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_0__[/* default */ "a"])(_ref,["children"]);return createTree(0,children,rest);};WrappedComponent.displayName="withStoreProviders(".concat(Component.displayName||Component.name,")");return WrappedComponent;}

/***/ }),

/***/ 107:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return createStore; });
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(1);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_0__);
// unstated-next: https://github.com/jamiebuilds/unstated-next
function createStore(useHook){var Context=react__WEBPACK_IMPORTED_MODULE_0___default.a.createContext(null);function Provider(props){var value=useHook();return react__WEBPACK_IMPORTED_MODULE_0___default.a.createElement(Context.Provider,{value:value},props.children);}function useStore(){var value=react__WEBPACK_IMPORTED_MODULE_0___default.a.useContext(Context);if(value===null){throw new Error("Component must be wrapped with <Store.Provider>");}return value;}return{Provider:Provider,useStore:useStore};}// function component 使用 hook 获取 store
// class component 使用 hoc 获取 store

/***/ }),

/***/ 112:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(230);

/***/ }),

/***/ 135:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(228);

/***/ }),

/***/ 137:
/***/ (function(module, exports) {

module.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHgAAAB4CAYAAAA5ZDbSAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyFpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQyIDc5LjE2MDkyNCwgMjAxNy8wNy8xMy0wMTowNjozOSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIChXaW5kb3dzKSIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDpERThDMTlBREREQTgxMUU4QjcwMTlGQ0VFNzg4MzlGMiIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDpERThDMTlBRUREQTgxMUU4QjcwMTlGQ0VFNzg4MzlGMiI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOkRFOEMxOUFCRERBODExRThCNzAxOUZDRUU3ODgzOUYyIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOkRFOEMxOUFDRERBODExRThCNzAxOUZDRUU3ODgzOUYyIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+AE41zAAAFI5JREFUeNrsXQmYFNW1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"

/***/ }),

/***/ 142:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return useBreadcrumbs; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "b", function() { return useNavs; });
/* harmony import */ var C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_toConsumableArray__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(14);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(1);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var aw_react_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(15);
/* harmony import */ var aw_react_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(aw_react_router__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _store_base_layout__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(51);
/* harmony import */ var _auth_authority__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(43);
/* harmony import */ var _tools_instance__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(32);
/* harmony import */ var _constants__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(33);
/** 修改面包屑 */function useBreadcrumbs(initBreadcrumbs){var store=_store_base_layout__WEBPACK_IMPORTED_MODULE_3__[/* default */ "a"].useStore();var setBreadcrumbs=function setBreadcrumbs(breadcrumbs){store.setBreadcrumbsAction([{name:'首页',routeName:'home'}].concat(Object(C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_toConsumableArray__WEBPACK_IMPORTED_MODULE_0__[/* default */ "a"])(breadcrumbs)));};Object(react__WEBPACK_IMPORTED_MODULE_1__["useEffect"])(function(){setBreadcrumbs(initBreadcrumbs);},[]);return{breadcrumbs:store.state.breadcrumbs,setBreadcrumbs:setBreadcrumbs};}var NAV_MAX_LEN=8;// 最后一个为 "更多"
/** 创建页面导航 */function useNavs(){var navs=[];var moreNavs={defaultName:'更多',subNav:[]};var routeInfos=aw_react_router__WEBPACK_IMPORTED_MODULE_2___default.a.instance().findMap(_constants__WEBPACK_IMPORTED_MODULE_6__[/* PAGE_NAMES */ "b"].map(function(t){return t.name;}));_constants__WEBPACK_IMPORTED_MODULE_6__[/* PAGE_NAMES */ "b"].forEach(function(r){var name=r.name;var info=routeInfos[name];if(!info){return;}var meta=info.meta||{};var strict=Object(_tools_instance__WEBPACK_IMPORTED_MODULE_5__[/* isUndefined */ "a"])(meta.authUnstrict)?true:!meta.authUnstrict;;if(_auth_authority__WEBPACK_IMPORTED_MODULE_4__[/* default */ "a"].checkAll(meta.auth||[],strict)){if(navs.length>=NAV_MAX_LEN-1){moreNavs.subNav.push({route:routeInfos[name]});return;}navs.push({route:routeInfos[name]});}});if(moreNavs.subNav.length>1){return[].concat(navs,[moreNavs]);}else if(moreNavs.subNav.length===1){return[].concat(navs,[moreNavs.subNav[0]]);}return navs;}

/***/ }),

/***/ 145:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";

// CONCATENATED MODULE: ./src/tools/hoc/ts-enhance.ts
// ts react 的一些推断辅助
/**
 * typescript defaultProps 的默认 props 辅助 hoc
 */var withDefaultProps=function withDefaultProps(defaultProps,Cmp){Cmp.defaultProps=defaultProps;return Cmp;};
// CONCATENATED MODULE: ./src/tools/hoc/index.ts
/* concated harmony reexport withDefaultProps */__webpack_require__.d(__webpack_exports__, "a", function() { return withDefaultProps; });


/***/ }),

/***/ 166:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(84);

/***/ }),

/***/ 170:
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(294);


/***/ }),

/***/ 172:
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin

/***/ }),

/***/ 173:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(425);

/***/ }),

/***/ 20:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/regenerator/index.js
var regenerator = __webpack_require__(8);
var regenerator_default = /*#__PURE__*/__webpack_require__.n(regenerator);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/asyncToGenerator.js
var asyncToGenerator = __webpack_require__(18);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/classCallCheck.js
var classCallCheck = __webpack_require__(12);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/createClass.js
var createClass = __webpack_require__(9);

// EXTERNAL MODULE: ./src/storage/base.ts + 1 modules
var base = __webpack_require__(53);

// EXTERNAL MODULE: ./src/config/index.ts
var src_config = __webpack_require__(21);

// CONCATENATED MODULE: ./src/storage/token.ts
/** token */// 浏览器退出时清除 token
var token_TokenStorage=/*#__PURE__*/function(){Object(createClass["a" /* default */])(TokenStorage,null,[{key:"instance",value:function instance(){if(this._instance===null){this._instance=new this();}return this._instance;}}]);function TokenStorage(){Object(classCallCheck["a" /* default */])(this,TokenStorage);this.storage=void 0;this.key=src_config["e" /* TOKEN_KEY */];this.tokenStore=null;this.storage=Object(base["a" /* default */])({hasMeta:false,session:true,prefix:''});}Object(createClass["a" /* default */])(TokenStorage,[{key:"set",// 内存缓存
value:function set(token){var cookie=Object(src_config["f" /* getCookie */])();if(cookie){cookie.set(this.key,token);}this.storage.setString(this.key,token);this.tokenStore=token;}/** 取值优先级 (内存 > cookie > sessionStorage) */},{key:"get",value:function get(){if(this.tokenStore){return this.tokenStore;}var cookie=Object(src_config["f" /* getCookie */])();// cookie 有则读 cookie，否则读 sessionStorage，然后同步到 cookie
var tokenFromLocal='';if(cookie){tokenFromLocal=cookie.get(this.key);}if(!tokenFromLocal){tokenFromLocal=this.storage.getString(this.key)||'';if(cookie){cookie.set(this.key,tokenFromLocal);}}this.tokenStore=tokenFromLocal;return this.tokenStore||'';}},{key:"clean",value:function clean(){var cookie=Object(src_config["f" /* getCookie */])();if(cookie){cookie.remove(this.key);}this.storage.remove(this.key);this.tokenStore=null;}}]);return TokenStorage;}();token_TokenStorage._instance=null;/* harmony default export */ var storage_token = (token_TokenStorage.instance());window.__TokenStorage__=token_TokenStorage.instance();
// EXTERNAL MODULE: ./node_modules/aw-react-router/lib/index.js
var lib = __webpack_require__(15);
var lib_default = /*#__PURE__*/__webpack_require__.n(lib);

// CONCATENATED MODULE: ./src/auth/token.ts
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return token_default; });
// 登出、获取用户 token 操作
var token_default=/*#__PURE__*/function(){function _default(){Object(classCallCheck["a" /* default */])(this,_default);}Object(createClass["a" /* default */])(_default,null,[{key:"set",// main.ts 中挂载的回调 (将不方便在这个类中做的事情放到 main.ts 中去做)
value:function set(token){storage_token.set(token);}},{key:"get",value:function get(){return storage_token.get();}},{key:"isLogin",value:function isLogin(){return!!this.get();}},{key:"clean",value:function clean(){storage_token.clean();}/**
   * 登出操作
   */},{key:"logout",value:function(){var _logout=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee(config){var status,_status;return regenerator_default.a.wrap(function _callee$(_context){while(1){switch(_context.prev=_context.next){case 0:if(!config.beforeCallback){_context.next=6;break;}_context.next=3;return config.beforeCallback(this.get());case 3:status=_context.sent;if(status){_context.next=6;break;}return _context.abrupt("return");case 6:this.clean();if(!config.afterCallback){_context.next=13;break;}_context.next=10;return config.afterCallback();case 10:_status=_context.sent;if(_status){_context.next=13;break;}return _context.abrupt("return");case 13:if(!(config.triggerGlobalCallback&&this.logoutCallback)){_context.next=16;break;}_context.next=16;return this.logoutCallback();case 16:if(config.reloadPage){window.location.href='/';}else{lib_default.a.instance().pushByName('login',{replace:true});}case 17:case"end":return _context.stop();}}},_callee,this);}));function logout(_x){return _logout.apply(this,arguments);}return logout;}()}]);return _default;}();token_default.key=storage_token.key;token_default.logoutCallback=null;

/***/ }),

/***/ 21:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "g", function() { return isDev; });
/* unused harmony export PUBLIC_URL */
/* unused harmony export API_ROOT */
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AUTH_API_ROOT; });
/* unused harmony export WORK_API_ROOT */
/* unused harmony export QUESTION_API_ROOT */
/* unused harmony export COACH_API_ROOT */
/* unused harmony export OFFICE_API_ROOT */
/* unused harmony export GUIDE_API_ROOT */
/* unused harmony export RESOURCE_API_ROOT */
/* unused harmony export DICTIONARY_API_ROOT */
/* unused harmony export ADMIN_URL */
/* unused harmony export RECORDING_PLATFORM */
/* unused harmony export EWEBEDITOR_URL */
/* unused harmony export EWEBEDITOR_STYLE */
/* unused harmony export EWEBEDITOR_SKIN */
/* unused harmony export create_ewebeditor_src */
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "c", function() { return LOCALSTROAGE_PREFIX; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "e", function() { return TOKEN_KEY; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "f", function() { return getCookie; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "b", function() { return DISABLED_EXPIRE; });
/* unused harmony export WK_STORAGE_EXPIRE */
/* unused harmony export DISABLED_CHAPTER_TREE_CACHE */
/* unused harmony export DISABLED_KNOW_TREE_CACHE */
/* unused harmony export PATCH_CALLBACK */
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "d", function() { return PATCH_CALLBACK_ROUTE_MIDDLEWARE; });
var W=window;var MAIN_CONFIG=W.MAIN_CONFIG;if(!MAIN_CONFIG){console.error('配置文件 MAIN_CONFIG.js 不存在，请检查!');}// ---------------------------- 开发环境判断 ----------------------------
var ENV_DEVELOPMENT='development';// const ENV_PRODUCTION = 'production';
/** 是否为 development 环境 */function isDev(){return "production"===ENV_DEVELOPMENT;}// ---------------------------- 项目配置相关 ----------------------------
/** PUBLIC_URL */var PUBLIC_URL="";// ---------------------------- 接口根地址 ----------------------------
var API_ROOT=MAIN_CONFIG.API_ROOT;var AUTH_API_ROOT=MAIN_CONFIG.AUTH_API_ROOT;var WORK_API_ROOT=MAIN_CONFIG.WORK_API_ROOT;var QUESTION_API_ROOT=MAIN_CONFIG.QUESTION_API_ROOT;var COACH_API_ROOT=MAIN_CONFIG.COACH_API_ROOT;var OFFICE_API_ROOT=MAIN_CONFIG.OFFICE_API_ROOT;var GUIDE_API_ROOT=MAIN_CONFIG.GUIDE_API_ROOT;var RESOURCE_API_ROOT=MAIN_CONFIG.RESOURCE_API_ROOT;var DICTIONARY_API_ROOT=MAIN_CONFIG.DICTIONARY_API_ROOT;// ---------------------------- page url ----------------------------
var ADMIN_URL=MAIN_CONFIG.ADMIN_URL;// 管理后台地址
var RECORDING_PLATFORM=MAIN_CONFIG.RECORDING_PLATFORM;// 录播平台登录页面
// ---------------------------- ewebeditor config ----------------------------
var EWEBEDITOR_URL='/ewebeditor/ewebeditor.htm';var EWEBEDITOR_STYLE='expand600';var EWEBEDITOR_SKIN='flat10';function create_ewebeditor_src(editor_key){return"".concat(EWEBEDITOR_URL,"?id=content").concat(editor_key,"&style=").concat(EWEBEDITOR_STYLE,"&skin=").concat(EWEBEDITOR_SKIN);}// ---------------------------- key ----------------------------
var LOCALSTROAGE_PREFIX=MAIN_CONFIG.LOCALSTROAGE_PREFIX;// localstroage prefix
var TOKEN_KEY=MAIN_CONFIG.TOKEN_KEY;// token key
// ---------------------------- cookie ----------------------------
var Cookies=window.Cookies;function getCookie(){if(!Cookies){console.log('js-cookie.js 未加载');return null;}return Cookies;}// ---------------------------- cache config ----------------------------
var DISABLED_EXPIRE=666666;// 不设置缓存
var WK_STORAGE_EXPIRE=MAIN_CONFIG.WK_STORAGE_EXPIRE||DISABLED_EXPIRE;var DISABLED_CHAPTER_TREE_CACHE=MAIN_CONFIG.DISABLED_CHAPTER_TREE_CACHE||false;var DISABLED_KNOW_TREE_CACHE=MAIN_CONFIG.DISABLED_KNOW_TREE_CACHE||false;// ---------------------------- 其他 ----------------------------
// 补丁函数
var PATCH_CALLBACK=MAIN_CONFIG.PATCH_CALLBACK||null;var PATCH_CALLBACK_ROUTE_MIDDLEWARE=function PATCH_CALLBACK_ROUTE_MIDDLEWARE(params){if(PATCH_CALLBACK){PATCH_CALLBACK(params);}return null;};// ---------------------------- 必有配置的加载检查 ----------------------------
var apiConfigArr=[API_ROOT,AUTH_API_ROOT,WORK_API_ROOT,QUESTION_API_ROOT,COACH_API_ROOT,OFFICE_API_ROOT,GUIDE_API_ROOT,RESOURCE_API_ROOT,DICTIONARY_API_ROOT,ADMIN_URL];var otherConfigArr=[LOCALSTROAGE_PREFIX,TOKEN_KEY];if(!apiConfigArr.every(function(i){return!!i;})){console.error('【MAIN_CONFIG】 API PREFIX 加载错误，请检查配置文件',apiConfigArr);}if(!otherConfigArr.every(function(i){return!!i;})){console.error('【MAIN_CONFIG】 OTHER CONFIG 加载错误，请检查配置文件',otherConfigArr);}

/***/ }),

/***/ 246:
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin

/***/ }),

/***/ 254:
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin

/***/ }),

/***/ 255:
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin

/***/ }),

/***/ 283:
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin

/***/ }),

/***/ 289:
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin

/***/ }),

/***/ 294:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);

// EXTERNAL MODULE: ./node_modules/antd/lib/message/style/index.js
var style = __webpack_require__(100);

// EXTERNAL MODULE: ./node_modules/antd/lib/message/index.js
var message = __webpack_require__(50);
var message_default = /*#__PURE__*/__webpack_require__.n(message);

// EXTERNAL MODULE: ./src/assets/styles/index.less
var styles = __webpack_require__(172);

// EXTERNAL MODULE: delegated ../node_modules/react/index.js from dll-reference main_library
var reactfrom_dll_reference_main_library = __webpack_require__(1);
var reactfrom_dll_reference_main_library_default = /*#__PURE__*/__webpack_require__.n(reactfrom_dll_reference_main_library);

// EXTERNAL MODULE: delegated ../node_modules/react-dom/index.js from dll-reference main_library
var react_domfrom_dll_reference_main_library = __webpack_require__(6);
var react_domfrom_dll_reference_main_library_default = /*#__PURE__*/__webpack_require__.n(react_domfrom_dll_reference_main_library);

// EXTERNAL MODULE: ./node_modules/fastclick/lib/fastclick.js
var fastclick = __webpack_require__(146);
var fastclick_default = /*#__PURE__*/__webpack_require__.n(fastclick);

// EXTERNAL MODULE: ./src/constants/index.ts + 3 modules
var constants = __webpack_require__(33);

// EXTERNAL MODULE: ./src/config/index.ts
var config = __webpack_require__(21);

// EXTERNAL MODULE: ./src/tools/aw/aw-store/index.tsx
var aw_store = __webpack_require__(31);

// EXTERNAL MODULE: ./src/store/auth/index.ts + 6 modules
var auth = __webpack_require__(58);

// EXTERNAL MODULE: ./node_modules/react-router-dom/es/HashRouter.js + 2 modules
var HashRouter = __webpack_require__(301);

// EXTERNAL MODULE: ./node_modules/history/esm/history.js + 3 modules
var esm_history = __webpack_require__(10);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/toConsumableArray.js + 3 modules
var toConsumableArray = __webpack_require__(14);

// EXTERNAL MODULE: ./node_modules/aw-react-router/lib/index.js
var lib = __webpack_require__(15);
var lib_default = /*#__PURE__*/__webpack_require__.n(lib);

// CONCATENATED MODULE: ./src/routes/suspense.tsx
/* harmony default export */ var suspense = (Object(lib["createSuspense"])({defaultDelay:0,defaultLoading:reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"common_page_placeholder"},"loading ... ...")}));
// EXTERNAL MODULE: ./src/layouts/Base/index.less
var Base = __webpack_require__(246);

// EXTERNAL MODULE: ./src/store/base-layout/index.ts
var base_layout = __webpack_require__(51);

// EXTERNAL MODULE: ./node_modules/antd/lib/row/style/index.js
var row_style = __webpack_require__(99);

// EXTERNAL MODULE: ./node_modules/antd/lib/row/index.js
var row = __webpack_require__(65);
var row_default = /*#__PURE__*/__webpack_require__.n(row);

// EXTERNAL MODULE: ./node_modules/antd/lib/col/style/index.js
var col_style = __webpack_require__(96);

// EXTERNAL MODULE: ./node_modules/antd/lib/col/index.js
var col = __webpack_require__(42);
var col_default = /*#__PURE__*/__webpack_require__.n(col);

// EXTERNAL MODULE: ./src/tools/hoc/index.ts + 1 modules
var hoc = __webpack_require__(145);

// EXTERNAL MODULE: ./src/tools/instance.ts
var instance = __webpack_require__(32);

// CONCATENATED MODULE: ./src/tools/index.ts
/**
 * safeGet 安全的获取对象值
 * obj = {a: {b: {c: 123}}}
 *
 * safeGet(obj, 'a.b.c') // 123
 * safeGet(obj, 'a.b.d', '我是默认值') // 我是默认值
 *
 * @param   {Object}  目标对象
 * @param   {String}  路径字符串 "a.b.c"
 * @param   {Any}     默认值
 * @return  {Any}     获取到的值
 */function safeGet(targetObj,keyString,defaultValue){var path=String(keyString).split('.');var result=path.reduce(function(preObj,curPath){return preObj!==null&&!Object(instance["a" /* isUndefined */])(preObj)?preObj[curPath]:preObj;},targetObj);return Object(instance["a" /* isUndefined */])(result)?Object(instance["a" /* isUndefined */])(defaultValue)?result:defaultValue:result;}
// EXTERNAL MODULE: ./src/assets/images/logo.png
var logo = __webpack_require__(92);
var logo_default = /*#__PURE__*/__webpack_require__.n(logo);

// EXTERNAL MODULE: ./node_modules/antd/lib/dropdown/style/index.js
var dropdown_style = __webpack_require__(74);

// EXTERNAL MODULE: ./node_modules/antd/lib/dropdown/index.js
var dropdown = __webpack_require__(66);
var dropdown_default = /*#__PURE__*/__webpack_require__.n(dropdown);

// EXTERNAL MODULE: ./node_modules/antd/lib/icon/style/index.js
var icon_style = __webpack_require__(165);

// EXTERNAL MODULE: ./node_modules/antd/lib/icon/index.js
var icon = __webpack_require__(26);
var icon_default = /*#__PURE__*/__webpack_require__.n(icon);

// EXTERNAL MODULE: ./node_modules/antd/lib/menu/style/index.js
var menu_style = __webpack_require__(75);

// EXTERNAL MODULE: ./node_modules/antd/lib/menu/index.js
var menu = __webpack_require__(24);
var menu_default = /*#__PURE__*/__webpack_require__.n(menu);

// EXTERNAL MODULE: ./node_modules/react-router-dom/es/Link.js
var Link = __webpack_require__(300);

// EXTERNAL MODULE: ./src/components/MyHeader/MenuItem.less
var MyHeader_MenuItem = __webpack_require__(254);

// EXTERNAL MODULE: ./src/components/MyHeader/UserAvatar.less
var MyHeader_UserAvatar = __webpack_require__(255);

// CONCATENATED MODULE: ./src/components/MyHeader/MenuItem.tsx
var MenuItem_menus=function menus(nav){var subNavs=nav.subNav||[];return reactfrom_dll_reference_main_library_default.a.createElement(menu_default.a,null,subNavs.map(function(n,i){return reactfrom_dll_reference_main_library_default.a.createElement(menu_default.a.Item,{key:i},reactfrom_dll_reference_main_library_default.a.createElement(Link["a" /* default */],{to:n.route.fullPath},n.route.meta.hanName));}));};var MenuItem_MenuItem=function MenuItem(_ref){var nav=_ref.nav,active=_ref.active;var classnames=['menu_item_component',active?'active':''].filter(Boolean).join(' ');var hasSubNav=nav.subNav&&nav.subNav.length>0&&nav.defaultName;return reactfrom_dll_reference_main_library_default.a.createElement("div",{className:classnames},hasSubNav?reactfrom_dll_reference_main_library_default.a.createElement(dropdown_default.a,{overlayClassName:"header_menu_item_content",overlay:MenuItem_menus(nav)},reactfrom_dll_reference_main_library_default.a.createElement("div",null,nav.defaultName," ",reactfrom_dll_reference_main_library_default.a.createElement(icon_default.a,{type:"down"}))):reactfrom_dll_reference_main_library_default.a.createElement(Link["a" /* default */],{to:nav.route.fullPath},nav.route.meta.hanName));};/* harmony default export */ var components_MyHeader_MenuItem = (MenuItem_MenuItem);
// EXTERNAL MODULE: ./src/assets/images/default_teacher_avatar.png
var default_teacher_avatar = __webpack_require__(137);
var default_teacher_avatar_default = /*#__PURE__*/__webpack_require__.n(default_teacher_avatar);

// CONCATENATED MODULE: ./src/components/MyHeader/UserAvatar.tsx
var UserAvatar_UserAvatar=function UserAvatar(){var store=auth["a" /* default */].useStore();var user=store.state.user;var settingsPath=lib_default.a.instance().getPathByName('settings');var menus=reactfrom_dll_reference_main_library_default.a.createElement(menu_default.a,null,reactfrom_dll_reference_main_library_default.a.createElement(menu_default.a.Item,null,reactfrom_dll_reference_main_library_default.a.createElement(Link["a" /* default */],{to:settingsPath},"\u4E2A\u4EBA\u4E2D\u5FC3")),reactfrom_dll_reference_main_library_default.a.createElement(menu_default.a.Item,null,reactfrom_dll_reference_main_library_default.a.createElement("p",{onClick:store.logoutAction},"\u9000\u51FA")));return reactfrom_dll_reference_main_library_default.a.createElement(dropdown_default.a,{overlay:menus},reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"user_avatar_component"},reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"user_avatar"},reactfrom_dll_reference_main_library_default.a.createElement("img",{src:user.fullPath||default_teacher_avatar_default.a,alt:"avatar"}))));};/* harmony default export */ var components_MyHeader_UserAvatar = (UserAvatar_UserAvatar);
// EXTERNAL MODULE: ./src/components/MyHeader/index.less
var components_MyHeader = __webpack_require__(283);

// CONCATENATED MODULE: ./src/components/MyHeader/index.tsx
var defaultProps={canGoRoot:true,navs:[]};function isActiveNav(nav){var currentRoute=lib_default.a.instance().getCurrentState();if(!currentRoute){return false;}var meta=currentRoute.meta;if(!meta){return false;}return meta.topNavName===safeGet(nav,'route.meta.topNavName','');}var MyHeader_MyHeader=function MyHeader(_ref){var navs=_ref.navs,canGoRoot=_ref.canGoRoot;/** 回到首页 */var gotoRootPage=function gotoRootPage(){lib_default.a.instance().pushByName('home');};return reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"my_header_component"},reactfrom_dll_reference_main_library_default.a.createElement(row_default.a,{className:"menu_inner"},reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:3},canGoRoot?reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"menu_logo",title:"\u70B9\u51FB\u56DE\u5230\u9996\u9875",onClick:gotoRootPage},reactfrom_dll_reference_main_library_default.a.createElement("img",{alt:"logo",src:logo_default.a}),reactfrom_dll_reference_main_library_default.a.createElement("span",{className:"logo_title"},"\u9996\u9875")):reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"menu_logo",style:{cursor:'pointer'}},reactfrom_dll_reference_main_library_default.a.createElement("img",{alt:"logo",src:logo_default.a}))),reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:19},navs.map(function(n,i){return reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"menu_wrapper_item",key:i,style:{width:100/navs.length+'%'}},reactfrom_dll_reference_main_library_default.a.createElement(components_MyHeader_MenuItem,{nav:n,active:isActiveNav(n)}));})),reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:2},reactfrom_dll_reference_main_library_default.a.createElement(components_MyHeader_UserAvatar,null))));};/* harmony default export */ var src_components_MyHeader = (Object(hoc["a" /* withDefaultProps */])(defaultProps,MyHeader_MyHeader));
// EXTERNAL MODULE: ./src/components/MyFooter/index.less
var components_MyFooter = __webpack_require__(289);

// CONCATENATED MODULE: ./src/components/MyFooter/index.tsx
var year=new Date().getFullYear();var MyFooter_MyFooter=function MyFooter(){return reactfrom_dll_reference_main_library_default.a.createElement("footer",{className:"my_footer_component"},reactfrom_dll_reference_main_library_default.a.createElement(row_default.a,{className:"inner"},reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:10},"Copyright \xA9 2009 - ",year," Edutech, Inc. All rights reserved."),reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:14},"\u7248\u6743\u6240\u6709\uFF1A\u4E0A\u6D77\u6613\u6559\u4FE1\u606F\u79D1\u6280\u6709\u9650\u516C\u53F8 \xA0\xA0\xA0\xA0\xA0",reactfrom_dll_reference_main_library_default.a.createElement("a",{rel:"noopener noreferrer",href:"http://www.educationtek.com/",target:"_blank"},"www.educationtek.com"))));};/* harmony default export */ var src_components_MyFooter = (MyFooter_MyFooter);
// EXTERNAL MODULE: ./node_modules/antd/lib/breadcrumb/style/index.js
var breadcrumb_style = __webpack_require__(290);

// EXTERNAL MODULE: ./node_modules/antd/lib/breadcrumb/index.js
var breadcrumb = __webpack_require__(93);
var breadcrumb_default = /*#__PURE__*/__webpack_require__.n(breadcrumb);

// CONCATENATED MODULE: ./src/layouts/Base/Breadcrumbs.tsx
var Breadcrumbs_Breadcrumbs=function Breadcrumbs(){var _BaseLayoutStore$useS=base_layout["a" /* default */].useStore(),state=_BaseLayoutStore$useS.state;var defaultHandler=Object(reactfrom_dll_reference_main_library["useCallback"])(function(b){if(!b.handler&&b.routeName){lib_default.a.instance().pushByName(b.routeName);}},[]);return reactfrom_dll_reference_main_library_default.a.createElement(breadcrumb_default.a,{separator:">"},state.breadcrumbs.map(function(b){var handler=b.handler||defaultHandler;return reactfrom_dll_reference_main_library_default.a.createElement(breadcrumb_default.a.Item,{key:b.name,onClick:function onClick(){return handler(b);}},b.name);}));};/* harmony default export */ var Base_Breadcrumbs = (Object(reactfrom_dll_reference_main_library["memo"])(Breadcrumbs_Breadcrumbs));
// EXTERNAL MODULE: ./src/layouts/Base/hooks.ts
var hooks = __webpack_require__(142);

// CONCATENATED MODULE: ./src/layouts/Base/index.tsx
var BaseLayout=Object(reactfrom_dll_reference_main_library["memo"])(function(_ref){var routerView=_ref.routerView;var navs=Object(hooks["b" /* useNavs */])();return reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"base_layout"},reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"header_wrapper"},reactfrom_dll_reference_main_library_default.a.createElement(src_components_MyHeader,{navs:navs})),reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"breadcrumb_wrapper"},reactfrom_dll_reference_main_library_default.a.createElement(Base_Breadcrumbs,null)),reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"main"},routerView()),reactfrom_dll_reference_main_library_default.a.createElement(src_components_MyFooter,null));});BaseLayout.displayName='BaseLayout';/* harmony default export */ var layouts_Base = (Object(aw_store["withStoreProviders"])(BaseLayout,[base_layout["a" /* default */]]));
// CONCATENATED MODULE: ./src/auth/map.ts
// 完整权限映射表
/** ---------------------------------------- 导学本 ---------------------------------------- */ // 导学本模块
var auth_guide='guide_book_tea';// 导学本新建
var auth_guide__gbt_guide_book_add='gbt_guide_book_add';// 导学本删除
var auth_guide__gbt_guide_book_del='gbt_guide_book_del';// 导学本移动
var auth_guide__gbt_guide_book_move='gbt_guide_book_move';// 导学本共享
var auth_guide__gbt_guide_book_share='gbt_guide_book_share';// 下载资源
var auth_guide__gbt_resource_download='gbt_resource_download';// 测验
var auth_guide__gbt_examination='gbt_examination';/** ---------------------------------------- 统计分析 ---------------------------------------- */ // 统计分析模块
var auth_statistics='statistics_tea';/** ---------------------------------------- 作业辅导 ---------------------------------------- */ // 作业辅导模块
var auth_coach='coach_tea';// 新建辅导
var auth_coach__ct_coach_add='ct_coach_add';// 删除辅导
var auth_coach__ct_coach_del='ct_coach_del';// 导入微课
var auth_coach__ct_WK_import='ct_WK_import';// 题库选题
var auth_coach__ct_questionbank_selection='ct_questionbank_selection';// 导出题目
var auth_coach__ct_question_export='ct_question_export';// 看讲解
var auth_coach__ct_read_explanation='ct_read_explanation';// 解析答案
var auth_coach__ct_analyze_answer='ct_analyze_answer';// 加入错题集
var auth_coach__ct_mistakecollection_add='ct_mistakecollection_add';// 移出辅导
var auth_coach__ct_coach_rmv='ct_coach_rmv';// 下载辅导
var auth_coach__ct_coach_download='ct_coach_download';// 点赞
var auth_coach__ct_thumb_up='ct_thumb-up';// 关联知识点
var auth_coach__ct_knowledge_associate='ct_knowledge_associate';/** ---------------------------------------- 网络作业 ---------------------------------------- */ // 网络作业模块
var auth_homework='homework_tea';// 布置作业
var auth_homework__ht_homework_arrange='ht_homework_arrange';// 删除
var auth_homework__ht_homework_del='ht_homework_del';// 发布到其他班级
var auth_homework__ht_otherclass_publish='ht_otherclass_publish';// 加入错题集
var auth_homework__ht_mistakecollection_add='ht_mistakecollection_add';// 复制作业
var auth_homework__ht_homework_copy='ht_homework_copy';// 下载作业
var auth_homework__ht_homework_download='ht_homework_download';// 下载答案
var auth_homework__ht_homework_answer='ht_homework_answer';// 分享作业
var auth_homework__ht_homework_share='ht_homework_share';// 参考资料
var auth_homework__ht_homework_attachment='ht_homework_attachment';// 求讲解
var auth_homework__ht_answer_temporary_save='ht_answer_temporary_save';// 暂存作业答案
var auth_homework__ht_explanation_seek='ht_explanation_seek';/** ---------------------------------------- 错题集 ---------------------------------------- */ // 错题集模块
var auth_mistakecollection='mistakecollection_tea';/** ---------------------------------------- 备课中心(个人) ---------------------------------------- */ // 个人备课中心模块
var auth_resource='resource_tea';// 上传资源
var auth_resource__rt_resource_upload='rt_resource_upload';// 删除资源
var auth_resource__rt_resource_del='rt_resource_del';// 使用到导学本
var auth_resource__rt_use_to_guide_book='rt_use_to_guide_book';// 发到校本资源
var auth_resource__rt_publish_to_school='rt_publish_to_school';// 下载资源
var auth_resource__rt_resource_download='rt_resource_download';/** ---------------------------------------- 备课中心(校本) ---------------------------------------- */ // 校本备课中心模块
var auth_resource_school='resource_school_tea';// 上传资源
var auth_resource_school__rst_resource_upload='rst_resource_upload';// 删除资源
var auth_resource_school__rst_resource_del='rst_resource_del';// 使用到导学本
var auth_resource_school__rst_use_to_guide_book='rst_use_to_guide_book';// 下载资源
var auth_resource_school__rst_resource_download='rst_resource_download';/** ---------------------------------------- 备课中心(公共) ---------------------------------------- */ // 公共备课中心模块
var auth_resource_public='resource_common_tea';// 下载资源
var auth_resource_public__rct_resource_download='rct_resource_download';// 使用到导学本
var auth_resource_public__rct_use_to_guide_book='rct_use_to_guide_book';/** ---------------------------------------- 题库(个人) ---------------------------------------- */ // 个人题库模块
var auth_questionbank='questionbank_tea';// 新增题目
var auth_questionbank__qt_question_add='qt_question_add';// 批量导入
var auth_questionbank__qt_question_batch_import='qt_question_batch_import';// 删除题目
var auth_questionbank__qt_question_del='qt_question_del';// 试题篮
var auth_questionbank__qt_question_basket='qt_question_basket';// 编辑题目
var auth_questionbank__qt_question_edit='qt_question_edit';// 解析答案
var auth_questionbank__qt_analyze_answer='qt_analyze_answer';// 看/录讲解
var auth_questionbank__qt_explanation_read_and_record='qt_explanation_read/record';// 加入校本题库
var auth_questionbank__qt_add_to_school_questionbank='qt_add_to_school_questionbank';/** ---------------------------------------- 题库(校本) ---------------------------------------- */ // 校本题库模块
var auth_questionbank_school='questionbank_school_tea';// 新增题目
var auth_questionbank_school__qst_question_add='qst_question_add';// 批量导入
var auth_questionbank_school__qst_question_batch_import='qst_question_batch_import';// 删除题目
var auth_questionbank_school__qst_question_del='qst_question_del';// 试题篮
var auth_questionbank_school__qst_question_basket='qst_question_basket';// 编辑题目
var auth_questionbank_school__qst_question_edit='qst_question_edit';// 解析答案
var auth_questionbank_school__qst_analyze_answer='qst_analyze_answer';// 看讲解
var auth_questionbank_school__qst_explanation_read='qst_explanation_read';// 录讲解
var auth_questionbank_school__qst_explanation_record='qst_explanation_record';/** ---------------------------------------- 题库(公共) ---------------------------------------- */ // 公共题库模块
var auth_questionbank_public='questionbank_common_tea';// 试题篮
var auth_questionbank_public__qct_question_basket='qct_question_basket';// 看解析
var auth_questionbank_public__qct_analyze_answer='qct_analyze_answer';/** ---------------------------------------- 智慧课堂 ---------------------------------------- */ // 智慧课堂模块
var auth_intelligent_classroom='intelligent_classroom';// 课堂报告
var auth_intelligent_classroom__classroom_report='classroom_report';/** ---------------------------------------- 课本 ---------------------------------------- */ // 课本模块
var auth_textbook='textbook_tea';// 个人课本 (添加、修改、删除)
var auth_textbook__tt_personal_textbook='tt_personal_textbook';// 校本课本 (添加、修改、删除)
var auth_textbook__tt_school_textbook='tt_school_textbook';/** ---------------------------------------- 个人中心 ---------------------------------------- */ // 班级设置
var auth_usercenter__class_setting='class_setting';/** ---------------------------------------- 使用向导 ---------------------------------------- */ // 使用向导模块
var auth_user_guidance='user_guidance_tea';/** ---------------------------------------- 入学考试 ---------------------------------------- */ // 入学考试模块
var auth_exam='exam_school_tea';/** ---------------------------------------- 录播平台 ---------------------------------------- */ // 录播平台模块
var auth_record_play='record_play_tea';
// CONCATENATED MODULE: ./src/auth/page.ts
// 模块权限
/** 智慧课堂权限 */var CLASSROOM_AUTH=[auth_intelligent_classroom];/** 作业辅导权限 */var COACH_AUTH=[auth_coach];/** 入学考试权限 */var EXAM_AUTH=[auth_exam];/** 导学本权限 */var GUIDE_AUTH=[auth_guide];/** 题库权限 */var QUESTIONS_AUTH=[auth_questionbank];/** 备课中心权限 */var RESOURCE_AUTH=[auth_resource];/** 统计分析权限 */var STATISTICS_AUTH=[auth_statistics];/** 课本权限 */var TEXTBOOK_AUTH=[auth_textbook];/** 使用向导权限 */var WIZARD_AUTH=[auth_user_guidance];/** 网络作业权限 */var WORK_AUTH=[auth_homework];/** 错题集权限 */var WRONG_AUTH=[auth_mistakecollection];
// CONCATENATED MODULE: ./src/routes/pages/classroom.tsx
var TOP_NAV_NAME='智慧课堂';var routes=[{name:'classroom',path:'/classroom',component:layouts_Base,meta:{hanName:TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(TOP_NAV_NAME),topNavName:TOP_NAV_NAME,auth:CLASSROOM_AUTH},children:[{"default":true,path:'/list',name:'classroom-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | classroom-list */ 0).then(__webpack_require__.bind(null, 471))}));},meta:{hanName:"".concat(TOP_NAV_NAME,"\u5217\u8868"),title:"\u6559\u5B66\u5E73\u53F0-".concat(TOP_NAV_NAME),topNavName:TOP_NAV_NAME,auth:CLASSROOM_AUTH}}]}];/* harmony default export */ var classroom = (routes);
// CONCATENATED MODULE: ./src/routes/pages/coach.tsx
var coach_TOP_NAV_NAME='作业辅导';var coach_routes=[{name:'coach',path:'/coach',component:layouts_Base,meta:{hanName:coach_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(coach_TOP_NAV_NAME),topNavName:coach_TOP_NAV_NAME,auth:COACH_AUTH},children:[{"default":true,path:'/list',name:'coach-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | coach_list */ 1).then(__webpack_require__.bind(null, 472))}));},meta:{hanName:coach_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(coach_TOP_NAV_NAME),topNavName:coach_TOP_NAV_NAME,auth:COACH_AUTH}}]}];/* harmony default export */ var coach = (coach_routes);
// CONCATENATED MODULE: ./src/routes/pages/exam.tsx
var exam_TOP_NAV_NAME='入学考试';var exam_routes=[{name:'exam',path:'/exam',component:layouts_Base,meta:{hanName:exam_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(exam_TOP_NAV_NAME),topNavName:exam_TOP_NAV_NAME,auth:EXAM_AUTH},children:[{"default":true,path:'/list',name:'exam-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | exam-list */ 2).then(__webpack_require__.bind(null, 473))}));},meta:{hanName:exam_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(exam_TOP_NAV_NAME),topNavName:exam_TOP_NAV_NAME,auth:EXAM_AUTH}}]}];/* harmony default export */ var exam = (exam_routes);
// CONCATENATED MODULE: ./src/routes/pages/guide.tsx
var guide_TOP_NAV_NAME='导学本';var guide_routes=[{name:'guide',path:'/guide',component:layouts_Base,meta:{hanName:guide_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(guide_TOP_NAV_NAME),topNavName:guide_TOP_NAV_NAME,auth:GUIDE_AUTH},children:[{"default":true,path:'/list',name:'guide-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | guide-list */ 3).then(__webpack_require__.bind(null, 474))}));},meta:{hanName:guide_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(guide_TOP_NAV_NAME),topNavName:guide_TOP_NAV_NAME,auth:GUIDE_AUTH}}]}];/* harmony default export */ var guide = (guide_routes);
// CONCATENATED MODULE: ./src/routes/pages/questions.tsx
var questions_TOP_NAV_NAME='题库';var questions_routes=[{name:'questions',path:'/questions',component:layouts_Base,meta:{hanName:questions_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(questions_TOP_NAV_NAME),topNavName:questions_TOP_NAV_NAME,auth:QUESTIONS_AUTH},children:[{"default":true,path:'/list',name:'questions-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | questions-list */ 8).then(__webpack_require__.bind(null, 475))}));},meta:{hanName:questions_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(questions_TOP_NAV_NAME),topNavName:questions_TOP_NAV_NAME,auth:QUESTIONS_AUTH}}]}];/* harmony default export */ var questions = (questions_routes);
// CONCATENATED MODULE: ./src/routes/pages/resource.tsx
var resource_TOP_NAV_NAME='备课中心';var resource_routes=[{name:'resource',path:'/resource',component:layouts_Base,meta:{hanName:resource_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(resource_TOP_NAV_NAME),topNavName:resource_TOP_NAV_NAME,auth:RESOURCE_AUTH},children:[{"default":true,path:'/list',name:'resource-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | resource-list */ 9).then(__webpack_require__.bind(null, 476))}));},meta:{hanName:resource_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(resource_TOP_NAV_NAME),topNavName:resource_TOP_NAV_NAME,auth:RESOURCE_AUTH}}]}];/* harmony default export */ var resource = (resource_routes);
// CONCATENATED MODULE: ./src/routes/pages/statistics.tsx
var statistics_TOP_NAV_NAME='统计分析';var statistics_routes=[{name:'statistics',path:'/statistics',component:layouts_Base,meta:{hanName:statistics_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(statistics_TOP_NAV_NAME),topNavName:statistics_TOP_NAV_NAME,auth:STATISTICS_AUTH},children:[{"default":true,name:'resource-view',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | resource-view */ 10).then(__webpack_require__.bind(null, 477))}));},meta:{hanName:statistics_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(statistics_TOP_NAV_NAME),topNavName:statistics_TOP_NAV_NAME,auth:STATISTICS_AUTH}}]}];/* harmony default export */ var statistics = (statistics_routes);
// CONCATENATED MODULE: ./src/routes/pages/textbook.tsx
var textbook_TOP_NAV_NAME='课本管理';var textbook_routes=[{name:'textbook',path:'/textbook',component:layouts_Base,meta:{hanName:textbook_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(textbook_TOP_NAV_NAME),topNavName:textbook_TOP_NAV_NAME,auth:TEXTBOOK_AUTH},children:[{"default":true,name:'textbook-chapter',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | textbook-chapter */ 13).then(__webpack_require__.bind(null, 478))}));},meta:{hanName:textbook_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(textbook_TOP_NAV_NAME),topNavName:textbook_TOP_NAV_NAME,auth:TEXTBOOK_AUTH}}]}];/* harmony default export */ var textbook = (textbook_routes);
// CONCATENATED MODULE: ./src/routes/pages/wizard.tsx
var wizard_TOP_NAV_NAME='使用向导';var wizard_routes=[{name:'wizard',path:'/wizard',component:layouts_Base,meta:{hanName:wizard_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(wizard_TOP_NAV_NAME),topNavName:wizard_TOP_NAV_NAME,auth:WIZARD_AUTH},children:[{"default":true,name:'wizard-wk',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | wizard-wk */ 15).then(__webpack_require__.bind(null, 479))}));},meta:{hanName:wizard_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(wizard_TOP_NAV_NAME),topNavName:wizard_TOP_NAV_NAME,auth:WIZARD_AUTH}}]}];/* harmony default export */ var wizard = (wizard_routes);
// CONCATENATED MODULE: ./src/routes/pages/work.tsx
var work_TOP_NAV_NAME='网络作业';var work_routes=[{name:'work',path:'/work',component:layouts_Base,meta:{hanName:work_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(work_TOP_NAV_NAME),topNavName:work_TOP_NAV_NAME,auth:WORK_AUTH},children:[{"default":true,path:'/list',name:'work-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | work-list */ 16).then(__webpack_require__.bind(null, 482))}));},meta:{hanName:work_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(work_TOP_NAV_NAME),topNavName:work_TOP_NAV_NAME,auth:WORK_AUTH}}]}];/* harmony default export */ var work = (work_routes);
// CONCATENATED MODULE: ./src/routes/pages/wrong.tsx
var wrong_TOP_NAV_NAME='错题集';var wrong_routes=[{name:'wrong',path:'/wrong',component:layouts_Base,meta:{hanName:wrong_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(wrong_TOP_NAV_NAME),topNavName:wrong_TOP_NAV_NAME,auth:WRONG_AUTH},children:[{"default":true,path:'/list',name:'wrong-list',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | wrong-list */ 17).then(__webpack_require__.bind(null, 480))}));},meta:{hanName:wrong_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(wrong_TOP_NAV_NAME),topNavName:wrong_TOP_NAV_NAME,auth:WRONG_AUTH}}]}];/* harmony default export */ var wrong = (wrong_routes);
// CONCATENATED MODULE: ./src/routes/pages/settings.tsx
var settings_TOP_NAV_NAME='个人设置';var settings_routes=[{name:'settings',path:'/settings',component:layouts_Base,meta:{hanName:settings_TOP_NAV_NAME,title:"\u6559\u5B66\u5E73\u53F0-".concat(settings_TOP_NAV_NAME),topNavName:settings_TOP_NAV_NAME,auth:[]},children:[{"default":true,name:'settings-default',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | settings */ 12).then(__webpack_require__.bind(null, 481))}));},meta:{hanName:"".concat(settings_TOP_NAV_NAME),title:"\u6559\u5B66\u5E73\u53F0-".concat(settings_TOP_NAV_NAME),topNavName:settings_TOP_NAV_NAME,auth:[]}}]}];/* harmony default export */ var settings = (settings_routes);
// CONCATENATED MODULE: ./src/routes/routes.tsx
var routes_routes=[{name:'login',path:'/login',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:Promise.all(/* import() | login */[__webpack_require__.e(19), __webpack_require__.e(5)]).then(__webpack_require__.bind(null, 468))}));},meta:{title:'登录',hanName:'登录'}},{name:'home',path:'/',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | home */ 4).then(__webpack_require__.bind(null, 483))}));},meta:{title:'首页',hanName:'首页'}}].concat(Object(toConsumableArray["a" /* default */])(classroom),Object(toConsumableArray["a" /* default */])(coach),Object(toConsumableArray["a" /* default */])(exam),Object(toConsumableArray["a" /* default */])(guide),Object(toConsumableArray["a" /* default */])(questions),Object(toConsumableArray["a" /* default */])(resource),Object(toConsumableArray["a" /* default */])(statistics),Object(toConsumableArray["a" /* default */])(textbook),Object(toConsumableArray["a" /* default */])(wizard),Object(toConsumableArray["a" /* default */])(work),Object(toConsumableArray["a" /* default */])(wrong),Object(toConsumableArray["a" /* default */])(settings),[{name:'404',path:'/404',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | notfound */ 7).then(__webpack_require__.bind(null, 469))}));},meta:{title:'404',hanName:'页面未找到'}},{name:'unauthorized',path:'/unauthorized',component:function component(props){return reactfrom_dll_reference_main_library_default.a.createElement(suspense,Object.assign({},props,{load:__webpack_require__.e(/* import() | unauthorized */ 14).then(__webpack_require__.bind(null, 470))}));},meta:{title:'unauthorized',hanName:'无权限访问'}}]);
// EXTERNAL MODULE: ./node_modules/react-router/es/Redirect.js + 1 modules
var Redirect = __webpack_require__(302);

// EXTERNAL MODULE: ./src/auth/token.ts + 1 modules
var token = __webpack_require__(20);

// EXTERNAL MODULE: ./src/auth/authority.ts
var authority = __webpack_require__(43);

// CONCATENATED MODULE: ./src/routes/middleware/auth.tsx
// 权限相关中间件
/**
 * 判断是否登录
 * whiteList: 不验证是否登录的路由 name 的白名单列表
 */function loginRequiredMiddleware(whiteList){return function(state){// 路由名在白名单中，不需要验证是否登录
if(whiteList&&whiteList.length){if(whiteList.indexOf(state.name)!==-1){return null;}}if(token["a" /* default */].isLogin()){return null;}return reactfrom_dll_reference_main_library_default.a.createElement(Redirect["a" /* default */],{to:"/login"});};}/**
 * 未登录才可访问
 * checkList: 该列表中的 route name 的路由未登录才可访问
 */function guestMiddleware(checkList){return function(state){if(checkList&&checkList.length){if(checkList.indexOf(state.name)===-1){return null;}}if(!token["a" /* default */].isLogin()){return null;}return reactfrom_dll_reference_main_library_default.a.createElement(Redirect["a" /* default */],{to:"/"});};}/** 用于权限判断的中间件 (也会判断是否登录了) */function authCheckMiddleware(){return function(state){var meta=state.meta||{};var auths=meta.auth||[];var strict=Object(instance["a" /* isUndefined */])(meta.authUnstrict)?true:!meta.authUnstrict;// 权限检查 (有配置才需检查)
if(!auths.length){return null;}if(authority["a" /* default */].checkAll(auths,strict)){return null;}return reactfrom_dll_reference_main_library_default.a.createElement(Redirect["a" /* default */],{to:"/unauthorized"});};}
// CONCATENATED MODULE: ./src/routes/index.tsx
// 加载路由配置
var awRouter=lib_default.a.instance();awRouter.load(routes_routes,{// history
history:Object(esm_history["createHashHistory"])(),// not found route name
notFoundRouteName:'404',// 全局路由中间件
middlewares:[// 设置 title
function(state){var meta=state.meta||{};var routeTitle=meta.title;if(routeTitle){document.title=routeTitle;}else{document.title=['教学平台',meta.hanName].filter(Boolean).join('-');}return null;},config["d" /* PATCH_CALLBACK_ROUTE_MIDDLEWARE */],// 暴露在 MAIN_CONFIG 中用于调试或打补丁的中间件
loginRequiredMiddleware(['404','unauthorized','login']),// 这些路由不需验证是否登录
guestMiddleware(['login']),// 只能未登录才可访问
authCheckMiddleware()]});var routes_MyRouter=function MyRouter(){return reactfrom_dll_reference_main_library_default.a.createElement(HashRouter["a" /* default */],null,awRouter.render()());};/* harmony default export */ var src_routes = (Object(reactfrom_dll_reference_main_library["memo"])(routes_MyRouter));
// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/classCallCheck.js
var classCallCheck = __webpack_require__(12);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/createClass.js
var createClass = __webpack_require__(9);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/possibleConstructorReturn.js + 2 modules
var possibleConstructorReturn = __webpack_require__(154);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/getPrototypeOf.js
var getPrototypeOf = __webpack_require__(153);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/inherits.js + 1 modules
var inherits = __webpack_require__(155);

// CONCATENATED MODULE: ./src/components/ErrorBoundary/index.tsx
// 错误边界
var ErrorBoundary_ErrorBoundary=/*#__PURE__*/function(_React$Component){Object(inherits["a" /* default */])(ErrorBoundary,_React$Component);function ErrorBoundary(props){var _this;Object(classCallCheck["a" /* default */])(this,ErrorBoundary);_this=Object(possibleConstructorReturn["a" /* default */])(this,Object(getPrototypeOf["a" /* default */])(ErrorBoundary).call(this,props));_this.state={hasError:false};_this.state={hasError:false};return _this;}Object(createClass["a" /* default */])(ErrorBoundary,[{key:"componentDidCatch",value:function componentDidCatch(error,errorInfo){console.error("ErrorBoundary: ",{error:error,errorInfo:errorInfo});this.setState({hasError:true});}},{key:"render",value:function render(){var _this$props=this.props,children=_this$props.children,ErrorComponent=_this$props.ErrorComponent;if(this.state.hasError){if(ErrorComponent){return reactfrom_dll_reference_main_library_default.a.createElement(ErrorComponent,null);}return reactfrom_dll_reference_main_library_default.a.createElement("h1",null,"Something went wrong.");}return children;}}],[{key:"getDerivedStateFromError",value:function getDerivedStateFromError(error){console.error("ErrorBoundary: ",error);return{hasError:true};}}]);return ErrorBoundary;}(reactfrom_dll_reference_main_library_default.a.Component);/* harmony default export */ var components_ErrorBoundary = (ErrorBoundary_ErrorBoundary);
// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/regenerator/index.js
var regenerator = __webpack_require__(8);
var regenerator_default = /*#__PURE__*/__webpack_require__.n(regenerator);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/asyncToGenerator.js
var asyncToGenerator = __webpack_require__(18);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/slicedToArray.js + 3 modules
var slicedToArray = __webpack_require__(16);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/objectSpread.js + 1 modules
var objectSpread = __webpack_require__(35);

// EXTERNAL MODULE: ./node_modules/p-min-delay/index.js
var p_min_delay = __webpack_require__(89);
var p_min_delay_default = /*#__PURE__*/__webpack_require__.n(p_min_delay);

// CONCATENATED MODULE: ./src/tools/hooks/async.ts
var ActionType;(function(ActionType){ActionType["START"]="START";ActionType["FINISH"]="FINISH";ActionType["HAS_ERROR"]="HAS_ERROR";})(ActionType||(ActionType={}));function promiseReducer(state,action){switch(action.type){case ActionType.START:return Object(objectSpread["a" /* default */])({},state,{initial:false,pending:true,errorMessage:'',hasError:false});case ActionType.FINISH:return Object(objectSpread["a" /* default */])({},state,{initial:false,pending:false},action.payload);case ActionType.HAS_ERROR:return Object(objectSpread["a" /* default */])({},state,{initial:false,pending:false},action.payload);default:return state;}}var startActionCreator=function startActionCreator(){return{type:ActionType.START};};var finishActionCreator=function finishActionCreator(result){return{type:ActionType.FINISH,payload:{hasError:false,errorMessage:'',result:result}};};var hasErrorActionCreator=function hasErrorActionCreator(){var errorMessage=arguments.length>0&&arguments[0]!==undefined?arguments[0]:'';var result=arguments.length>1&&arguments[1]!==undefined?arguments[1]:null;return{type:ActionType.HAS_ERROR,payload:{hasError:true,errorMessage:errorMessage,result:result}};};function usePromise(func,initialState){var _ref=arguments.length>2&&arguments[2]!==undefined?arguments[2]:{},delay=_ref.delay,checkError=_ref.checkError;var _useReducer=Object(reactfrom_dll_reference_main_library["useReducer"])(promiseReducer,{initial:true,pending:true,hasError:false,errorMessage:'',result:initialState}),_useReducer2=Object(slicedToArray["a" /* default */])(_useReducer,2),state=_useReducer2[0],dispatch=_useReducer2[1];var run=Object(reactfrom_dll_reference_main_library["useCallback"])(/*#__PURE__*/Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee(){var resp,e;return regenerator_default.a.wrap(function _callee$(_context){while(1){switch(_context.prev=_context.next){case 0:dispatch(startActionCreator());_context.prev=1;if(!delay){_context.next=8;break;}_context.next=5;return p_min_delay_default()(func(),delay);case 5:resp=_context.sent;_context.next=11;break;case 8:_context.next=10;return func();case 10:resp=_context.sent;case 11:if(!checkError){_context.next=16;break;}e=checkError(resp);if(!e.hasError){_context.next=16;break;}dispatch(hasErrorActionCreator(e.errorMessage,resp));return _context.abrupt("return");case 16:dispatch(finishActionCreator(resp));_context.next=23;break;case 19:_context.prev=19;_context.t0=_context["catch"](1);dispatch(hasErrorActionCreator(String(_context.t0)));console.warn('usePromise error: ',_context.t0);case 23:case"end":return _context.stop();}}},_callee,null,[[1,19]]);})),[delay,checkError]);return Object(reactfrom_dll_reference_main_library["useMemo"])(function(){return Object(objectSpread["a" /* default */])({},state,{run:run});},[state,run]);}function useRequest(func,initialState){var _ref3=arguments.length>2&&arguments[2]!==undefined?arguments[2]:{},delay=_ref3.delay,checkError=_ref3.checkError;var check=Object(reactfrom_dll_reference_main_library["useMemo"])(function(){if(!checkError){return function(result){if(!result.status){return{hasError:true,errorMessage:result.message||''};}return{hasError:false,errorMessage:''};};}return checkError;},[]);var pr=usePromise(func,{status:true,message:'',data:initialState},{delay:delay,checkError:check});return Object(objectSpread["a" /* default */])({},pr,{result:pr.result.data});}
// CONCATENATED MODULE: ./src/App.tsx
var App_App=function App(){var _AuthStore$useStore=auth["a" /* default */].useStore(),getUserInfoAction=_AuthStore$useStore.getUserInfoAction;var _useRequest=useRequest(getUserInfoAction,{}),pending=_useRequest.pending,run=_useRequest.run;Object(reactfrom_dll_reference_main_library["useEffect"])(function(){run();},[]);return reactfrom_dll_reference_main_library_default.a.createElement(components_ErrorBoundary,null,pending?reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"common_page_placeholder"},"loading ..."):reactfrom_dll_reference_main_library_default.a.createElement(src_routes,null));};/* harmony default export */ var src_App = (Object(aw_store["withStoreProviders"])(App_App,[auth["a" /* default */]]));
// CONCATENATED MODULE: ./src/main.tsx
// ---------------------------------------- css ----------------------------------------
// ---------------------------------------- main ----------------------------------------
// ---------------------------------------- 其他 ----------------------------------------
// ---------------------------------------- 错误处理 ----------------------------------------
window.onerror=function(event,source,fileno,columnNumber,error){var error_data={title:document.getElementsByTagName('title')[0].innerText,error_event:event,script_url:source,line_number:fileno,column_number:columnNumber,error_detail:error&&error.message||'',error_stack:error&&error.stack||'',user_agent:window.navigator.userAgent,location_href:window.location.href,token:token["a" /* default */].get()};var event_string=String(error_data.error_event||'');if(event_string==='ResizeObserver loop limit exceeded'){return;}var SCRIPT_ERROR='Unexpected token <';// 脚本错误 (一般是服务器中的前端代码更新了，浏览器缓存中的 js 代码由于不一致而报错)
var CHUNK_ERROR='chunk';if(event_string.indexOf(SCRIPT_ERROR)!==-1||event_string.indexOf(CHUNK_ERROR)!==-1){message_default.a.warning('发现新版本，请使用刷新网页');}};// ---------------------------------------- 挂载 ----------------------------------------
fastclick_default.a.attach(document.body);// ---------------------------------------- token ----------------------------------------
window.addEventListener('storage',function(ev){var key=ev.key;var val=ev.newValue;if(key===token["a" /* default */].key&&val===null){token["a" /* default */].clean();}});// logout 时的 callback
token["a" /* default */].logoutCallback=function(){window.localStorage.clear();window.sessionStorage.clear();};// ======================================= RENDER =======================================
function createApp(){return reactfrom_dll_reference_main_library_default.a.createElement(src_App,null);}function main_render(){react_domfrom_dll_reference_main_library_default.a.render(createApp(),document.getElementById(constants["a" /* APP_ROOT_ID */]));}if(Object(config["g" /* isDev */])()){// DEVELOPMENT
var devRender=function devRender(){if(false){}main_render();};try{devRender();}catch(error){console.error(error);main_render();}}else{// PRODUCTION
main_render();}

/***/ }),

/***/ 31:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var _type__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(105);
/* harmony import */ var _type__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_type__WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (checked) */ if(__webpack_require__.o(_type__WEBPACK_IMPORTED_MODULE_0__, "createStore")) __webpack_require__.d(__webpack_exports__, "createStore", function() { return _type__WEBPACK_IMPORTED_MODULE_0__["createStore"]; });

/* harmony reexport (checked) */ if(__webpack_require__.o(_type__WEBPACK_IMPORTED_MODULE_0__, "withStoreProviders")) __webpack_require__.d(__webpack_exports__, "withStoreProviders", function() { return _type__WEBPACK_IMPORTED_MODULE_0__["withStoreProviders"]; });

/* harmony import */ var _with__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(106);
/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "withStoreProviders", function() { return _with__WEBPACK_IMPORTED_MODULE_1__["a"]; });

/* harmony import */ var _store__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(107);
/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "createStore", function() { return _store__WEBPACK_IMPORTED_MODULE_2__["a"]; });



/***/ }),

/***/ 32:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* unused harmony export isPromise */
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return isUndefined; });
/** 判断是否为 promise */function isPromise(val){return val&&typeof val.then==='function';}/** 判断是否未 undefined */function isUndefined(val){return val===void 0;}

/***/ }),

/***/ 33:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";

// CONCATENATED MODULE: ./src/constants/app.ts
/** react 挂载的 dom root id */var APP_ROOT_ID='root';
// CONCATENATED MODULE: ./src/constants/pages.ts
var PAGE_NAMES=[{name:'resource'},{name:'coach'},{name:'guide'},{name:'wrong'},{name:'classroom'},{name:'work'},{name:'statistics'},{name:'questions'},{name:'exam'}];
// CONCATENATED MODULE: ./src/constants/user.ts
// 用户相关常量
// 用户类型
var USER_TYPE;(function(USER_TYPE){USER_TYPE[USER_TYPE["SUPER_ADMIN"]=0]="SUPER_ADMIN";USER_TYPE[USER_TYPE["ADMIN"]=1]="ADMIN";USER_TYPE[USER_TYPE["TEACHER"]=2]="TEACHER";USER_TYPE[USER_TYPE["STUDENT"]=3]="STUDENT";})(USER_TYPE||(USER_TYPE={}));/**
 * 角色类型
 * - 1: 教师 TEACHER
 * - 2: 班主任 HEAD_TEACHER
 * - 3: 学生 STUDENTS
 * - 4: 校长 PRINCIPAL
 * - 5: 年级主任 GRADE_DIRECTOR
 * - 6: 科目组长 SUBJECT_LEADER
 */var ROLE_TYPE;// 是否为默认角色
(function(ROLE_TYPE){ROLE_TYPE[ROLE_TYPE["TEACHER"]=1]="TEACHER";ROLE_TYPE[ROLE_TYPE["HEAD_TEACHER"]=2]="HEAD_TEACHER";ROLE_TYPE[ROLE_TYPE["STUDENTS"]=3]="STUDENTS";ROLE_TYPE[ROLE_TYPE["PRINCIPAL"]=4]="PRINCIPAL";ROLE_TYPE[ROLE_TYPE["GRADE_DIRECTOR"]=5]="GRADE_DIRECTOR";ROLE_TYPE[ROLE_TYPE["SUBJECT_LEADER"]=6]="SUBJECT_LEADER";})(ROLE_TYPE||(ROLE_TYPE={}));var ROLE_STATUS;(function(ROLE_STATUS){ROLE_STATUS[ROLE_STATUS["IS"]=1]="IS";ROLE_STATUS[ROLE_STATUS["NO"]=0]="NO";})(ROLE_STATUS||(ROLE_STATUS={}));
// CONCATENATED MODULE: ./src/constants/index.ts
/* concated harmony reexport APP_ROOT_ID */__webpack_require__.d(__webpack_exports__, "a", function() { return APP_ROOT_ID; });
/* concated harmony reexport PAGE_NAMES */__webpack_require__.d(__webpack_exports__, "b", function() { return PAGE_NAMES; });
/* concated harmony reexport USER_TYPE */__webpack_require__.d(__webpack_exports__, "d", function() { return USER_TYPE; });
/* unused concated harmony import ROLE_TYPE */
/* concated harmony reexport ROLE_STATUS */__webpack_require__.d(__webpack_exports__, "c", function() { return ROLE_STATUS; });


/***/ }),

/***/ 40:
/***/ (function(module, exports) {

module.exports = main_library;

/***/ }),

/***/ 43:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_classCallCheck__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(12);
/* harmony import */ var C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_createClass__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(9);
// 权限模块
var Authority=/*#__PURE__*/function(){Object(C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_createClass__WEBPACK_IMPORTED_MODULE_1__[/* default */ "a"])(Authority,null,[{key:"instance",value:function instance(opt){if(this._instance===null){this._instance=new this(opt||{});}return this._instance;}/** 不判断权限 (用于调试) */}]);function Authority(_ref){var _ref$mock=_ref.mock,mock=_ref$mock===void 0?false:_ref$mock;Object(C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_classCallCheck__WEBPACK_IMPORTED_MODULE_0__[/* default */ "a"])(this,Authority);this.mock=void 0;this.loaded=false;this.authStore=[];this.mock=mock;}/** 是否已加载过权限 */Object(C_Users_wutong_Desktop_react_yj_node_modules_babel_preset_react_app_node_modules_babel_runtime_helpers_esm_createClass__WEBPACK_IMPORTED_MODULE_1__[/* default */ "a"])(Authority,[{key:"load",/** 加载权限 */value:function load(){var authArr=arguments.length>0&&arguments[0]!==undefined?arguments[0]:[];if(this.loaded){return;}this.authStore=authArr;this.loaded=true;}/** 加载权限 (map 形式) */},{key:"loadMap",value:function loadMap(){var authMap=arguments.length>0&&arguments[0]!==undefined?arguments[0]:{};if(this.loaded){return;}var authArr=[];for(var key in authMap){if(authMap.hasOwnProperty(key)){var item=authMap[key];authArr.push(key);authArr=authArr.concat(item);}}this.load(authArr);}/** 重新加载权限 */},{key:"reload",value:function reload(){var authArr=arguments.length>0&&arguments[0]!==undefined?arguments[0]:[];this.authStore=authArr;this.loaded=true;}/** 检查权限 */},{key:"check",value:function check(auth){if(!this.loaded){throw new Error('Authority: 还未加载权限数据');}if(this.mock){return true;}if(!this.authStore.length){return false;}return this.authStore.indexOf(auth)!==-1;}/**
   * 检查权限
   * strict: true，auths 中必须全部存在才通过，否则一个通过即可
   */},{key:"checkAll",value:function checkAll(auths){var _this=this;var strict=arguments.length>1&&arguments[1]!==undefined?arguments[1]:true;if(strict){return auths.every(function(a){return _this.check(a);});}return auths.some(function(a){return _this.check(a);});}}]);return Authority;}();Authority._instance=null;var instance=Authority.instance({mock:false});/* harmony default export */ __webpack_exports__["a"] = (instance);window.__Authority__=instance;

/***/ }),

/***/ 51:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var react_hanger__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(69);
/* harmony import */ var _tools_aw_aw_store__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(31);
// base layout 的一些状态
var initialState={breadcrumbs:[]};/** 设置 layout 的面包屑 */function setBreadcrumbsAction(state){return function(breadcrumbs){state.setState({breadcrumbs:breadcrumbs});};}function useBaseLayout(){var state=Object(react_hanger__WEBPACK_IMPORTED_MODULE_0__[/* useSetState */ "a"])(initialState);return{state:state.state,// actions
setBreadcrumbsAction:setBreadcrumbsAction(state)};}/* harmony default export */ __webpack_exports__["a"] = (Object(_tools_aw_aw_store__WEBPACK_IMPORTED_MODULE_1__["createStore"])(useBaseLayout));

/***/ }),

/***/ 53:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/classCallCheck.js
var classCallCheck = __webpack_require__(12);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/createClass.js
var createClass = __webpack_require__(9);

// CONCATENATED MODULE: ./src/tools/aw/aw-storage/index.ts
var SEPARATOR='$META|META$';// 不验证是否过期的过期时间设置
var NO_VALIDATE_EXPIRE=0;var aw_storage_AWStorage=/*#__PURE__*/function(){/** 是否禁用缓存 */ /** 开启 log */ /** 前缀 */ /** 是否使用 sessionStorage 存储 (默认使用 localStorage 存储) */ /** 是否有 meta 信息 */ /** 过期时间: 单位毫秒 (0 为不过期) */ // 一分钟为 1000 * 60，一小时为 1000 * 60 * 60
/** 判断是否过期的函数 (会和 expire 一块使用) */function AWStorage(_ref){var _ref$disabled=_ref.disabled,disabled=_ref$disabled===void 0?false:_ref$disabled,_ref$enableLog=_ref.enableLog,enableLog=_ref$enableLog===void 0?false:_ref$enableLog,_ref$prefix=_ref.prefix,prefix=_ref$prefix===void 0?'':_ref$prefix,_ref$session=_ref.session,session=_ref$session===void 0?false:_ref$session,_ref$expire=_ref.expire,expire=_ref$expire===void 0?0:_ref$expire,_ref$hasMeta=_ref.hasMeta,hasMeta=_ref$hasMeta===void 0?true:_ref$hasMeta,checkExpire=_ref.checkExpire;Object(classCallCheck["a" /* default */])(this,AWStorage);this.disabled=void 0;this.enableLog=void 0;this.prefix=void 0;this.session=void 0;this.hasMeta=void 0;this.expire=void 0;this.checkExpire=void 0;this.disabled=disabled;this.enableLog=enableLog;this.prefix=prefix;this.session=session;this.hasMeta=hasMeta;this.expire=expire;this.checkExpire=checkExpire;}Object(createClass["a" /* default */])(AWStorage,[{key:"log",value:function log(key,action){var msg=arguments.length>2&&arguments[2]!==undefined?arguments[2]:'';if(!this.enableLog){return;}console.log('AWStorage-'+action+''+key,msg);}},{key:"getStorage",value:function getStorage(){if(!window.sessionStorage||!window.localStorage){throw new Error('该浏览器不支持本地存储');}return this.session?window.sessionStorage:window.localStorage;}/** 获取 key */},{key:"getKey",value:function getKey(key){if(!this.prefix){return key;}return this.prefix+key;}/** 创建 meta 信息 */},{key:"createMeta",value:function createMeta(){var defaultMeta={updateTime:new Date().getTime(),expire:this.expire};return JSON.stringify(defaultMeta);}/** 是否过期 false=未过期 */},{key:"isExpired",value:function isExpired(meta){// 没有过期信息
if(!meta.expire||!meta.updateTime||meta.expire===NO_VALIDATE_EXPIRE){return false;}// 执行配置的验证函数
if(this.checkExpire&&this.checkExpire(meta.expire)){return true;}// 验证是否过期
var now=new Date().getTime();var time=now-meta.updateTime;return time>=meta.expire;}/** 获取 meta 信息，并且返回真正的 value */},{key:"parseMetaAndReturnRealValue",value:function parseMetaAndReturnRealValue(key){var rawValue=this.getStorage().getItem(this.getKey(key))||'';var valueArr=rawValue.split(SEPARATOR);// 没有 meta 信息
if(valueArr.length===1){return rawValue;}var metaString=valueArr[0];var realValue=valueArr[1];if(!realValue){return null;}try{var metaObj=JSON.parse(metaString);// meta 信息验证
// 1. 是否过期
if(this.isExpired(metaObj)){this.remove(key);return null;}return realValue;}catch(err){console.warn(err);return null;}}},{key:"setItem",value:function setItem(key,val){if(this.disabled){return;}if(this.hasMeta){var meta=this.createMeta();var value=meta+SEPARATOR+String(val);this.getStorage().setItem(this.getKey(key),value);}else{this.getStorage().setItem(this.getKey(key),String(val));}this.log(key,'SET');}/** 返回空字符串 (表示没获取到) */},{key:"getItem",value:function getItem(key){if(this.disabled){return null;}var value='';if(this.hasMeta){value=this.parseMetaAndReturnRealValue(key);}else{value=this.getStorage().getItem(this.getKey(key))||'';}this.log(key,'GET');return value||'';}// ------------------- storeage -------------------
/** 设置 string */},{key:"setString",value:function setString(key,val){this.setItem(key,val);}/** 获取 string */},{key:"getString",value:function getString(key){return this.getItem(key);}/** 设置 number */},{key:"setNumber",value:function setNumber(key,val){this.setItem(key,String(val));}/** 获取 number */},{key:"getNumber",value:function getNumber(key){var value=this.getItem(key);return value===null?null:Number(value);}/** 设置 json */},{key:"setJson",value:function setJson(key,val,middleware){try{var data=JSON.stringify(val);var dealData=middleware?middleware(data):data;this.setItem(key,dealData);return true;}catch(err){console.warn(err);return false;}}/** 获取 json */},{key:"getJson",value:function getJson(key,middleware){var data=this.getItem(key);if(!data){return null;}var dealData=middleware?middleware(data):data;try{return dealData!==''?JSON.parse(dealData):null;}catch(err){console.warn(err);return null;}}/** 删除 */},{key:"remove",value:function remove(){var _this=this;for(var _len=arguments.length,keys=new Array(_len),_key=0;_key<_len;_key++){keys[_key]=arguments[_key];}if(!keys||!keys.length){return false;}keys.forEach(function(key){return _this.getStorage().removeItem(_this.getKey(key));});return true;}}]);return AWStorage;}();/* harmony default export */ var aw_storage = (aw_storage_AWStorage);
// EXTERNAL MODULE: ./src/config/index.ts
var config = __webpack_require__(21);

// CONCATENATED MODULE: ./src/storage/base.ts
var base_AWStorageFactory=function AWStorageFactory(_ref){var _ref$enableLog=_ref.enableLog,enableLog=_ref$enableLog===void 0?false:_ref$enableLog,_ref$disabled=_ref.disabled,disabled=_ref$disabled===void 0?false:_ref$disabled,_ref$prefix=_ref.prefix,prefix=_ref$prefix===void 0?config["c" /* LOCALSTROAGE_PREFIX */]:_ref$prefix,_ref$session=_ref.session,session=_ref$session===void 0?false:_ref$session,_ref$hasMeta=_ref.hasMeta,hasMeta=_ref$hasMeta===void 0?true:_ref$hasMeta,_ref$expire=_ref.expire,expire=_ref$expire===void 0?NO_VALIDATE_EXPIRE:_ref$expire,_ref$checkExpire=_ref.checkExpire,checkExpire=_ref$checkExpire===void 0?function(e){return e===config["b" /* DISABLED_EXPIRE */];}:_ref$checkExpire;return new aw_storage({enableLog:enableLog,disabled:disabled,session:session,prefix:prefix,hasMeta:hasMeta,expire:expire,checkExpire:checkExpire});};/* harmony default export */ var base = __webpack_exports__["a"] = (base_AWStorageFactory);

/***/ }),

/***/ 58:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/objectSpread.js + 1 modules
var objectSpread = __webpack_require__(35);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/regenerator/index.js
var regenerator = __webpack_require__(8);
var regenerator_default = /*#__PURE__*/__webpack_require__.n(regenerator);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/asyncToGenerator.js
var asyncToGenerator = __webpack_require__(18);

// EXTERNAL MODULE: delegated ../node_modules/react/index.js from dll-reference main_library
var reactfrom_dll_reference_main_library = __webpack_require__(1);

// EXTERNAL MODULE: ./node_modules/react-hanger/es6/index.js + 9 modules
var es6 = __webpack_require__(69);

// EXTERNAL MODULE: ./src/tools/aw/aw-store/index.tsx
var aw_store = __webpack_require__(31);

// EXTERNAL MODULE: ./src/constants/index.ts + 3 modules
var constants = __webpack_require__(33);

// EXTERNAL MODULE: ./src/auth/authority.ts
var authority = __webpack_require__(43);

// EXTERNAL MODULE: ./src/auth/token.ts + 1 modules
var auth_token = __webpack_require__(20);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/classCallCheck.js
var classCallCheck = __webpack_require__(12);

// EXTERNAL MODULE: ./node_modules/babel-preset-react-app/node_modules/@babel/runtime/helpers/esm/createClass.js
var createClass = __webpack_require__(9);

// EXTERNAL MODULE: ./node_modules/antd/lib/message/style/index.js
var style = __webpack_require__(100);

// EXTERNAL MODULE: ./node_modules/antd/lib/message/index.js
var message = __webpack_require__(50);
var message_default = /*#__PURE__*/__webpack_require__.n(message);

// EXTERNAL MODULE: ./node_modules/axios/index.js
var axios = __webpack_require__(148);
var axios_default = /*#__PURE__*/__webpack_require__.n(axios);

// EXTERNAL MODULE: ./node_modules/qs/lib/index.js
var lib = __webpack_require__(149);
var lib_default = /*#__PURE__*/__webpack_require__.n(lib);

// CONCATENATED MODULE: ./src/services/type.ts
var URLENCODED_CONTENT_TYPE={'Content-Type':'application/x-www-form-urlencoded;charset=UTF-8'};var JSON_CONTENT_TYPE={'Content-Type':'application/json;charset=UTF-8'};var CommonErrorMessage='网络错误';
// CONCATENATED MODULE: ./src/services/code.ts
// 成功
var OK_CODE=200;// 用户未登录
var NO_LOGIN_CODE=464;// 其他错误
var OTHER_ERROR_CODES=[599,// 加载失败
532,// 调用第三方错误
9999,// 服务不可用
-1];
// CONCATENATED MODULE: ./src/services/base.ts
var base_axios=axios_default.a.create({headers:URLENCODED_CONTENT_TYPE});// request
base_axios.interceptors.request.use(function(config){// 处理参数
if(!config.params){config.params={};}if(!config.data){config.data={};}var token=auth_token["a" /* default */].get();if(token){config.params.token=token;// query 上携带 token
}if(!config.json){config.data=lib_default.a.stringify(config.data);}else{config.headers=Object.assign({},config.headers||{},JSON_CONTENT_TYPE);}return config;},function(error){message_default.a.error(error);return Promise.reject(error);});// response
base_axios.interceptors.response.use(function(response){// 响应数据具体在 services 中处理
return response;},function(error){message_default.a.error(error);return Promise.reject(error);});var base_Services=/*#__PURE__*/function(){function Services(){Object(classCallCheck["a" /* default */])(this,Services);}Object(createClass["a" /* default */])(Services,[{key:"errorResponseData",value:function errorResponseData(config,result){var msg=result&&result.msg?result.msg:CommonErrorMessage;if(config.showErrorMessage){message_default.a.error(msg);}return{status:false,message:msg,data:result?result.data:null};}},{key:"responseData",value:function responseData(config,result){var statusCode=Number(result.code);// 如有携带 token，存储
if(result.token){auth_token["a" /* default */].set(result.token);}// 请求成功
if(statusCode===OK_CODE){return{status:true,message:result.msg||'成功',data:result.data};}// 请求失败 (code 处理)
if(statusCode===NO_LOGIN_CODE){auth_token["a" /* default */].logout({triggerGlobalCallback:true,reloadPage:true});return this.errorResponseData(config);}if(OTHER_ERROR_CODES.indexOf(statusCode)!==-1){return this.errorResponseData(config,result);}return this.errorResponseData(config,result);}},{key:"request",value:function(){var _request=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee(config){var resp;return regenerator_default.a.wrap(function _callee$(_context){while(1){switch(_context.prev=_context.next){case 0:_context.prev=0;_context.next=3;return base_axios(config);case 3:resp=_context.sent;return _context.abrupt("return",this.responseData(config,resp.data));case 7:_context.prev=7;_context.t0=_context["catch"](0);console.warn(_context.t0);return _context.abrupt("return",this.errorResponseData(config));case 11:case"end":return _context.stop();}}},_callee,this,[[0,7]]);}));function request(_x){return _request.apply(this,arguments);}return request;}()},{key:"get",value:function(){var _get=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee2(config){return regenerator_default.a.wrap(function _callee2$(_context2){while(1){switch(_context2.prev=_context2.next){case 0:return _context2.abrupt("return",this.request(Object.assign({},config,{method:'GET'})));case 1:case"end":return _context2.stop();}}},_callee2,this);}));function get(_x2){return _get.apply(this,arguments);}return get;}()},{key:"post",value:function(){var _post=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee3(config){return regenerator_default.a.wrap(function _callee3$(_context3){while(1){switch(_context3.prev=_context3.next){case 0:return _context3.abrupt("return",this.request(Object.assign({json:false},config,{method:'POST'})));case 1:case"end":return _context3.stop();}}},_callee3,this);}));function post(_x3){return _post.apply(this,arguments);}return post;}()}],[{key:"instance",value:function instance(){if(this._instance===null){this._instance=new this();}return this._instance;}}]);return Services;}();base_Services._instance=null;/* harmony default export */ var base = (base_Services.instance());window.__Services__=base_Services.instance();
// EXTERNAL MODULE: ./src/config/index.ts
var src_config = __webpack_require__(21);

// CONCATENATED MODULE: ./src/services/user/index.ts
var user_default=/*#__PURE__*/function(){function _default(){Object(classCallCheck["a" /* default */])(this,_default);}Object(createClass["a" /* default */])(_default,null,[{key:"authAPI",value:function authAPI(api){return src_config["a" /* AUTH_API_ROOT */]+api;}},{key:"login",/** 登录 */value:function(){var _login=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee(_ref){var username,password,resp;return regenerator_default.a.wrap(function _callee$(_context){while(1){switch(_context.prev=_context.next){case 0:username=_ref.username,password=_ref.password;_context.next=3;return base.get({url:this.authAPI('/efficiencyLogin'),params:{username:username,password:password},showErrorMessage:false});case 3:resp=_context.sent;return _context.abrupt("return",resp);case 5:case"end":return _context.stop();}}},_callee,this);}));function login(_x){return _login.apply(this,arguments);}return login;}()/** 获取用户信息 */},{key:"getUserInfo",value:function(){var _getUserInfo=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee2(){var resp;return regenerator_default.a.wrap(function _callee2$(_context2){while(1){switch(_context2.prev=_context2.next){case 0:_context2.next=2;return base.get({url:this.authAPI('/api/user/efficiency-select-one')});case 2:resp=_context2.sent;return _context2.abrupt("return",resp);case 4:case"end":return _context2.stop();}}},_callee2,this);}));function getUserInfo(){return _getUserInfo.apply(this,arguments);}return getUserInfo;}()/** 获取默认学期 */},{key:"getDetaultTerm",value:function(){var _getDetaultTerm=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee3(){var resp,_data,data;return regenerator_default.a.wrap(function _callee3$(_context3){while(1){switch(_context3.prev=_context3.next){case 0:_context3.next=2;return base.get({url:this.authAPI('/api/term/get-default-terminfo')});case 2:resp=_context3.sent;_data=resp.data||{};data={termid:_data.termid,termname:_data.yearname+'学年'+_data.termname};return _context3.abrupt("return",{status:resp.status,message:resp.message,data:data});case 6:case"end":return _context3.stop();}}},_callee3,this);}));function getDetaultTerm(){return _getDetaultTerm.apply(this,arguments);}return getDetaultTerm;}()}]);return _default;}();
// EXTERNAL MODULE: ./src/storage/base.ts + 1 modules
var storage_base = __webpack_require__(53);

// CONCATENATED MODULE: ./src/storage/term.ts
/** 学期 */var term_TermStorage=/*#__PURE__*/function(){Object(createClass["a" /* default */])(TermStorage,null,[{key:"instance",value:function instance(){if(this._instance===null){this._instance=new this();}return this._instance;}}]);function TermStorage(){Object(classCallCheck["a" /* default */])(this,TermStorage);this.storage=void 0;this.key='termid';this.storage=Object(storage_base["a" /* default */])({hasMeta:false});}Object(createClass["a" /* default */])(TermStorage,[{key:"clean",value:function clean(){return this.storage.remove(this.key);}},{key:"get",value:function get(){return this.storage.getNumber(this.key);}},{key:"set",value:function set(id){this.storage.setNumber(this.key,id);}}]);return TermStorage;}();term_TermStorage._instance=null;/* harmony default export */ var storage_term = (term_TermStorage.instance());window.__TermStorage__=term_TermStorage.instance();
// CONCATENATED MODULE: ./src/storage/role.ts
/** 角色 */var role_RoleStorage=/*#__PURE__*/function(){Object(createClass["a" /* default */])(RoleStorage,null,[{key:"instance",value:function instance(){if(this._instance===null){this._instance=new this();}return this._instance;}}]);function RoleStorage(){Object(classCallCheck["a" /* default */])(this,RoleStorage);this.storage=void 0;this.roleStore=null;this.key='role';this.storage=Object(storage_base["a" /* default */])({hasMeta:false});}Object(createClass["a" /* default */])(RoleStorage,[{key:"clean",value:function clean(){this.roleStore=null;return this.storage.remove(this.key);}},{key:"get",value:function get(){if(this.roleStore){return this.roleStore;}try{var jsonString=this.storage.getString(this.key);if(!jsonString){return null;}var data=JSON.parse(jsonString);if(!data.roleid||!data.roletype){return null;}var roleData={roleid:Number(data.roleid),roletype:Number(data.roletype)};this.roleStore=roleData;return roleData;}catch(err){console.warn(err);return null;}}},{key:"set",value:function set(role){this.roleStore=role;this.storage.setString(this.key,JSON.stringify(role));}},{key:"isTeacherRole",value:function isTeacherRole(){var role=this.get();if(!role){return true;// 默认老师
}// 1 为教师角色 (注意: 这里的 type 和用户 category 不一样，那个是判断学生老师管理员身份的)
return role.roletype===1;}}]);return RoleStorage;}();role_RoleStorage._instance=null;/* harmony default export */ var storage_role = (role_RoleStorage.instance());window.__RoleStorage__=role_RoleStorage.instance();
// CONCATENATED MODULE: ./src/store/auth/index.ts
// 用户相关 (登录登出、权限 ...)
var initialState={user:null,terms:[],currentTerm:null,currentRole:null};function setUserToState(data,state){var user=data.userEntity||{};var terms=data.termVos||[];authority["a" /* default */].loadMap(user.acl||{});state.setState({user:user,terms:terms});}/** 设置默认学期 */function setDefaultTerm(_x,_x2){return _setDefaultTerm.apply(this,arguments);}/** 设置默认角色 */function _setDefaultTerm(){_setDefaultTerm=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee3(state,defaultTermID){var terms,localTermID,localTerm,defaultTerm,term;return regenerator_default.a.wrap(function _callee3$(_context3){while(1){switch(_context3.prev=_context3.next){case 0:terms=state.state.terms||[];if(terms.length){_context3.next=4;break;}storage_term.clean();return _context3.abrupt("return");case 4:localTermID=storage_term.get();if(!localTermID){_context3.next=10;break;}localTerm=terms.find(function(t){return t.termid===localTermID;});if(!localTerm){_context3.next=10;break;}setCurrentTermAction(state)(localTerm.termid);return _context3.abrupt("return");case 10:if(!defaultTermID){_context3.next=15;break;}defaultTerm=terms.find(function(t){return t.termid===defaultTermID;});if(!defaultTerm){_context3.next=15;break;}setCurrentTermAction(state)(defaultTerm.termid);return _context3.abrupt("return");case 15:_context3.next=17;return user_default.getDetaultTerm();case 17:term=_context3.sent;setCurrentTermAction(state)(term.data.termid);case 19:case"end":return _context3.stop();}}},_callee3);}));return _setDefaultTerm.apply(this,arguments);}function setDefaultRole(state){var setRole=function setRole(roleid){var inList=roles.find(function(r){return r.roleid===roleid;});if(inList){setCurrentRoleAction(state)(roleid);return true;}return false;};if(!state.state.user){return;}var roles=state.state.user.aclrole||[];if(!roles.length){storage_role.clean();return;}var localRole=storage_role.get();if(localRole){if(setRole(localRole.roleid)){return;}}var defaultRole=roles.find(function(r){return r.activestatus===constants["c" /* ROLE_STATUS */].IS;});if(defaultRole){if(setRole(defaultRole.roleid)){return;}}var firstRole=roles[0];if(firstRole){setRole(firstRole.roleid);}}/** 设置学期 */function setCurrentTermAction(state){return function(termid){if(state.state.currentTerm&&state.state.currentTerm.termid===termid){return false;}var term=state.state.terms.find(function(t){return t.termid===termid;});if(!term){return false;}state.setState({currentTerm:term});storage_term.set(term.termid);return true;};}/** 设置角色 */function setCurrentRoleAction(state){return function(roleid){if(!state.state.user){return false;}if(state.state.currentRole&&state.state.currentRole.roleid===roleid){return false;}var roles=state.state.user.aclrole||[];var role=roles.find(function(r){return r.roleid===roleid;});if(!role){return false;}state.setState({currentRole:role});storage_role.set({roleid:role.roleid,roletype:role.type});return true;};}/** 登录 */function loginAction(state){return(/*#__PURE__*/function(){var _ref=Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee(username,password){var result;return regenerator_default.a.wrap(function _callee$(_context){while(1){switch(_context.prev=_context.next){case 0:_context.next=2;return user_default.login({username:username.trim(),password:password.trim()});case 2:result=_context.sent;if(result.status&&result.data){setUserToState(result.data,state);}return _context.abrupt("return",result);case 5:case"end":return _context.stop();}}},_callee);}));return function(_x3,_x4){return _ref.apply(this,arguments);};}());}/** 获取用户信息 */function getUserInfoAction(state){// force: 覆盖用户数据
return(/*#__PURE__*/Object(asyncToGenerator["a" /* default */])(/*#__PURE__*/regenerator_default.a.mark(function _callee2(){var force,result,_args2=arguments;return regenerator_default.a.wrap(function _callee2$(_context2){while(1){switch(_context2.prev=_context2.next){case 0:force=_args2.length>0&&_args2[0]!==undefined?_args2[0]:false;if(auth_token["a" /* default */].isLogin()){_context2.next=3;break;}return _context2.abrupt("return",{status:false,message:'用户未登录',data:null});case 3:if(!(!force&&state.state.user)){_context2.next=5;break;}return _context2.abrupt("return",{status:true,message:'成功',data:state.state.user});case 5:_context2.next=7;return user_default.getUserInfo();case 7:result=_context2.sent;if(result.status&&result.data){setUserToState(result.data,state);}return _context2.abrupt("return",result);case 10:case"end":return _context2.stop();}}},_callee2);})));}/** 登出 */function logoutAction(_ref3){var setState=_ref3.setState;return function(){setState(Object(objectSpread["a" /* default */])({},initialState));auth_token["a" /* default */].logout({triggerGlobalCallback:true,reloadPage:true});};}function useAuth(){var state=Object(es6["a" /* useSetState */])(initialState);Object(reactfrom_dll_reference_main_library["useEffect"])(function(){setDefaultTerm(state);},[state.state.terms]);Object(reactfrom_dll_reference_main_library["useEffect"])(function(){setDefaultRole(state);},[state.state.user]);return{state:state.state,// actions
loginAction:loginAction(state),getUserInfoAction:getUserInfoAction(state),logoutAction:logoutAction(state),setCurrentTermAction:setCurrentTermAction(state),setCurrentRoleAction:setCurrentRoleAction(state),// getters
/** 获取权限 table */authGetter:function authGetter(){return authority["a" /* default */].authStore||[];},/** 获取学期数据 */termsGetter:function termsGetter(){return state.state.terms||[];},/** 是否为老师 */isTeacherGetter:function isTeacherGetter(){return state.state.user?state.state.user.category===constants["d" /* USER_TYPE */].TEACHER:false;},/** 角色列表 */rolesGetter:function rolesGetter(){var info=state.state.user;if(!info){return[];}return info.aclrole||[];}};}/* harmony default export */ var auth = __webpack_exports__["a"] = (Object(aw_store["createStore"])(useAuth));

/***/ }),

/***/ 6:
/***/ (function(module, exports, __webpack_require__) {

module.exports = (__webpack_require__(40))(427);

/***/ }),

/***/ 92:
/***/ (function(module, exports) {

module.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEkAAAAmCAYAAACI0bZTAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyFpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQyIDc5LjE2MDkyNCwgMjAxNy8wNy8xMy0wMTowNjozOSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIChXaW5kb3dzKSIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDpENkMxM0IyRDczOTUxMUU4ODA0MkZDRUUwNjhGQkM5MyIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDpENkMxM0IyRTczOTUxMUU4ODA0MkZDRUUwNjhGQkM5MyI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOkQ2QzEzQjJCNzM5NTExRTg4MDQyRkNFRTA2OEZCQzkzIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOkQ2QzEzQjJDNzM5NTExRTg4MDQyRkNFRTA2OEZCQzkzIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+Wu5wuwAABntJREFUeNrcmmlsVFUU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"

/***/ })

},[[170,11,18]]]);
//# sourceMappingURL=main.ceb6fa8d.chunk.js.map