import { useNavigate, useParams, Link } from 'react-router-dom'
import { useGetPostsPostId, useUpdatePost } from '../../../api/api'
import {
  Button,
  Center,
  Container,
  Divider,
  FormControl,
  HStack,
  Heading,
  Input,
  Loading,
  useSnacks
} from '@yamada-ui/react'
import MarkdownEditor from '@uiw/react-markdown-editor'
import { useState } from 'react'

export const UpdatePostRoute = () => {
  const { id } = useParams<{ id: string }>()
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const { data, isLoading, isError } = useGetPostsPostId(Number(id!))

  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [titleError, setTitleError] = useState('')
  const [contentError, setContentError] = useState('')

  const { snack } = useSnacks()

  const navigate = useNavigate()
  const { mutate } = useUpdatePost()

  const validdateTitle = (value: string) => {
    if (value === '') {
      setTitleError('タイトルを入力してください。')
      return false
    }
    if (value.length > 100) {
      setTitleError('タイトルは100文字以内で入力してください。')
      return false
    }
    setTitleError('')
    return true
  }

  const validateContent = (value: string) => {
    if (value === '') {
      setContentError('内容を入力してください。')
      return false
    }
    setContentError('')
    return true
  }

  const handleTitleChange = (event: { target: { value: string } }) => {
    const value = event.target.value
    setTitle(value)
    validdateTitle(value)
  }

  const handleContentChange = (value: string) => {
    setContent(value)
    validateContent(value)
  }

  const handleSubmit = (event: { preventDefault: () => void }) => {
    event.preventDefault()
    const isTitleValid = validdateTitle(title)
    const isContentValid = validateContent(content)

    if (isTitleValid && isContentValid) {
      mutate(
        { data: { id: Number(id!), title: title, body: content } },
        {
          onSuccess: () => {
            snack({
              status: 'success',
              description: '投稿の更新に成功しました。'
            })
            navigate('/')
          },
          onError: () => {
            snack({
              status: 'error',
              description: '投稿の更新に失敗しました。'
            })
          }
        }
      )
    }
  }

  return (
    <Container>
      <Heading size="lg">投稿を編集する</Heading>
      {isLoading && (
        <Center>
          <Loading variant="circles" size="6xl" color="cyan.500" />
        </Center>
      )}
      {isError && (
        <Center>
          <Heading>エラーが発生しました</Heading>
        </Center>
      )}
      <Divider variant="solid" />
      <form onSubmit={handleSubmit}>
        <FormControl label="タイトル" isRequired isInvalid={titleError !== ''} errorMessage={titleError} mb={4}>
          <Input
            type="text"
            placeholder="タイトルを入力してください。"
            defaultValue={data?.title}
            onChange={handleTitleChange}
          />
        </FormControl>
        <FormControl label="内容" isRequired isInvalid={contentError !== ''} errorMessage={contentError} mb={4}>
          <MarkdownEditor value={data?.body} height="200px" onChange={handleContentChange} />
        </FormControl>
        <HStack>
          <Link to={`/posts/${id}`}>
            <Button colorScheme="primary" variant={'outline'}>
              キャンセル
            </Button>
          </Link>
          <Button onClick={handleSubmit} type="submit" colorScheme="primary">
            変更を保存する
          </Button>
        </HStack>
      </form>
    </Container>
  )
}
