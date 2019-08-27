;(function (name, definition, context) {
  if ((typeof module !== 'undefined') && module.exports)
    module.exports = definition() // node 环境
  else if ((typeof context['define'] === 'function') && (context['define']['amd'] || context['define']['cmd']))
    define(definition)            // amd cmd 规范环境，如 seajs requirejs
  else
    context[name] = definition()  // 浏览器环境
})('MAIN_CONFIG', function () {

  var api_root = window.API_ROOT;
  var url_root = window.URL_ROOT;

  return {
    // ---------------------------- api prefix ----------------------------
    API_ROOT: api_root,
    URL_ROOT: url_root,

    // ---------------------------- key ----------------------------
    LOCALSTROAGE_PREFIX: 'aw_',      // localstroage prefix
    TOKEN_KEY:           'aw_token', // token key

    // ---------------------------- other ----------------------------
    PATCH_CALLBACK: function () {}
  }

}, this);
