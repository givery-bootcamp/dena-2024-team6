import { Routes, Route } from 'react-router-dom'

import { HelloWorld } from '../features/helloworld'
import { PostsRoute } from '../pages/posts/PostsRoute'

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<PostsRoute />} />
      <Route path="/posts" element={<HelloWorld />} />
    </Routes>
  )
}
