import { AppRoute } from './AppRoute'
import { Box } from '@yamada-ui/react'
import './App.scss'

function App() {
  return (
    <Box w="100vw" minH="100vh" bgColor="#98C9DE">
      <div className="app-body">
        <AppRoute />
      </div>
    </Box>
  )
}

export default App
