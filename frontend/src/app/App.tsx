import { AppRoute } from './AppRoute'
import { Box } from '@yamada-ui/react'
import './App.scss'
import { Header } from './Header'

function App() {
  return (
    <Box w="100vw" minH="100vh" bgColor="#98C9DE">
      <Header />
      <div className="app-body">
        <AppRoute />
      </div>
    </Box>
  )
}

export default App
