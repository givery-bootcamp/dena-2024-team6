import { Box, Text } from '@yamada-ui/react'
type attributeDisplayProps = {
  labelName: string
  value: string
}

export const AttributeDisplay = ({ labelName, value }: attributeDisplayProps) => {
  return (
    <Text color="gray.200">
      {labelName}
      <Box as="span" color="black">
        {value}
      </Box>
    </Text>
  )
}
