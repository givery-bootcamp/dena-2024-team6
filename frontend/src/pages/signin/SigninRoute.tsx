import { Card, CardBody, CardHeader, Container, Text, HStack, Heading, CardFooter, VStack } from '@yamada-ui/react'
import { useEffect, useState } from 'react'
import { MOCK_POSTS, post } from '../../shared/models'
import dayjs from 'dayjs'

export const SigninRoute = () => {
  const [posts, setPosts] = useState<Array<post>>(MOCK_POSTS)
  return (
    <Container>
      <Heading size="lg">投稿一覧</Heading>
      <VStack w="full">
        {posts.map((post) => (
          <Card key={post.id} variant="outline" w="full">
            <CardHeader>
              <Heading size="md">{post.title}</Heading>
            </CardHeader>

            <CardBody>
              <Text>{post.body}</Text>
            </CardBody>
            <CardFooter>
              <HStack>
                <Text>{post.userName}</Text>
                <Text>更新日時： {dayjs(post.updatedAt).format('YYYY年M月D日 HH:mm:ss')}</Text>
              </HStack>
            </CardFooter>
          </Card>
        ))}
      </VStack>
    </Container>
  )
}
