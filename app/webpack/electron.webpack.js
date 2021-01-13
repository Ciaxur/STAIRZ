const path = require('path');

module.exports = {
  entry: path.resolve(__dirname, '..', 'src', 'app.ts'),
  devtool: 'source-map',
  resolve: {
    extensions: [ '.tsx', '.ts', '.js' ],
  },
  target: 'electron-main',
  module: {
    rules: [
      {
        test: /\.(js|ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        },
      },
    ],
  },
  node: {
    __dirname: false,
  },
  output: {
    path: path.resolve(__dirname, '..', 'dist'),
    filename: 'app.js',
  },
};