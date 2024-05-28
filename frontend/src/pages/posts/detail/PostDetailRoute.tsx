import { Container, Text, HStack, Heading, Divider } from '@yamada-ui/react'
import { useState } from 'react'
import { MOCK_POSTS, post } from '../../../shared/models'
import dayjs from 'dayjs'
import { AttributeDisplay } from './AttributeDisplay'
import { useParams } from 'react-router-dom'

export const PostDetailRoute = () => {
  const { id } = useParams<{ id: string }>()
  const [post, setPost] = useState<post>(() => {
    const foundPost = MOCK_POSTS.find((post) => post.id === id)
    return foundPost || MOCK_POSTS[0]
  })

  return (
    <Container>
      <Heading size="lg">{post.title}</Heading>
      <HStack>
        <AttributeDisplay labelName="作成日時：" value={dayjs(post.createdAt).format('YYYY年M月D日 HH:mm:ss')} />
        <AttributeDisplay labelName="更新日時：" value={dayjs(post.updatedAt).format('YYYY年M月D日 HH:mm:ss')} />
        <AttributeDisplay labelName="ユーザー名：" value={post.userName} />
      </HStack>
      <Divider variant="solid" />
      <Text>{post.body}</Text>
    </Container>
  )
}
