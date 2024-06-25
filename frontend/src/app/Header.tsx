import { Box, Button } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { usePostSignout, useGetUser } from '../api/api'

export const Header = () => {
  const { data: user, isLoading, isError } = useGetUser()
  const navigate = useNavigate()

  const signoutMutation = usePostSignout({
    mutation: {
      onSuccess: () => {
        navigate('/signin')
      }
    }
  })

  const handleButtonClick = () => {
    if (user) {
      signoutMutation.mutate()
    } else {
      navigate('/signin')
    }
  }

  return (
    <header className="app-header">
      サンプルアプリケーション
      <Box display="flex" alignItems="center">
        {/* {isLoading && <div>Loading...</div>} */}
        {/* {isError && <div>ユーザー情報の取得に失敗しました。</div>} */}
        {user && <Box mr={2}>{user.user_name}</Box>}
        {/* {!isLoading && !isError && ( */}
        <Button variant="solid" onClick={handleButtonClick}>
          {user ? 'サインアウト' : 'サインイン'}
        </Button>
        {/* } */}
      </Box>
    </header>
  )
}
