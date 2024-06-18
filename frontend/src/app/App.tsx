import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { AppRoute } from './AppRoute'
import { Button, Spacer, Box } from '@yamada-ui/react'

import './App.scss'

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const navigate = useNavigate()

  const handleButtonClick = () => {
    if (isAuthenticated) {
      // サインアウト処理
      console.log('Signout 処理')
      setIsAuthenticated(false)
    } else {
      // サインイン処理
      console.log('Signin 処理')
      navigate('/signin')
      setIsAuthenticated(true)
    }
  }

  return (
    <div className="app-root">
      <header className="app-header">
        サンプルアプリケーション
        <Box display="flex" alignItems="center">
          {isAuthenticated && <Box mr={2}>ユーザー</Box>}
          <Button variant="solid" onClick={handleButtonClick}>
            {isAuthenticated ? 'サインアウト' : 'サインイン'}
          </Button>
        </Box>
      </header>
      <main className="app-body container">
        <AppRoute />
      </main>
    </div>
  )
}

export default App
