import { useDisclosure, Flex, Box, Dialog } from '@yamada-ui/react'
import { useNavigate, useParams } from 'react-router-dom'
import { useDeletePost, useGetPost, useGetCurrentUser, useListPostComments } from '@api/hooks'
import { PostDetailCard } from './components/PostDetailCard'
import { CommentCard } from './components/CommentCard'
import { CreatePostCommentCard } from './components/createCommentCard'
import { useEffect, useRef } from 'react'
export const PostDetailRoute = () => {
  const endOfCommentRef = useRef<HTMLDivElement>(null)
  const navigate = useNavigate()
  const { isOpen, onOpen, onClose } = useDisclosure()
  const { id } = useParams<{ id: string }>()
  const { data, isError } = useGetPost(id!)
  const { data: commentList } = useListPostComments(id!, {
    query: {
      refetchInterval: 500
    }
  })
  const { data: user } = useGetCurrentUser()
  const { mutate } = useDeletePost()

  const handleDelete = () => {
    mutate(
      {
        postid: id!
      },
      {
        onSuccess: () => {
          navigate('/')
        }
      }
    )
  }

  useEffect(() => {
    endOfCommentRef?.current?.scrollIntoView({ behavior: 'smooth' })
  }, [commentList])

  return (
    <Flex w="full" flexDir="column" gap="lg">
      <Box px="md" py="md">
        <PostDetailCard
          title={data?.title}
          body={data?.body}
          userName={data?.user_name}
          createdAt={data?.created_at ? new Date(data.created_at) : undefined}
          isAuthor={data?.user_id == user?.user_id}
          isError={isError}
          onEdit={() => {
            navigate(`/posts/${id}/edit`)
          }}
          onDelete={() => {
            onOpen()
          }}
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
        <Flex
          flexDir="column"
          gap="md"
          px="md"
          py="sm"
          h="45vh"
          overflowY="scroll"
          _scrollbar={{
            w: '8px',
            height: '32px'
          }}
          _scrollbarThumb={{
            bgColor: 'black'
          }}
        >
          {commentList?.map((c) => (
            <CommentCard
              key={c.id}
              userName={c.user_name}
              body={c.body}
              createdAt={c.created_at ? new Date(c.created_at) : undefined}
            />
          ))}
          <Box h="1px" ref={endOfCommentRef} />
        </Flex>
        <CreatePostCommentCard id={Number(id)} onSuccess={() => {}} />
      </Flex>
      <Dialog
        header={data?.title + 'の削除'}
        isOpen={isOpen}
        onClose={onClose}
        cancel="キャンセル"
        onCancel={onClose}
        success={{
          colorScheme: 'danger',
          children: '削除する'
        }}
        onSuccess={() => {
          handleDelete()
        }}
      >
        投稿を削除しますか？削除すると関連する情報やコメントや閲覧できなくなります
      </Dialog>
    </Flex>
  )
}
