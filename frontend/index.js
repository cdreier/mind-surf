import 'react-hot-loader/patch'
import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'mobx-react'
import { AppContainer } from 'react-hot-loader'
import App from './src/App'

import TreeModel from './src/models/TreeModel'

const models = {
  tree: new TreeModel(),
}


renderWithHotReload(App)

if (module.hot) {
  module.hot.accept('./src/App', () => {
    const NextRoot = require('./src/App.js').default
    renderWithHotReload(NextRoot)
  })
}

function renderWithHotReload(RootElement) {
  ReactDOM.render(
    <AppContainer>
      <Provider {...models}>
        <RootElement />
      </Provider>
    </AppContainer>,
    document.getElementById('mount')
  )
}
