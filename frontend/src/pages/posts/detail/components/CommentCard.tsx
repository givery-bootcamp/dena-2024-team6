import { Text, Avatar, Box, Flex, HStack, SkeletonCircle } from '@yamada-ui/react'
import { memo } from 'react'
import dayjs from 'dayjs'

type CommentCardProps = {
  userName?: string
  userIconURL?: string
  body?: string
  createdAt?: Date
}

export const CommentCard = memo(function ({ userName, userIconURL, body, createdAt }: CommentCardProps) {
  return (
    <Flex w="full" borderRadius="lg" gap="md">
      {userName ? <Avatar  color="White" bg="#583474" size="sm" name={userName} src={userIconURL} /> : <SkeletonCircle w="32px" h="28px" />}
      <Flex gap="sm" flexDir="column">
        <HStack gap="lg">
          <Text color="#ffffff">{userName}</Text>
          <Text fontSize="sm" color="neutral.50">
            {dayjs(createdAt).format('YYYY年M月D日 hh:mm s秒')}
          </Text>
        </HStack>
        <Text color="#ffffff" fontSize="md">
          {body}
        </Text>
      </Flex>
    </Flex>
  )
})
