import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react'
import { useGetUser } from '../../api/api'
import { User } from '../../api/model'

interface UserContextType {
  user: User | undefined
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>
  refetchUser: () => void
}

const UserContext = createContext<UserContextType | undefined>(undefined)

export const UserProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useState<User | undefined>(undefined)
  const { data, isLoading, isError, refetch } = useGetUser()

  useEffect(() => {
    if (!isLoading && !isError && data) {
      setUser(data)
    }
  }, [data, isLoading, isError])

  return <UserContext.Provider value={{ user, setUser, refetchUser: refetch }}>{children}</UserContext.Provider>
}

export const useUser = () => {
  const context = useContext(UserContext)
  if (!context) {
    throw new Error('useUser must be used within a UserProvider')
  }
  return context
}
