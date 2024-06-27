import { Divider, Text } from '@yamada-ui/react'
import { AttributeDisplay } from './AttributeDisplay'

export const CommentRoute = ({ id }: { id: number }) => {
  // mock
  const commentList = [
    {
      id: 1,
      post_id: 1,
      user_id: 1,
      user_name: 'user1',
      body: 'body1',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    },
    {
      id: 2,
      post_id: 1,
      user_id: 2,
      user_name: 'user2',
      body: 'body2',
      created_at: '2021-01-01T00:00:00',
      updated_at: '2021-01-01T00:00:00'
    }
  ]

  return (
    <>
      {commentList.map((comment) => (
        <>
          <Divider variant="solid" />
          <Text>{comment.body}</Text>
          <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
            <AttributeDisplay labelName="ユーザー名：" value={comment.user_name} />
          </div>
        </>
      ))}
    </>
  )
}
