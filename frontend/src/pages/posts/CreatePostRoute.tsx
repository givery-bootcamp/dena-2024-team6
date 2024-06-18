import { useState } from 'react'
import { Button, Container, Divider, FormControl, HStack, Heading, Input } from '@yamada-ui/react'
import MarkdownEditor from '@uiw/react-markdown-editor'
import { Link } from 'react-router-dom'

export const CreatePostRoute = () => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const handleSubmit = (event: { preventDefault: () => void }) => {
    event.preventDefault()
    console.log('submit', { title, content })
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
          <Button type="submit" colorScheme="primary">
            投稿する
          </Button>
        </HStack>
      </form>
    </Container>
  )
}
