import { useCreatePostComments, useLikePost } from '@api/hooks'
import { Box, Button, Flex, HStack, Icon, IconButton, Input, useSnacks } from '@yamada-ui/react'
import { MessageSquareText, Heart } from 'lucide-react'
import { useState } from 'react'

interface CreatePostCommentCardProps {
  id: number
  onSuccess: () => void
}

export const CreatePostCommentCard = ({ id, onSuccess }: CreatePostCommentCardProps) => {
  const [body, setBody] = useState('')
  const { mutate } = useCreatePostComments()
  const { snack } = useSnacks()
  const {mutate:mutateLike} = useLikePost()

  const handleBodyChange = (event: { target: { value: string } }) => {
    setBody(event.target.value)
  }

  const handleLike = () => {
    // いいねの処理
    mutateLike({postId: String(id)})
    console.log('いいね')

  }

  function handleSubmit(event: { preventDefault: () => void }): void {
    event.preventDefault()
    const trimmedBody = body.trim()
    if (trimmedBody === '') {
      snack({
        status: 'error',
        description: 'コメントを入力してください。'
      })
      return
    }
    mutate(
      { postId: String(id), data: { body: trimmedBody } },
      {
        onSuccess: () => {
          setBody('')
          snack({
            status: 'success',
            description: 'コメントの投稿に成功しました。'
          })
          onSuccess() // コメント投稿成功後に親コンポーネントで指定された関数を呼び出す
        },
        onError: () => {
          snack({
            status: 'error',
            description: 'コメントの投稿に失敗しました。'
          })
        }
      }
    )
  }

  return (
    <Flex px="md" py="lg" justifyContent="space-between">
      <Box w="75vw" p="sm" bgColor="whiteAlpha.800" borderRadius="md" _hover={{ bgColor: 'neutral.100' }}>
        <form onSubmit={handleSubmit}>
          <HStack>
            <Icon size="lg" as={MessageSquareText} />
            <Input type="text" placeholder="コメントを入力してください" value={body} onChange={handleBodyChange} />
            <Button
              color="White"
              bgGradient="linear(to-r, #00D1FF,#8ABBF5, #DEC9EB)"
              fontWeight="bold"
              fontSize="16px"
              fontFamily="Inter"
              w="100px"
              type="submit"
            >
              送信
            </Button>
          </HStack>
        </form>
      </Box>
      <IconButton colorScheme="whiteAlpha" variant="ghost" as={Heart} onClick={handleLike}/>
    </Flex>
  )
}
