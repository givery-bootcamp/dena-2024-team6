import { Button, Divider, HStack, Input, useSnacks } from '@yamada-ui/react'
import { useState } from 'react'
import { usePostComments } from '../../../api/api'

export const CreatePostCommentRoute = ({ id }: { id: number }) => {
  const [body, setBody] = useState('')
  const [bodyError, setBodyError] = useState<string | null>(null)
  const { mutate } = usePostComments()

  const { snack } = useSnacks()

  const validdateBody = (value: string) => {
    if (value === '') {
      setBodyError('コメントを入力してください。')
      return false
    }
    setBodyError(null)
    return true
  }

  const handleBodyChange = (event: { target: { value: string } }) => {
    const value = event.target.value
    setBody(value)
    validdateBody(value)
  }
  function handleSubmit(event: { preventDefault: () => void }): void {
    event.preventDefault()
    const isBodyValid = validdateBody(body)
    if (isBodyValid) {
      mutate(
        { data: { post_id: String(id), body: body } },
        {
          onSuccess: () => {
            setBody('')
            snack({
              status: 'success',
              description: '投稿に成功しました。'
            })
          },
          onError: () => {
            snack({
              status: 'error',
              description: '投稿に失敗しました。'
            })
          }
        }
      )
    }
    return
  }

  return (
    <>
      <Divider variant="solid" />
      <form onSubmit={handleSubmit}>
        <HStack>
          <Input type="text" placeholder="コメントを入力してください" value={body} onChange={handleBodyChange} />
          <Button colorScheme="primary" type="submit">
            送信
          </Button>
        </HStack>
      </form>
    </>
  )
}
