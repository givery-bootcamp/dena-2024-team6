import { Markdown } from '@yamada-ui/markdown'
import {
  Text,
  Avatar,
  Box,
  HStack,
  Heading,
  VStack,
  Flex,
  Divider,
  SkeletonCircle,
  Skeleton,
  SkeletonText,
  Center,
  Image,
  IconButton
} from '@yamada-ui/react'
import { memo } from 'react'
import dayjs from 'dayjs'
import { Edit2, Trash2 } from 'lucide-react'

type PostDetailCardProps = {
  title?: string
  body?: string
  userName?: string
  userIconURL?: string
  createdAt?: Date
  isAuthor?: boolean
  isError?: boolean
  onEdit: () => void
  onDelete: () => void
}

export const PostDetailCard = memo(function ({
  title,
  body,
  userName,
  userIconURL,
  createdAt,
  isAuthor = false,
  isError = false,
  onEdit,
  onDelete
}: PostDetailCardProps) {
  return (
    <Box w="full" bgColor="#ffffff" borderRadius="lg" p="md">
      {!isError ? (
        <Flex gap="md">
          {userName ? (
            <Avatar color="White" bg="#583474" size="sm" name={userName} src={userIconURL} />
          ) : (
            <SkeletonCircle w="32px" h="28px" />
          )}
          <VStack>
            {userName ? (
              <Flex justifyContent="space-between">
                <HStack>
                  <Text color="black">{userName}</Text>
                  <Text fontSize="sm" color="neutral.300">
                    {dayjs(createdAt).format('YYYY年M月D日')}
                  </Text>
                </HStack>
                {isAuthor ? (
                  <HStack>
                    <IconButton size="xs" as={Edit2} variant="ghost" onClick={onEdit} />
                    <IconButton size="xs" as={Trash2} colorScheme="danger" variant="ghost" onClick={onDelete} />
                  </HStack>
                ) : null}
              </Flex>
            ) : (
              <Skeleton w="40%" />
            )}
            {title ? (
              <Flex flexDir="column" gap="sm">
                <Heading fontSize="lg">{title}</Heading>
                <Divider bgColor="black" />
              </Flex>
            ) : (
              <Flex flexDir="column" gap="sm">
                <Skeleton />
                <Divider bgColor="black" />
              </Flex>
            )}
            {body ? (
              <Box overflow="scroll" maxH={{ base: '20vh' }}>
                <Markdown components={markdownComponents}>{body}</Markdown>
              </Box>
            ) : (
              <Box overflow="scroll">
                <SkeletonText lineClamp={6} />
              </Box>
            )}
          </VStack>
        </Flex>
      ) : (
        <Center>
          <Heading>エラーが発生しました</Heading>
        </Center>
      )}
    </Box>
  )
})

const markdownComponents = {
  img: (props: any) => <Image display="block" w={{ xs: 'full', lg: '20vw' }} maxH="full" {...props} />
}
