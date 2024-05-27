import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'

import './main.scss'
import App from './app/App'
import { store } from './shared/store'
import { UIProvider } from '@yamada-ui/react'

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
root.render(
  <BrowserRouter>
    <UIProvider>
      <Provider store={store}>
        <App />
      </Provider>
    </UIProvider>
  </BrowserRouter>
)
