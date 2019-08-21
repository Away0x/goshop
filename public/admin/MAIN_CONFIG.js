;(function (name, definition, context) {
  if ((typeof module !== 'undefined') && module.exports)
    module.exports = definition() // node 环境
  else if ((typeof context['define'] === 'function') && (context['define']['amd'] || context['define']['cmd']))
    define(definition)            // amd cmd 规范环境，如 seajs requirejs
  else
    context[name] = definition()  // 浏览器环境
})('MAIN_CONFIG', function () {

  var api_root = window.API_ROOT || 'http://www.ischool365.com:10892';

  return {
    // ---------------------------- api prefix ----------------------------
    API_ROOT: api_root,
    AUTH_API_ROOT: api_root + '/auth',
    WORK_API_ROOT: api_root + '/homework',
    QUESTION_API_ROOT: api_root + '/tiku',
    COACH_API_ROOT: api_root + '/coach',
    OFFICE_API_ROOT: api_root + '/office',
    GUIDE_API_ROOT: api_root + '/daoxue',
    RESOURCE_API_ROOT: api_root + '/resource',
    DICTIONARY_API_ROOT: api_root + '/word',

    // ---------------------------- page url ----------------------------
    ADMIN_URL: '/admin-s', // 管理后台 url

    // ---------------------------- key ----------------------------
    LOCALSTROAGE_PREFIX: 'aw_',      // localstroage prefix
    TOKEN_KEY:           'aw_token', // token key

    // ---------------------------- cache config ----------------------------
    WK_STORAGE_EXPIRE: 1000 * 60 * 60,  // 微课大师本地缓存的过期时间 (1小时过期)
    DISABLED_CHAPTER_TREE_CACHE: false, // 是否禁用树缓存 (章节树第一层节点的内存缓存)
    DISABLED_KNOW_TREE_CACHE: false,    // 是否禁用树缓存 (知识点树的内存缓存)

    // ---------------------------- third party ----------------------------
    RECORDING_PLATFORM: 'http://10.10.10.100/auth.html', // 录播平台登录页面

    // ---------------------------- other ----------------------------
    // 补丁
    PATCH_CALLBACK: function (route) {}
  }

}, this);
