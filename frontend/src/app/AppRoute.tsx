import { Routes, Route } from 'react-router-dom'

import { PostsRoute } from '../pages/posts/PostsRoute'
import { PostDetailRoute } from '../pages/posts/detail/PostDetailRoute'
import { CreatePostRoute } from '../pages/posts/CreatePostRoute'
import { UpdatePostRoute } from '../pages/posts/update/UpdatePostRoute'
import { SigninRoute } from '../pages/signin/SigninRoute'
import { SignupRoute } from '../pages/signup/SignupRoute'

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<PostsRoute />} />
      <Route path="/signin" element={<SigninRoute />} />
      <Route path="/signup" element={<SignupRoute />} />
      <Route path="/posts/:id" element={<PostDetailRoute />} />
      <Route path="/posts/:id/edit" element={<UpdatePostRoute />} />
      <Route path="/posts/new" element={<CreatePostRoute />} />
    </Routes>
  )
}
