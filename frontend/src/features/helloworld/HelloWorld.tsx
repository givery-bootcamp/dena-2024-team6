import { useEffect } from 'react'

import { useAppDispatch, useAppSelector } from '../../shared/hooks'
import { APIService } from '../../shared/services'
import { Button, Text, Wrap } from '@yamada-ui/react'

export function HelloWorld() {
  // ここからロジック
  const { hello } = useAppSelector((state) => state.hello)
  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(APIService.getHello())
  }, [dispatch])

  // ここからUI
  return (
    <Wrap>
      <Button>{hello?.lang}</Button>
      <Button>{hello?.message}</Button>
    </Wrap>
  )
}
