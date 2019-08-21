const path = require('path');
const webpack = require('webpack');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;
const paths = require('./paths');

module.exports = {
  mode: 'production',
  entry: {
    main: [
      'react-app-polyfill/ie9',
      'react-app-polyfill/stable',
      'react',
      'react-dom',
    ]
  },
  output: {
    path: paths.appPublic + '/dll',
    filename: '[name].dll.js',
    library: '[name]_library'
  },
  plugins: [
    new webpack.DllPlugin({
      path: path.join(__dirname, '.', '[name]-dll-manifest.json'),
      name: '[name]_library',
      context: paths.appPublic,
    }),
    new BundleAnalyzerPlugin({
      // analyzerMode: 'static'
    }),
  ]
};
