import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  Box,
  Button,
  Center,
  Container,
  Divider,
  FormControl,
  Input,
  Label,
  Snacks,
  Text,
  useSnacks
} from '@yamada-ui/react'
import { useSignIn, useGetCurrentUser } from '@api/hooks'

export const SigninRoute = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const { snack, snacks } = useSnacks()
  const { refetch } = useGetCurrentUser()
  const navigate = useNavigate()
  const { mutate, isPending } = useSignIn()

  const validateUsername = (value: string) => {
    if (value === '') {
      setUsernameError('ユーザー名を入力してください。')
      return false
    }
    const regex = /^[a-zA-Z0-9]+$/
    if (!regex.test(value)) {
      setUsernameError('ユーザー名は英数のみ許可されます。記号は使用できません。')
      return false
    } else {
      setUsernameError('')
      return true
    }
  }
  const validatePassword = (value: string) => {
    if (value === '') {
      setPasswordError('パスワードを入力してください。')
      return false
    }
    const regex = /^[\x20-\x7E]+$/
    if (!regex.test(value)) {
      setPasswordError('パスワードはASCII範囲の英数記号のみ許可されます。')
      return false
    } else {
      setPasswordError('')
      return true
    }
  }
  const handleUsernameChange = (event: { target: { value: string } }) => {
    const value = event.target.value
    setUsername(value)
    validateUsername(value)
  }
  const handlePasswordChange = (event: { target: { value: string } }) => {
    const value = event.target.value
    setPassword(value)
    validatePassword(value)
  }
  const handleSubmit = () => {
    const isUsernameValid = validateUsername(username)
    const isPasswordValid = validatePassword(password)
    if (isUsernameValid && isPasswordValid) {
      mutate(
        { data: { user_name: username, password: password } },
        {
          onSuccess: () => {
            refetch()
            navigate('/')
          },
          onError: () => {
            snack({
              title: 'エラー',
              description: 'ユーザーが登録されていません',
              variant: 'solid',
              status: 'error'
            })
          }
        }
      )
    }
  }

  return (
    <Container>
      {/* 画面中央に置きたい */}
      <Center h="80vh">
        <Box w="full" p="md" bg="White" borderRadius="md">
          {/* 間隔を空ける */}
          <Box h="10px" />
          <Center>
            <Text as="h1" fontWeight="bold" fontSize="32px" fontFamily="Kaushan Script">
              Cheer Topics
            </Text>
          </Center>
          <Box h="10px" />
          <Snacks snacks={snacks} gutter={[0, 'md']} />
          <FormControl isRequired isInvalid={usernameError !== ''} errorMessage={usernameError}>
            <Label fontWeight="bold" fontSize="16px" fontFamily="Inter">
              ユーザネーム
            </Label>
            <Input
              type="text"
              placeholder="ユーザー名を入力してください。"
              isRequired
              value={username}
              onChange={handleUsernameChange}
            />
          </FormControl>
          <Box h="10px" />
          <FormControl isRequired isInvalid={passwordError !== ''} errorMessage={passwordError}>
            <Label fontWeight="bold" fontSize="16px" fontFamily="Inter">
              パスワード
            </Label>
            <Input
              type="password"
              placeholder="パスワードを入力してください。"
              isRequired
              value={password}
              onChange={handlePasswordChange}
            />
          </FormControl>
          <Box h="30px" />
          <Center>
            <Button
              onClick={handleSubmit}
              color="White"
              bgGradient="linear(to-r, #00D1FF,#8ABBF5, #DEC9EB)"
              isLoading={isPending}
              loadingIcon="dots"
              fontWeight="bold"
              fontSize="16px"
              fontFamily="Inter"
            >
              ログイン
            </Button>
          </Center>
          <Box h="30px" />
          <Divider variant="solid" />
          <Box h="20px" />
          <Center>
            <Text as="a" href="/signup" fontSize="14px" textDecoration="underline" fontFamily="Inter" color="#656565">
              またはアカウントを新規作成
            </Text>
          </Center>
          <Box h="10px" />
        </Box>
      </Center>
    </Container>
  )
}
