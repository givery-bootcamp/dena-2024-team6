import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'

import './main.scss'
import App from './app/App'
import { store } from './shared/store'
import { UIProvider } from '@yamada-ui/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
const queryClient = new QueryClient()

root.render(
  <BrowserRouter>
    <QueryClientProvider client={queryClient}>
      <UIProvider>
        <Provider store={store}>
          <App />
        </Provider>
      </UIProvider>
    </QueryClientProvider>
  </BrowserRouter>
)
