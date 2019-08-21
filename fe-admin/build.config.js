const path = require('path')

module.exports = {
  entry: 'src/main',
  alias: {
    '@': path.join(__dirname, 'src'),
  },
  // 修改 ant design less 样式中的变量
  // antd-mobile https://github.com/ant-design/ant-design-mobile/blob/master/components/style/themes/default.less
  antDesignLessModifyVars: {
    'brand-primary': 'red'
  },
}