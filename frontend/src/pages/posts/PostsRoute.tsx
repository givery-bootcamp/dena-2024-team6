import {
  Loading,
  Card,
  CardBody,
  Container,
  Text,
  HStack,
  Heading,
  VStack,
  Center,
  Button,
  Flex,
  Spacer,
  Box,
  Divider,
  useDisclosure
} from '@yamada-ui/react'
import { useListPosts } from '@api/hooks'
import { Link } from 'react-router-dom'
import { useEffect } from 'react'
import { Beer, FileText } from 'lucide-react'
import { CreatePostModal } from './CreatePostRoute'

export const PostsRoute = () => {
  // API取得
  const { data, isLoading, isError, refetch } = useListPosts()
  const { isOpen, onOpen, onClose } = useDisclosure()

  useEffect(() => {
    refetch()
  }, [])

  return (
    <Container>
      <Box h="30px" />
      <Heading size="lg">
        <Center>
          <Text fontWeight="extrabold" fontSize="18px" fontFamily="Inter">
            LET'S RELEASE YOUR PASSION
          </Text>
        </Center>
        <Box h="10px" />
        <Center>
          <Text fontWeight="extrabold" fontSize="14px" fontFamily="Inter" color="#E64545">
            {data?.length}&nbsp;
          </Text>
          <Text fontSize="14px" fontFamily="Inter">
            tables are exist.
          </Text>
        </Center>
        <Box h="10px" />
      </Heading>
      {isLoading && (
        <Center>
          <Loading variant="circles" size="6xl" color="White" />
        </Center>
      )}
      {isError && (
        <Center>
          <Heading>エラーが発生しました</Heading>
        </Center>
      )}
      <Divider variant="solid" />
      <VStack h="500px" p="md" w="full" overflow="auto">
        {data?.map((post) => (
          <Card
            key={post.post_id}
            bg="White"
            w="full"
            _hover={{
              cursor: 'pointer',
              bgColor: 'gray.50'
            }}
          >
            <Link to={`/posts/${post.post_id}`}>
              <CardBody>
                <HStack>
                  <Beer />
                  <Text fontWeight="bold" fontSize="16px" fontFamily="Inter" color="#E64545">
                    0
                  </Text>
                  <Spacer />
                  <Text fontWeight="bold" fontSize="14px" fontFamily="Inter">
                    {post.title}
                  </Text>
                </HStack>
              </CardBody>
            </Link>
          </Card>
        ))}
      </VStack>
      <Divider variant="solid" />
      <Flex justifyContent="flex-end">
        <Button
          bg="White"
          borderWidth="4px"
          borderColor="#F092FF"
          borderRadius="full"
          variant="outline"
          h="80px"
          w="80px"
          onClick={onOpen}
        >
          <Center>
            <VStack gap="0">
              <Box h="10px" />
              <Center>
                <FileText color="#EA67D5" size="30" />
              </Center>
              <Text fontFamily="Inter" fontWeight="bold" fontSize="12px" color="#583474">
                ポスト
              </Text>
            </VStack>
          </Center>
        </Button>
      </Flex>
      <CreatePostModal isOpen={isOpen} onClose={onClose} fetchPosts={refetch} />
    </Container>
  )
}
