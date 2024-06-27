import { Button, Center, Divider, Modal, ModalBody, ModalFooter, ModalHeader } from '@yamada-ui/react'
import { Link } from 'react-router-dom'

interface props {
  isOpen: boolean
  onClose: () => void
  signoutMutate: () => void
  refetch: () => void
}

export const SignoutModal = ({ isOpen, onClose, signoutMutate, refetch }: props) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <Center>
        <ModalHeader>警告</ModalHeader>
      </Center>
      <Divider variant="solid" my={2} />
      <ModalBody>サインアウトを行います。よろしいですか？</ModalBody>
      <Divider variant="solid" my={2} />
      <ModalFooter>
        <Button variant="ghost" onClick={onClose}>
          とじる
        </Button>
        <Link to="/">
          <Button
            onClick={() => {
              signoutMutate()
              refetch()
              onClose()
            }}
            colorScheme="primary"
          >
            サインアウト
          </Button>
        </Link>
      </ModalFooter>
    </Modal>
  )
}
