import { Box, HStack, Text, VStack } from '@yamada-ui/react'
import { AttributeDisplay } from './AttributeDisplay'
import dayjs from 'dayjs'
// import { useGetComments } from '../../../api/api'

export const CommentRoute = ({ id }: { id: number }) => {
  // mock
  const mockcommentList = [
    {
      id: 1,
      post_id: 1,
      user_id: 1,
      user_name: 'user1',
      body: 'body1',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    },
    {
      id: 2,
      post_id: 1,
      user_id: 2,
      user_name: 'user2',
      body: 'body2',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    },
    {
      id: 2,
      post_id: 1,
      user_id: 2,
      user_name: 'user2',
      body: 'body2',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    },
    {
      id: 2,
      post_id: 1,
      user_id: 2,
      user_name: 'user2',
      body: 'body2',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    },
    {
      id: 2,
      post_id: 1,
      user_id: 2,
      user_name: 'user2',
      body: 'body2',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    }
  ]
  // const { data: commentList } = useGetComments(id)
  const comments =mockcommentList

  return (
    <>
      <VStack h="md" p="md" w="full" overflow="auto">
        {comments.map((comment) => (
          <>
            <Box w="full" p="md" border="2px solid" borderColor="gray.50" borderRadius="md">
              <Text>{comment.body}</Text>
              <HStack>
                <Text>{comment.user_name}</Text>
                <AttributeDisplay
                  labelName="作成日時："
                  value={dayjs(comment.created_at).format('YYYY年M月D日 HH:mm:ss')}
                />
              </HStack>
            </Box>
          </>
        ))}
      </VStack>
    </>
  )
}
