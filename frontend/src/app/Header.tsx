import { Box, Button, Spacer, Stack, Text } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { useEffect } from 'react'
import { useSignOut, useGetCurrentUser } from '@api/hooks'
import { CircleUserRound, LogOut, LogIn, ArrowLeft } from 'lucide-react'
import { useUser } from '@shared/provider/UserProvider'
export const Header = () => {
  const navigate = useNavigate()
  const { currentUser, setCurrentUser } = useUser()
  const { data: user, isError } = useGetCurrentUser()

  const { mutate: signoutMutate } = useSignOut({
    mutation: {
      onSuccess: () => {
        setCurrentUser(null)
        navigate('/signin')
      }
    }
  })

  useEffect(() => {
    setCurrentUser(user!)
  }, [user])

  useEffect(() => {
    if (isError) {
      setCurrentUser(null)
    }
  }, [isError])

  function initialUser({ user_name }: { user_name: string | undefined }) {
    if (user_name) {
      return user_name[0].toUpperCase()
    }
    return null
  }

  const handleGoBack = () => {
    navigate(-1)
  }

  return (
    <header className="app-header">
      <Button h="40px" w="40px" borderRadius="full" onClick={handleGoBack} bg="none">
        <Stack>
          <ArrowLeft color="#646464" size="30" />
        </Stack>
      </Button>
      {isError || !currentUser ? (
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
          <Button h="40px" w="40px" borderRadius="full" bg="#583474" color="White">
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
            }}
          >
            <LogOut size="24" color="#646464" />
          </Button>
        </>
      )}
    </header>
  )
}
