import { AppRoute } from './AppRoute'
import { Header } from './Header'
import './App.scss'
import { UserProvider } from '../shared/hooks/UserProvider'

function App() {
  return (
    <div className="app-root">
      <UserProvider>
        <Header />
        <main className="app-body container">
          <AppRoute />
        </main>
      </UserProvider>
    </div>
  )
}

export default App
