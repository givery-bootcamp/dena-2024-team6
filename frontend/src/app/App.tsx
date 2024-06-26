import { AppRoute } from './AppRoute'
import { Header } from './Header'
import './App.scss'

function App() {
  return (
    <div className="app-root">
      <Header />
      <main className="app-body container">
        <AppRoute />
      </main>
    </div>
  )
}

export default App
