import { Routes, Route } from 'react-router-dom'

import { PostsRoute } from '../pages/posts/PostsRoute'
import { SigninRoute } from '../pages/signin/SigninRoute'
import { PostDetailRoute } from '../pages/posts/detail/PostDetailRoute'
import { CreatePostRoute } from '../pages/posts/CreatePostRoute'
import { UpdatePostRoute } from '../pages/posts/update/UpdatePostRoute'

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/signin" element={<SigninRoute />} />
      <Route path="/" element={<PostsRoute />} />
      <Route path="/posts/:id" element={<PostDetailRoute />} />
      <Route path="/posts/:id/edit" element={<UpdatePostRoute />} />
      <Route path="/posts/new" element={<CreatePostRoute />} />
    </Routes>
  )
}
