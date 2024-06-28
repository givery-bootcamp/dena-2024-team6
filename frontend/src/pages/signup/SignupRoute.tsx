import { useState } from 'react'
import {
  Container,
  FormControl,
  Input,
  Button,
  Snacks,
  useSnacks,
  Box,
  Center,
  Divider,
  Label,
  Text
} from '@yamada-ui/react'
export const SignupRoute = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [confirmPasswordError, setConfirmPasswordError] = useState('')
  const { snack, snacks } = useSnacks()
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
  const handleUsernameChange = (event: { target: { value: any } }) => {
    const value = event.target.value
    setUsername(value)
    validateUsername(value)
  }
  const handlePasswordChange = (event: { target: { value: any } }) => {
    const value = event.target.value
    setPassword(value)
    validatePassword(value)
    if (confirmPassword !== '' && value !== confirmPassword) {
      setConfirmPasswordError('パスワードが一致しません')
    } else {
      setConfirmPasswordError('')
    }
  }
  const handleConfirmPasswordChange = (event: { target: { value: any } }) => {
    const value = event.target.value
    setConfirmPassword(value)
    if (password !== value) {
      setConfirmPasswordError('パスワードが一致しません')
    } else {
      setConfirmPasswordError('')
    }
  }
  const handleSubmit = () => {
    const isUsernameValid = validateUsername(username)
    const isPasswordValid = validatePassword(password)
    const isConfirmPasswordValid = confirmPassword === password
    if (isUsernameValid && isPasswordValid && isConfirmPasswordValid) {
      // アカウント作成処理をここに追加します
      // 例: register({ username, password })
      snack({
        title: '成功',
        description: 'アカウントが作成されました。',
        variant: 'solid',
        status: 'success'
      })
    } else if (!isConfirmPasswordValid) {
      setConfirmPasswordError('パスワードが一致しません')
    }
  }
  return (
    <Container>
      <Center h="80vh">
        <Box w="full" p="md" bg="White" borderRadius="md">
          <Box h="10px" />
          <Center>
            <Text as="h1" fontWeight="bold" fontSize="32px" fontFamily="Kaushan Script">
              Cheer Topics
            </Text>
          </Center>
          <Box h="10px" />
          <Snacks snacks={snacks} gutter={[0, 'md']} />
          <FormControl isRequired isInvalid={usernameError !== ''} errorMessage={usernameError}>
            <Label fontWeight="bold" fontSize="14px" fontFamily="Inter">
              ユーザネーム
            </Label>
            <Input
              type="text"
              placeholder="ユーザー名を入力してください。"
              fontSize="14px"
              isRequired
              value={username}
              onChange={handleUsernameChange}
            />
          </FormControl>
          <Box h="14px" />
          <FormControl isRequired isInvalid={passwordError !== ''} errorMessage={passwordError}>
            <Label fontWeight="bold" fontSize="14px" fontFamily="Inter">
              パスワード
            </Label>
            <Input
              type="password"
              placeholder="パスワードを入力してください。"
              fontSize="14px"
              isRequired
              value={password}
              onChange={handlePasswordChange}
            />
          </FormControl>
          <Box h="14px" />
          <FormControl isRequired isInvalid={confirmPasswordError !== ''} errorMessage={confirmPasswordError}>
            <Label fontWeight="bold" fontSize="14px" fontFamily="Inter">
              パスワード(確認)
            </Label>
            <Input
              type="password"
              placeholder="もう一度入力してください。"
              fontSize="14px"
              isRequired
              value={confirmPassword}
              onChange={handleConfirmPasswordChange}
            />
          </FormControl>
          <Box h="30px" />
          <Center>
            <Button
              onClick={handleSubmit}
              color="White"
              bgGradient="linear(to-r, #00D1FF,#8ABBF5, #DEC9EB)"
              loadingIcon="dots"
              fontWeight="bold"
              fontSize="14px"
              fontFamily="Inter"
            >
              アカウント登録
            </Button>
          </Center>
          <Box h="30px" />
          <Divider variant="solid" />
          <Box h="20px" />
          <Center>
            <Text as="a" href="/signin" fontSize="14px" textDecoration="underline" fontFamily="Inter" color="#656565">
              既にアカウントがある場合はログイン
            </Text>
          </Center>
          <Box h="10px" />
        </Box>
      </Center>
    </Container>
  )
}
