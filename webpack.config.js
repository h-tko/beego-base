var webpack = require('webpack');
var path = require('path');

module.exports = {
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
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: 'file-loader',
        }, {
            test: /\.js[x]?$/,
            exclude: /node_modules/,
            loader: 'babel',
            query: {
                presets: ['react', 'es2015']
            }
        }],
    },
    devtool: '#source-map',
    plugins: [
        new webpack.ProvidePlugin({
            'React': 'react',
            'ReactDOM': 'react-dom',
        }),
    ],
}
