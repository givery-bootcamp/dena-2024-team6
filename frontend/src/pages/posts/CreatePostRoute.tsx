import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  Button,
  Divider,
  FormControl,
  Input,
  useSnacks,
  Modal,
  ModalOverlay,
  ModalHeader,
  ModalBody,
  ModalFooter
} from '@yamada-ui/react'
import MarkdownEditor from '@uiw/react-markdown-editor'
import { useCreatePost } from '@api/hooks'

type props = {
  isOpen: boolean
  onClose: () => void
  fetchPosts: () => void
}

export const CreatePostModal = ({ isOpen, onClose, fetchPosts }: props) => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [titleError, setTitleError] = useState<string | null>(null)
  const [contentError, setContentError] = useState<string | null>(null)
  const { snack } = useSnacks()
  const navigate = useNavigate()
  const { mutate } = useCreatePost()
  const validateTitle = (value: string) => {
    if (value === '') {
      setTitleError('タイトルを入力してください。')
      return false
    }
    if (value.length > 100) {
      setTitleError('タイトルは100文字以内で入力してください。')
      return false
    }
    setTitleError(null)
    return true
  }
  const validateContent = (value: string) => {
    if (value === '') {
      setContentError('内容を入力してください。')
      return false
    }
    setContentError(null)
    return true
  }
  const handleTitleChange = (event: { target: { value: string } }) => {
    const value = event.target.value
    setTitle(value)
    validateTitle(value)
  }
  const handleContentChange = (value: string) => {
    setContent(value)
    validateContent(value)
  }
  const handleSubmit = (event: { preventDefault: () => void }) => {
    event.preventDefault()
    const isTitleValid = validateTitle(title)
    const isContentValid = validateContent(content)
    if (isTitleValid && isContentValid) {
      mutate(
        { data: { title, body: content } },
        {
          onSuccess: () => {
            onClose()
            navigate('/')
            fetchPosts()
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
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalHeader>投稿の新規作成</ModalHeader>
      <ModalBody>
        <Divider variant="solid" />
        <form onSubmit={handleSubmit}>
          <FormControl label="タイトル" isRequired isInvalid={!!titleError} errorMessage={titleError} mb={4}>
            <Input type="text" placeholder="タイトルを入力してください。" value={title} onChange={handleTitleChange} />
          </FormControl>
          <FormControl label="本文" isRequired isInvalid={!!contentError} errorMessage={contentError} mb={4}>
            <MarkdownEditor value={content} height="200px" onChange={handleContentChange} />
          </FormControl>
        </form>
      </ModalBody>
      <ModalFooter>
        <Button variant="outline" onClick={onClose}>
          キャンセル
        </Button>
        <Button
          onClick={handleSubmit}
          color="White"
          bgGradient="linear(to-r, #00D1FF,#8ABBF5, #DEC9EB)"
          loadingIcon="dots"
          fontWeight="bold"
          fontSize="16px"
          fontFamily="Inter"
        >
          投稿を作成する
        </Button>
      </ModalFooter>
    </Modal>
  )
}
