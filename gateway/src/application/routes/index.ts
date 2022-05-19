import { FastifyInstance } from 'fastify'

import cardsRoutes from './cards'

export default async function routes (fastify: FastifyInstance) {
  fastify.get('/health', async () => {
    return { alive: true }
  })

  fastify.register(cardsRoutes, { prefix: 'cards' })
}
