
var packageJSON = require('./package.json');
var path = require('path');
var webpack = require('webpack');
module.exports ={
    devtool: 'source-map',
    entry: {
        index: path.join(__dirname, 'web/index.jsx'),
    },
    // mode: "development",
    output: {
        path: path.join(__dirname, 'web/public/generated'),
        filename: 'app-bundle.js'
    },
    resolve: {extensions: ['.js', '.jsx']},
    plugins: [
        new webpack.LoaderOptionsPlugin({debug: false}),
        new webpack.DefinePlugin({
            'process.env.NODE_ENV' : JSON.stringify('development')
            // 'process.env.NODE_ENV' : JSON.stringify('production')
        })
    ],
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                exclude: /(node_modules|bower_components)/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.js$/,
                enforce: 'pre',
                use: ['source-map-loader'],
            },
            {
                test: /\.m?js$/,
                exclude: /(node_modules|bower_components)/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.css$/,
                use: [
                    {
                        loader: "css-loader"
                    }
                ]
            },
            {
                test: /\.s[ac]ss$/i,
                include: [
                    path.resolve(__dirname, 'node_modules'),
                    path.resolve(__dirname, 'web')
                ],
                use: [
                    // Creates `style` nodes from JS strings
                    "style-loader",
                    // Translates CSS into CommonJS
                    "css-loader",
                    // Compiles Sass to CSS
                    // "sass-loader",
                ],
            },
            {
                test: /\.(png|woff|woff2|eot|}ttf|svg)$/,
                use: [
                    {
                        loader: 'url-loader, options: { limit: 100000 } }]'
                    }
                ]
            }
        ]
    },
    devServer: {
        noInfo: false,
        quiet: false,
        lazy: false,
        watchOptions: {
            poll: true
        }
    }
}
