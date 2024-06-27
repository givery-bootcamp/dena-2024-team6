import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Container, FormControl, Input, Snacks, useSnacks, Text } from '@yamada-ui/react'
import { usePostSignin, useGetUser } from '../../api/api'
export const SigninRoute = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const { snack, snacks } = useSnacks()
  const { refetch } = useGetUser()
  const navigate = useNavigate()
  const { mutate, isPending } = usePostSignin()
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
      <Snacks snacks={snacks} gutter={[0, 'md']} />
      <FormControl label="ユーザー名" isRequired isInvalid={usernameError !== ''} errorMessage={usernameError}>
        <Input
          type="text"
          placeholder="ユーザー名を入力してください。"
          isRequired
          value={username}
          onChange={handleUsernameChange}
        />
      </FormControl>
      <FormControl label="パスワード" isRequired isInvalid={passwordError !== ''} errorMessage={passwordError}>
        <Input
          type="password"
          placeholder="パスワードを入力してください。"
          isRequired
          value={password}
          onChange={handlePasswordChange}
        />
      </FormControl>
      <Button onClick={handleSubmit} isLoading={isPending} loadingIcon="dots" colorScheme="primary">
        サインイン
      </Button>
      <Text
        as="a"
        mt={4}
        onClick={() => {
          navigate('/signup')
        }}
        cursor="pointer"
        textDecoration="underline"
      >
        または、アカウントを新規作成
      </Text>
    </Container>
  )
}
