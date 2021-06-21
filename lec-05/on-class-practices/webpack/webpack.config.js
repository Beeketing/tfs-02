const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: path.resolve(__dirname, "src/index.js"),
  output: {
    path: path.resolve(__dirname, "dist"),
    filename: "bundle.js",
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        use: "babel-loader",
      },
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"],
      },
    ],
  },
  devServer: {
    port: 8081,
    contentBase: path.resolve(__dirname, "public"),
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: "Hello world",
      template: path.resolve(__dirname, "public/index.html"),
    }),
  ],
};
