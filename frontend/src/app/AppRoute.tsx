import { Routes, Route } from 'react-router-dom'

import { HelloWorld } from '../features/helloworld'
import { PostsRoute } from '../pages/posts/PostsRoute'
import { SigninRoute } from '../pages/signin/SigninRoute'

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/signin" element={<SigninRoute />} />
      <Route path="/" element={<PostsRoute />} />
      <Route path="/posts" element={<HelloWorld />} />
    </Routes>
  )
}
