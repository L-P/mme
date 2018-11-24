const CleanWebpackPlugin = require('clean-webpack-plugin');

module.exports = {
  lintOnSave: false,
  baseUrl: undefined,
  outputDir: '../static',
  assetsDir: undefined,
  runtimeCompiler: undefined,
  productionSourceMap: false,
  parallel: undefined,
  css: undefined,
  configureWebpack: {
    plugins: [
      new CleanWebpackPlugin(
        [`${__dirname}/../static`],
        { root: `${__dirname}/..` },
      ),
    ],
  },
};
