
import "core-js/stable"
import "regenerator-runtime/runtime"
import React from 'react'
import ReactDOM from 'react-dom'
import thunk from 'redux-thunk';
import {Provider} from 'react-redux'
import {composeWithDevTools} from 'redux-devtools-extension';
import {createStore, applyMiddleware} from 'redux'

import './index.css'
import App from './App'
import rootReducer from './reducers'

const store = createStore(rootReducer, composeWithDevTools(applyMiddleware(thunk)));

ReactDOM.render(<Provider store={store}><App/></Provider>, document.getElementById('root'))
