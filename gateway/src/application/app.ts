import fastify from 'fastify'

import logger from '../utils/logger'

const app = fastify({ logger: false })

app.get('/health', async () => {
  logger.info('GET /health was called')
  return { alive: true }
})

export default app
