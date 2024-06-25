import { Box, Button } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { getUser } from '../api/api'

export const Header = () => {
  // const currentUser = getUser()
  const navigate = useNavigate()

  // console.log({ currentUser })

  // const handleButtonClick = () => {
  //   if (currentUser !== null) {
  //     // サインアウト処理
  //     console.log('SignOut 処理')
  //   } else {
  //     // サインイン処理
  //     console.log('Signin 処理')
  //     navigate('/signin')
  //   }
  // }

  return (
    <header className="app-header">
      サンプルアプリケーション
      {/* <Box display="flex" alignItems="center">
        {currentUser !== null && <Box mr={2}>ユーザー</Box>}
        <Button variant="solid" onClick={handleButtonClick}>
          {currentUser !== null ? 'サインアウト' : 'サインイン'}
        </Button>
      </Box> */}
    </header>
  )
}
