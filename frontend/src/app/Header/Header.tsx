import { useEffect, useState } from 'react'
import { Box, Button, Loading, Text, useDisclosure } from '@yamada-ui/react'
import { useNavigate } from 'react-router-dom'
import { usePostSignout, useGetUser } from '../../api/api'
import { SignoutModal } from './SignoutModal'

export const Header = () => {
  const navigate = useNavigate()
  const { data: user, isError, refetch, isFetching } = useGetUser()
  const [currentUser, setCurrentUser] = useState(user)
  const { isOpen, onOpen, onClose } = useDisclosure()

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
    <>
      <header className="app-header">
        サンプルアプリケーション
        <Box display="flex" alignItems="center">
          {isFetching ? (
            <Loading variant="circles" size="6xl" color="cyan.500" />
          ) : isError || !currentUser ? (
            <Button onClick={() => signoutMutate()}>サインイン</Button>
          ) : (
            <>
              <Text mr={2}>{currentUser.user_name}</Text>
              <Button onClick={onOpen}>サインアウト</Button>
            </>
          )}
        </Box>
      </header>
      <SignoutModal isOpen={isOpen} onClose={onClose} signoutMutate={signoutMutate} refetch={refetch} />
    </>
  )
}
