import { PrismaClient } from '@prisma/client'
import { FastifyInstance, FastifyPluginCallback } from 'fastify'

import { clientsRoutes } from './routes/clientsRoutes'
import { transactionsRoutes } from './routes/transactionsRoutes'

const routes = (prisma: PrismaClient): FastifyPluginCallback => {
  return async (fastify: FastifyInstance, _, done) => {
    fastify.get('/health', async () => {
      return { ok: true }
    })

    fastify.register(transactionsRoutes(prisma), { prefix: '/transactions' })
    fastify.register(clientsRoutes(prisma), { prefix: '/clients' })

    done()
  }
}

export default routes
