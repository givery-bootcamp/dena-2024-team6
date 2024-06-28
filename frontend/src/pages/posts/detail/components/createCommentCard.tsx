import { useCreatePostComments } from '@api/hooks'
import { Text, Box, Button, Divider, Flex, HStack, Icon, IconButton, Input, useSnacks } from '@yamada-ui/react'
import { MessageSquareText, Heart } from 'lucide-react'
import { useState } from 'react'

export const CreatePostCommentCard = ({ id }: { id: number }) => {
  const [body, setBody] = useState('')
  const [bodyError, setBodyError] = useState<string | null>(null)
  const { mutate } = useCreatePostComments()

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
        { postId: String(id), data: { body: body } },
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
    <Flex px="md" py="lg" justifyContent="space-between">
      <Box
        w="75vw"
        p="sm"
        bgColor="whiteAlpha.800"
        borderRadius="md"
        _hover={{
          bgColor: 'neutral.100'
        }}
      >
        <form onSubmit={handleSubmit}>
          <HStack>
            <Icon size="lg" as={MessageSquareText} />
            <Input type="text" placeholder="コメントを入力してください" value={body} onChange={handleBodyChange} />
            <Button
              onClick={handleSubmit}
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
      <IconButton colorScheme="whiteAlpha" variant="ghost" as={Heart} />
    </Flex>
  )
}
