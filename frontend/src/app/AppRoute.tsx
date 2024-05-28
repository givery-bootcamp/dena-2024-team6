import { Routes, Route } from 'react-router-dom'

import { PostsRoute } from '../pages/posts/PostsRoute'
import { PostDetailRoute } from '../pages/posts/detail/PostDetailRoute'

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<PostsRoute />} />
      <Route path="/posts/:id" element={<PostDetailRoute />} />
    </Routes>
  )
}
