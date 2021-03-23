import React from 'react'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk'
import { createStore , applyMiddleware } from 'redux'
import './App.css'
import Dashboard from './components/Dashboard/Dashboard.js'
import rootReducer from './reducers/index.js'
import Nav from './components/Nav/Nav.js'

const endpoint = process.env.REACT_APP_ENDPOINT

const store = createStore(rootReducer, applyMiddleware(thunk))

function App() {


  return (
    <Provider store={store}>
      <div className="App">
        <Dashboard endpoint={endpoint}/>
        <div>{process.env.REACT_APP_ENDPOINT}</div>
      </div>
    </Provider>
  )
}

export default App
