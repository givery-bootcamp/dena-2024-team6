import { Box, Button, Loading, Text } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { useSignOut, useGetCurrentUser } from '@api/hooks'

export const Header = () => {
  const navigate = useNavigate()
  const { data: user, isError, refetch, isFetching } = useGetCurrentUser()
  const [currentUser, setCurrentUser] = useState(user)

  const { mutate: signoutMutate } = useSignOut({
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
    <Box bgColor="white">
      サンプルアプリケーション
      <Box display="flex" alignItems="center">
        {isFetching ? (
          <Loading variant="circles" size="6xl" color="cyan.500" />
        ) : isError || !currentUser ? (
          <Button onClick={() => navigate('/signin')}>サインイン</Button>
        ) : (
          <>
            <Text mr={2}>{currentUser.user_name}</Text>
            <Button
              onClick={() => {
                signoutMutate()
                refetch()
              }}
            >
              サインアウト
            </Button>
          </>
        )}
      </Box>
    </Box>
  )
}
