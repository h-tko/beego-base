var webpack = require('webpack');
var path = require('path');
var HtmlwebpackPlugin = require('html-webpack-plugin');
var merge = require('webpack-merge');

const ENV = process.env.npm_lifecycle_event;
process.env.BABEL_ENV = ENV;

const common = {
    entry: {
        'bundle': path.join(__dirname, 'static/js/app.js'),
    },
    output: {
        path: path.join(__dirname, 'static/dist'),
        filename: "[name].js",
    },
    cache: true,
    module: {
        loaders: [{
            test: /\.css/, loader: 'style-loader!css-loader',
        }, {
            test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: 'url-loader?limit=10000&mimetype=application/font-woff',
        }, {
            test: /\.(jpg|png)$/,
            loader: 'url-loader',
        }, {
            test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
            loader: 'url-loader?mimetype=image/svg+xml',
        }, {
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: 'file-loader',
        }, {
            test: /\.js?$/,
            exclude: /node_modules/,
            loader: 'babel-loader',
            query: {
                presets: ['react', 'es2015']
            }
        }],
    },
    plugins: [
        new webpack.ProvidePlugin({
            jQuery: 'jquery',
            $: 'jquery',
            Tether: 'tether',
            React: 'react',
            ReactDOM: 'react-dom',
        }),
        new HtmlwebpackPlugin({
            title: 'beego-base',
        }),
    ],
};

if (ENV === 'start' || !ENV) {
    module.exports = merge(common, {
        devtool: '#source-map',
        devServer: {
            contentBase: path.join(__dirname, 'static/dist'),
            publicPath: '/static/',
            historyApiFallback: false,
            hot: true,
            inline: true,
            host: "0.0.0.0",
            port: 3000,
            proxy: {
                '**': {
                    target: 'http://localhost:8080',
                    secure: false,
                },
            },
        },
        plugins: [
            new webpack.HotModuleReplacementPlugin()
        ]
    });
} else if (ENV === 'build') {
    module.exports =common;
}
