const webpack = require('webpack')
const merge = require('webpack-merge')
const common = require('./webpack.config.common.js')

module.exports = merge.smart(common, {
  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify('production')
      }
    })
  ]
})