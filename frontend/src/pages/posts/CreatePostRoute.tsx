import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Container, Divider, FormControl, HStack, Heading, Input,useSnacks } from '@yamada-ui/react'
import MarkdownEditor from '@uiw/react-markdown-editor'
import { Link } from 'react-router-dom'
import { usePostPosts } from '../../api/api'

export const CreatePostRoute = () => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const {snack} = useSnacks()

  const navigate = useNavigate()
  const {mutate} = usePostPosts()

  const validdateTitle = (value: string) => {
    if (value === '') {
      snack({
        status: 'error',
        description: 'タイトルを入力してください。'
      })
      return false
    }
    if (value.length > 100) {
      snack({
        status: 'error',
        description: 'タイトルは100文字以内で入力してください。'
      })
      return false
    }
    return true
  }

  const validateContent = (value: string) => {
    if (value === '') {
      snack({
        status: 'error',
        description: '内容を入力してください。'
      })
      return false
    }
    return true
  }

  const handleSubmit = (event: { preventDefault: () => void }) => {
    event.preventDefault()
    const isTitleValid = validdateTitle(title)
    const isContentValid = validateContent(content)

    if (isTitleValid && isContentValid) {
      mutate(
        {data:{title:title, body:content}},
        {
          onSuccess: () => {
            snack({
              status: 'success',
              description: '投稿に成功しました。'
            
            })
            navigate('/')
          },
          onError: () => {
            snack(
              {
                status: 'error',
                description: '投稿に失敗しました。'
              }
            )
          }
        }
      )
    console.log('submit', { title, content })
    }
  }

  return (
    <Container>
      <Heading size="lg">新しい投稿を作成する</Heading>
      <Divider variant="solid" />
      <form onSubmit={handleSubmit}>
        <FormControl label="タイトル" isRequired mb={4}>
          <Input
            type="text"
            placeholder="タイトルを入力してください。"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
        </FormControl>
        <FormControl label="内容" isRequired mb={4}>
          <MarkdownEditor value={content} height="200px" onChange={(value) => setContent(value)} />
        </FormControl>
        <HStack>
          <Link to="/">
            <Button colorScheme="primary" variant={'outline'}>
              キャンセル
            </Button>
          </Link>
          <Button onClick={handleSubmit} type="submit" colorScheme="primary">
            投稿する
          </Button>
        </HStack>
      </form>
    </Container>
  )
}
