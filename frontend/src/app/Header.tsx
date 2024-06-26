import { Box, Button, Text } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { usePostSignout, useGetUser } from '../api/api'

export const Header = () => {
  const navigate = useNavigate()
  const { data: user, isError, refetch } = useGetUser()
  const [currentUser, setCurrentUser] = useState(user)

  const { mutate: signoutMutate } = usePostSignout({
    mutation: {
      onSuccess: () => {
        setCurrentUser(undefined)
        navigate('/signin')
      }
    }
  })

  useEffect(() => {
    if (!isError) {
      setCurrentUser(user)
    } else {
      setCurrentUser(undefined)
    }
  }, [user, isError])

  const handleButtonClick = () => {
    if (currentUser) {
      signoutMutate()
    } else {
      navigate('/signin')
    }
  }

  return (
    <header className="app-header">
      サンプルアプリケーション
      <Box display="flex" alignItems="center">
        {currentUser && <Text mr={2}>{currentUser.user_name}</Text>}
        <Button variant="solid" onClick={handleButtonClick}>
          {currentUser ? 'サインアウト' : 'サインイン'}
        </Button>
      </Box>
    </header>
  )
}
