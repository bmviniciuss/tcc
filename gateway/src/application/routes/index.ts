import { FastifyInstance } from 'fastify'

import logger from '../../utils/logger'
import transactionsRoutes from './transactions'

export default async function routes (fastify: FastifyInstance) {
  fastify.get('/health', async () => {
    logger.info('GET /health was called')
    return { alive: true }
  })

  fastify.register(transactionsRoutes, { prefix: 'transactions' })
}
