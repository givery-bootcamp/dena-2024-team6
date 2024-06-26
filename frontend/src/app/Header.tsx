import { Box, Button, Text } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { usePostSignout, useGetUser } from '../api/api'

export const Header = () => {
  const navigate = useNavigate()
  const { data: user, isError } = useGetUser()
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
    if (isError) {
      setCurrentUser(undefined)
    } else {
      setCurrentUser(user)
    }
  }, [user, isError])

  return (
    <header className="app-header">
      サンプルアプリケーション
      <Box display="flex" alignItems="center">
        {isError || !currentUser ? (
          <Button onClick={() => signoutMutate()}>サインイン</Button>
        ) : (
          <>
            <Text mr={2}>{currentUser.user_name}</Text>
            <Button onClick={() => signoutMutate()}>サインアウト</Button>
          </>
        )}
      </Box>
    </header>
  )
}
