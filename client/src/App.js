import React from 'react'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk'
import { createStore , applyMiddleware } from 'redux'
import './App.css'
import Dashboard from './components/Dashboard/Dashboard.js'
import rootReducer from './reducers/index.js'


const ep = process.env.REACT_APP_ENDPOINT

const store = createStore(rootReducer, applyMiddleware(thunk))

function App() {


  return (
    <Provider store={store}>
      <div className="App">
        <Dashboard endpoint={ep}/>
        <div>{process.env.REACT_APP_ENDPOINT}</div>
      </div>
    </Provider>
  )
}

export default App
