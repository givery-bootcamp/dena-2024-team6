import { Container, Text, HStack, Heading } from '@yamada-ui/react'
import { useState } from 'react'
import { MOCK_POSTS, post } from '../../../shared/models'
import dayjs from 'dayjs'
import { AttributeDisplay } from './AttributeDisplay'

export const PostDetailRoute = () => {
  // API取得
  const [post, setPost] = useState<post>(MOCK_POSTS[0])
  return (
    <Container>
      <Heading size="lg">{post.title}</Heading>
      <HStack>
        <AttributeDisplay labelName="作成日時：" value={dayjs(post.createdAt).format('YYYY年M月D日 HH:mm:ss')} />
        <AttributeDisplay labelName="更新日時：" value={dayjs(post.updatedAt).format('YYYY年M月D日 HH:mm:ss')} />
        <AttributeDisplay labelName="ユーザー名：" value={post.userName} />
      </HStack>
      <Text>{post.body}</Text>
    </Container>
  )
}
