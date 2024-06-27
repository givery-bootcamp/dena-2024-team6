import { defineConfig } from 'orval'

export default defineConfig({
  api: {
    output: {
      mode: 'split',
      target: 'src/api/hooks.ts',
      schemas: 'src/api/model',
      client: 'react-query',
      clean: true,
      mock: true,
      override: {
        mutator: {
          path: './src/shared/libs/axios.ts',
          name: 'customInstance'
        }
      }
    },
    input: {
      target: '../docs/api.yaml'
    }
  }
})
