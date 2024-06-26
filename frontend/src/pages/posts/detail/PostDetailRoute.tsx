import { Container, Text, HStack, Heading, Divider, Center, Loading } from '@yamada-ui/react'
import dayjs from 'dayjs'
import { AttributeDisplay } from './AttributeDisplay'
import { useParams, Link } from 'react-router-dom'
import { useGetPostsPostId } from '../../../api/api'
import { Markdown } from '@yamada-ui/markdown';

export const PostDetailRoute = () => {
  const { id } = useParams<{ id: string }>()
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const { data, isLoading, isError } = useGetPostsPostId(Number(id!))

  return (
    <Container>
      <Heading size="lg">{data?.title}</Heading>
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
      {data && (
        <HStack>
          <AttributeDisplay labelName="作成日時：" value={dayjs(data?.created_at).format('YYYY年M月D日 HH:mm:ss')} />
          <AttributeDisplay labelName="更新日時：" value={dayjs(data?.updated_at).format('YYYY年M月D日 HH:mm:ss')} />
          <AttributeDisplay labelName="ユーザー名：" value={data?.user_name} />
        </HStack>
      )}
      <Divider variant="solid" />
      <Text>
      <Markdown>{data?.body}</Markdown>
    </Text>
      <Link to="/">戻る</Link>
    </Container>
  )
}
