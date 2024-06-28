import { useDisclosure, Flex, Box, Textarea, HStack } from '@yamada-ui/react'
import { Link, useParams } from 'react-router-dom'
import { useDeletePost, useGetPost, useGetCurrentUser } from '@api/hooks'
import { PostDetailCard } from './components/PostDetailCard'
import { CommentCard } from './components/CommentCard'

export const PostDetailRoute = () => {
  const { id } = useParams<{ id: string }>()
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const { data, isLoading, isError } = useGetPost(id!)
  const { data: user } = useGetCurrentUser()
  const { mutate } = useDeletePost()
  const { isOpen, onOpen, onClose } = useDisclosure()

  const handleDelete = () => {
    mutate({
      postid: id!
    })
  }

  const dummy = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]

  return (
    <Flex flexDir="column" gap="lg">
      <Box px="md" py="md">
        <PostDetailCard
          title={data?.title}
          body={data?.body}
          userName={'futa.matsuo'}
          createdAt={data?.created_at ? new Date(data.created_at) : undefined}
          isError={isError}
        />
      </Box>
      <Flex
        w="full"
        flexDir="column"
        gap="md"
        bgGradient="linear(to-b, transparent, blackAlpha.800)"
        position="absolute"
        bottom="0px"
      >
        <Flex flexDir="column" gap="md" px="md" py="sm" h="45vh" overflow="scroll">
          {dummy.map((d) => (
            <CommentCard
              key={d}
              userName="さくらみこ"
              body="JavaScriptは難しいにぇ..."
              createdAt={data?.created_at ? new Date(data.created_at) : undefined}
            />
          ))}
        </Flex>
        <HStack px="md" py="lg">
          <Box w="60vw" bgColor="whiteAlpha.900" borderRadius="md">
            <Textarea placeholder="コメントを入力..." />
          </Box>
        </HStack>
      </Flex>
    </Flex>
  )
}
