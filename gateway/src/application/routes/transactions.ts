import { FastifyInstance } from 'fastify'

import logger from '../../utils/logger'

export default async function transactionsRoutes (fastify: FastifyInstance) {
  fastify.route({
    method: 'POST',
    url: '/',
    handler: async () => {
      logger.info('POST /transactions called')
      return { transactions: true }
    }
  })
}
