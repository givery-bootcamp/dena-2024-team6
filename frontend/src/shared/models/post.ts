export type post = {
  id: string
  title: string
  body: string
  userId: string
  userName: string
  createdAt: string
  updatedAt: string
}

export const MOCK_POSTS: Array<post> = [
  {
    id: 'mock-id',
    title: 'モックタイトル',
    body: 'hoge',
    userId: 'mock-user-id1',
    userName: 'Funobu',
    createdAt: '2024-06-02T00:00:00+09:00',
    updatedAt: '2024-06-02T00:00:00+09:00'
  },
  {
    id: '1',
    title: 'モックタイトル1',
    body: 'こんにちは\nこんにちは',
    userId: 'mock-user-id1',
    userName: 'わだ',
    createdAt: '2024-06-02T00:00:00+09:00',
    updatedAt: '2024-06-02T00:00:00+09:00'
  },
  {
    id: '2',
    title: 'モックタイトル2',
    body: 'こんにちは\nこんにちは',
    userId: 'mock-user-id1',
    userName: 'わだ',
    createdAt: '2024-06-02T00:00:00+09:00',
    updatedAt: '2024-06-02T00:00:00+09:00'
  },
  {
    id: '3',
    title: 'モックタイトル3',
    body: 'こんにちは\nこんにちは',
    userId: 'mock-user-id1',
    userName: 'わだ',
    createdAt: '2024-06-02T00:00:00+09:00',
    updatedAt: '2024-06-02T00:00:00+09:00'
  },
  {
    id: '4',
    title: 'モックタイトル4',
    body: 'こんにちは\nこんにちは',
    userId: 'mock-user-id1',
    userName: 'わだ',
    createdAt: '2024-06-02T00:00:00+09:00',
    updatedAt: '2024-06-02T00:00:00+09:00'
  }
]
