import { Box, Button, Center, Loading, Spacer, Text } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { useSignOut, useGetCurrentUser } from '@api/hooks'
import { CircleUserRound, LogOut, LogIn } from 'lucide-react'

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

  function initialUser({ user_name }: { user_name: string | undefined }) {
    if (user_name) {
      return user_name[0].toUpperCase()
    }
    return null
  }

  return (
    <header className="app-header">
      {isFetching ? (
        <Center>
          <Loading variant="circles" size="6xl" color="#98C9DE" />
        </Center>
      ) : isError || !currentUser ? (
        <>
          <CircleUserRound size="40" />
          <Box w="10px" />
          <Text fontWeight="bold" fontFamily="Inter" fontSize="16px">
            匿名ユーザ
          </Text>
          <Spacer />
          <Button bg="none" onClick={() => navigate('/signin')}>
            <LogIn size="20" color="#646464" />
            <Text fontFamily="Inter" fontSize="12px" color="#646464">
              ログイン
            </Text>
          </Button>
        </>
      ) : (
        <>
          <Button h="40px" w="20px" borderRadius="full" bg="#583474" color="White">
            {initialUser({ user_name: currentUser.user_name })}
          </Button>
          <Box w="10px" />
          <Text fontWeight="bold" fontFamily="Inter" fontSize="16px">
            {currentUser.user_name}
          </Text>
          <Spacer />
          <Button
            bg="none"
            onClick={() => {
              signoutMutate()
              refetch()
            }}
          >
            <LogOut size="20" color="#646464" />
          </Button>
        </>
      )}
    </header>
  )
}
