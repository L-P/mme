const CleanWebpackPlugin = require('clean-webpack-plugin');

module.exports = {
  lintOnSave: false,
  baseUrl: '/',
  outputDir: 'dist',
  assetsDir: './_',
  runtimeCompiler: undefined,
  productionSourceMap: false,
  parallel: undefined,
  css: undefined,
  configureWebpack: {
    plugins: [
      new CleanWebpackPlugin(['dist']),
    ],
  },
};
