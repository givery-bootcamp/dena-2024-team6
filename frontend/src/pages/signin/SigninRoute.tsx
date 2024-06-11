import { useState } from 'react'
import { Button, Container, FormControl, Input } from '@yamada-ui/react'

export const SigninRoute = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')

  const validateUsername = (value: string) => {
    if (value === '') {
      setUsernameError('ユーザー名を入力してください。')
      return
    }
    const regex = /^[a-zA-Z0-9]+$/
    if (!regex.test(value)) {
      setUsernameError('ユーザー名は英数のみ許可されます。記号は使用できません。')
    } else {
      setUsernameError('')
    }
  }

  const validatePassword = (value: string) => {
    if (value === '') {
      setPasswordError('パスワードを入力してください。')
      return
    }
    const regex = /^[\x20-\x7E]+$/ // ASCII範囲の印刷可能な文字
    if (!regex.test(value)) {
      setPasswordError('パスワードはASCII範囲の英数記号のみ許可されます。')
    } else {
      setPasswordError('')
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

  return (
    <Container>
      <FormControl label="ユーザー名" isRequired isInvalid={usernameError !== ''} errorMessage={usernameError}>
        <Input
          type="text"
          placeholder="ユーザー名を入力してください。"
          value={username}
          onChange={handleUsernameChange}
        />
      </FormControl>

      <FormControl label="パスワード" isRequired isInvalid={passwordError !== ''} errorMessage={passwordError}>
        <Input
          type="password"
          placeholder="パスワードを入力してください。"
          value={password}
          onChange={handlePasswordChange}
        />
      </FormControl>

      <Button
        onClick={() => {
          // TODO: エラーがない場合にサインインAPIを叩き、ルーティング
        }}
      >
        サインイン
      </Button>
    </Container>
  )
}
