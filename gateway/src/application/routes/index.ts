import { FastifyInstance } from 'fastify'

import paymentsRoutes from '../../http/handlers/payments'
import logger from '../../utils/logger'

export default async function routes (fastify: FastifyInstance) {
  fastify.get('/health', async () => {
    logger.info('GET /health was called')
    return { alive: true }
  })

  fastify.register(paymentsRoutes, { prefix: 'payments' })
}
