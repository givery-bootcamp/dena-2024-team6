import {
  Loading,
  Card,
  CardBody,
  CardHeader,
  Container,
  Text,
  HStack,
  Heading,
  CardFooter,
  VStack,
  Center
} from '@yamada-ui/react'
import dayjs from 'dayjs'
import { useGetPosts } from '../../api/api'
import { useNavigate } from 'react-router-dom'

export const PostsRoute = () => {
  // API取得
  const { data, isLoading, isError } = useGetPosts()
  const navigate = useNavigate()

  return (
    <Container>
      <Heading size="lg">投稿一覧</Heading>
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
      <VStack w="full">
        {data?.map((post) => (
          <Card
            key={post.id}
            variant="outline"
            w="full"
            onClick={() => navigate(`/posts/${post.id}`)}
            _hover={{
              cursor: 'pointer',
              bgColor: 'gray.50'
            }}
          >
            <CardHeader>
              <Heading size="md">{post.title}</Heading>
            </CardHeader>

            <CardBody>
              <Text>{post.body}</Text>
            </CardBody>
            <CardFooter>
              <HStack>
                <Text>{post.user_name}</Text>
                <Text>更新日時： {dayjs(post.updated_at).format('YYYY年M月D日 HH:mm:ss')}</Text>
              </HStack>
            </CardFooter>
          </Card>
        ))}
      </VStack>
    </Container>
  )
}
