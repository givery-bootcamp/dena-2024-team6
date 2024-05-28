import { Card, CardBody, CardHeader, Container, Text, HStack, Heading, CardFooter, VStack } from '@yamada-ui/react'
import { useEffect, useState } from 'react'
import { MOCK_POSTS, post } from '../../shared/models'
import dayjs, { Dayjs } from 'dayjs'

export const PostsRoute = () => {
  // API取得
  const [posts, setPosts] = useState<Array<post>>(MOCK_POSTS)
  return <Container></Container>
}
