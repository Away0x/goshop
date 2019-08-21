(window["webpackJsonp"] = window["webpackJsonp"] || []).push([[16],{

/***/ 482:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);

// EXTERNAL MODULE: delegated ../node_modules/react/index.js from dll-reference main_library
var reactfrom_dll_reference_main_library = __webpack_require__(1);
var reactfrom_dll_reference_main_library_default = /*#__PURE__*/__webpack_require__.n(reactfrom_dll_reference_main_library);

// EXTERNAL MODULE: ./src/layouts/Base/hooks.ts
var hooks = __webpack_require__(142);

// EXTERNAL MODULE: ./node_modules/antd/lib/row/style/index.js
var style = __webpack_require__(99);

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

// CONCATENATED MODULE: ./src/components/ChapterKnowTree/TreeTypeSwitch.tsx
var TreeTypeSwitch_TreeTypeSwitch=function TreeTypeSwitch(){return reactfrom_dll_reference_main_library_default.a.createElement(row_default.a,{className:"chapter_know_tree__tree_type_switch"},reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:12},"\u7AE0\u8282"),reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:12},"\u77E5\u8BC6\u70B9"));};/* harmony default export */ var ChapterKnowTree_TreeTypeSwitch = (TreeTypeSwitch_TreeTypeSwitch);
// CONCATENATED MODULE: ./src/components/ChapterKnowTree/index.tsx
// import {
//   Row,
//   Col,
// } from 'antd';
// import MyTree from '@/components/MyTree';
// enum TreeType {
//   CHAPTER = 'chapter', // 章节树
//   KNOW = 'know'        // 知识点树
// }
var defaultProps={height:620,hasKnowTree:true,hasFilter:true};var ChapterKnowTree_ChapterKnowTree=function ChapterKnowTree(){return reactfrom_dll_reference_main_library_default.a.createElement("div",{className:"chapter_know_tree_component"},reactfrom_dll_reference_main_library_default.a.createElement(ChapterKnowTree_TreeTypeSwitch,null));};/* harmony default export */ var components_ChapterKnowTree = (Object(hoc["a" /* withDefaultProps */])(defaultProps,Object(reactfrom_dll_reference_main_library["memo"])(ChapterKnowTree_ChapterKnowTree)));
// CONCATENATED MODULE: ./src/layouts/LeftTree/index.tsx
var LeftTree_LeftTreeLayout=function LeftTreeLayout(props){return reactfrom_dll_reference_main_library_default.a.createElement(row_default.a,{className:"left_tree_layout"},reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:4},reactfrom_dll_reference_main_library_default.a.createElement(components_ChapterKnowTree,null)),reactfrom_dll_reference_main_library_default.a.createElement(col_default.a,{span:20},props.children));};/* harmony default export */ var LeftTree = (LeftTree_LeftTreeLayout);
// CONCATENATED MODULE: ./src/pages/Work/List/index.tsx
var List_WorkListPage=function WorkListPage(){Object(hooks["a" /* useBreadcrumbs */])([{name:'网络作业'}]);return reactfrom_dll_reference_main_library_default.a.createElement(LeftTree,null,reactfrom_dll_reference_main_library_default.a.createElement("div",null,"work list"));};/* harmony default export */ var List = __webpack_exports__["default"] = (Object(reactfrom_dll_reference_main_library["memo"])(List_WorkListPage));

/***/ })

}]);
//# sourceMappingURL=work-list.995b001f.chunk.js.map