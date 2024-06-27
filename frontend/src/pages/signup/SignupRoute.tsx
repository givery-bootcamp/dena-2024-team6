import { useState } from 'react'
import { Container, FormControl, Input, Button, Snacks, useSnacks } from '@yamada-ui/react'

export const SignupRoute = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')

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
  }

  const handleSubmit = () => {
    const isUsernameValid = validateUsername(username)
    const isPasswordValid = validatePassword(password)
    if (isUsernameValid && isPasswordValid) {
      // アカウント作成処理をここに追加します
      // 例: register({ username, password })
      snack({
        title: '成功',
        description: 'アカウントが作成されました。',
        variant: 'solid',
        status: 'success'
      })
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
      <Button onClick={handleSubmit} colorScheme="primary">
        作成
      </Button>
    </Container>
  )
}
