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
  Center,
  Button,
  Flex,
  Spacer
} from '@yamada-ui/react'
import dayjs from 'dayjs'
import { useGetPosts } from '../../api/api'
import { Link } from 'react-router-dom'
import { useEffect } from 'react'
import { Markdown } from '@yamada-ui/markdown'

export const PostsRoute = () => {
  // API取得
  const { data, isLoading, isError, refetch } = useGetPosts()
  useEffect(() => {
    refetch()
  }, [])

  return (
    <Container>
      <Heading size="lg">
        <Flex gap="md">
          <Text> 投稿一覧</Text>
          <Spacer />
          <Link to="/posts/new">
            <Button colorScheme="primary">新規作成</Button>
          </Link>
        </Flex>
      </Heading>
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
            _hover={{
              cursor: 'pointer',
              bgColor: 'gray.50'
            }}
          >
            <Link to={`/posts/${post.id}`}>
              <CardHeader>
                <Heading size="md">{post.title}</Heading>
              </CardHeader>

              <CardBody>
                <Markdown>{post.body}</Markdown>
              </CardBody>
              <CardFooter>
                <HStack>
                  <Text>{post.user_name}</Text>
                  <Text>更新日時： {dayjs(post.updated_at).format('YYYY年M月D日 HH:mm:ss')}</Text>
                </HStack>
              </CardFooter>
            </Link>
          </Card>
        ))}
      </VStack>
    </Container>
  )
}
