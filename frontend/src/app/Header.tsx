import { Box, Button, Text } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { usePostSignout } from '../api/api'
import { useUser } from '../shared/hooks/UserProvider'

export const Header = () => {
  const navigate = useNavigate()
  const { user, setUser } = useUser()
  const { mutate: signoutMutate } = usePostSignout({
    mutation: {
      onSuccess: () => {
        setUser(undefined)
        navigate('/signin')
      }
    }
  })

  const handleButtonClick = () => {
    if (user) {
      signoutMutate()
    } else {
      navigate('/signin')
    }
  }

  return (
    <header className="app-header">
      サンプルアプリケーション
      <Box display="flex" alignItems="center">
        {user && <Text mr={2}>{user.user_name}</Text>}
        <Button variant="solid" onClick={handleButtonClick}>
          {user ? 'サインアウト' : 'サインイン'}
        </Button>
      </Box>
    </header>
  )
}
