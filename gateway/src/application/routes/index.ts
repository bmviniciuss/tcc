import { FastifyInstance } from 'fastify'

import cardsRoutes from './cards'
import paymentRoutes from './payment'

export default async function routes (fastify: FastifyInstance) {
  fastify.get('/health', async () => {
    return { alive: true }
  })

  fastify.register(cardsRoutes, { prefix: 'cards' })
  fastify.register(paymentRoutes, { prefix: 'payments' })
}
