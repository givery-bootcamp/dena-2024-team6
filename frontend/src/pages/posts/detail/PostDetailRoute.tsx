import {
  Container,
  Text,
  HStack,
  Heading,
  Divider,
  Center,
  Loading,
  Button,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  useDisclosure
} from '@yamada-ui/react'
import dayjs from 'dayjs'
import { AttributeDisplay } from './AttributeDisplay'
import { Link, useParams } from 'react-router-dom'
import { useDeletePost, useGetPostsPostId, useGetUser } from '../../../api/api'
import { Markdown } from '@yamada-ui/markdown'
import { CommentRoute } from './PostCommentsRoute'
import { CreatePostRoute } from '../CreatePostRoute'
import { CreatePostCommentRoute } from './CreatePostCommentRoute'

export const PostDetailRoute = () => {
  const { id } = useParams<{ id: string }>()
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const { data, isLoading, isError } = useGetPostsPostId(Number(id!))
  const { data: user } = useGetUser()
  const { mutate } = useDeletePost()
  const { isOpen, onOpen, onClose } = useDisclosure()

  const handleDelete = () => {
    mutate(Number(id!))
  }

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

      <HStack>
        {data?.user_id === user?.id ? (
          <>
            <Link to={`/posts/${id}/edit`}>
              <Button onClick={() => console.log('edit')} colorScheme="primary">
                編集
              </Button>
            </Link>
            <Button onClick={onOpen} colorScheme="primary">
              削除
            </Button>
            <>
              <Modal isOpen={isOpen} onClose={onClose}>
                <Center>
                  <ModalHeader>警告</ModalHeader>
                </Center>
                <Divider variant="solid" />
                <ModalBody>削除したら元に戻せません。削除しますか？</ModalBody>
                <Divider variant="solid" />
                <ModalFooter>
                  <Button variant="ghost" onClick={onClose}>
                    とじる
                  </Button>
                  <>
                    <Link to="/">
                      <Button onClick={handleDelete} colorScheme="primary">
                        削除
                      </Button>
                    </Link>
                  </>
                </ModalFooter>
              </Modal>
            </>
          </>
        ) : null}
      </HStack>
      <CommentRoute id={Number(id!)} />
      <Link to="/">
        <Button colorScheme="primary" variant={'outline'}>
          戻る
        </Button>
      </Link>
      <CreatePostCommentRoute id={Number(id!)} />
    </Container>
  )
}
