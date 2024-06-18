import { Routes, Route } from 'react-router-dom'

import { PostsRoute } from '../pages/posts/PostsRoute'
import { SigninRoute } from '../pages/signin/SigninRoute'
import { PostDetailRoute } from '../pages/posts/detail/PostDetailRoute'

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/signin" element={<SigninRoute />} />
      <Route path="/" element={<PostsRoute />} />
      <Route path="/posts/:id" element={<PostDetailRoute />} />
    </Routes>
  )
}
