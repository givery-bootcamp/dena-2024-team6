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
  Image
} from '@yamada-ui/react'
import { memo } from 'react'
import dayjs from 'dayjs'

type PostDetailCardProps = {
  title?: string
  body?: string
  userName?: string
  userIconURL?: string
  createdAt?: Date
  isError?: boolean
}

export const PostDetailCard = memo(function ({
  title,
  body,
  userName,
  userIconURL,
  createdAt,
  isError = false
}: PostDetailCardProps) {
  return (
    <Box w="full" bgColor="#ffffff" borderRadius="lg" p="md">
      {!isError ? (
        <Flex gap="md">
          {userName ? <Avatar size="sm" name={userName} src={userIconURL} /> : <SkeletonCircle w="32px" h="28px" />}
          <VStack>
            {userName ? (
              <HStack>
                <Text color="black">Funobu</Text>
                <Text fontSize="sm" color="neutral.300">
                  {dayjs(createdAt).format('YYYY年M月D日')}
                </Text>
              </HStack>
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