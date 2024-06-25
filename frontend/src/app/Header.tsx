import { Box, Button } from '@yamada-ui/react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

export const Header = () => {
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
    <header className="app-header">
      サンプルアプリケーション
      <Box display="flex" alignItems="center">
        {isAuthenticated && <Box mr={2}>ユーザー</Box>}
        <Button variant="solid" onClick={handleButtonClick}>
          {isAuthenticated ? 'サインアウト' : 'サインイン'}
        </Button>
      </Box>
    </header>
  )
}
