import { FastifyInstance, FastifyPluginCallback } from 'fastify'
import logger from '../../utils/logger'
import { PrismaClient } from '@prisma/client'
import { transactionsRoutes } from './routes/transactionsRoutes'
import { clientsRoutes } from './routes/clientsRoutes'

const routes = (prisma: PrismaClient): FastifyPluginCallback => {
  return async (fastify: FastifyInstance, _, done) => {
    const l = logger.child({ label: 'routes' })

    fastify.get('/health', async () => {
      return { ok: true }
    })

    fastify.register(transactionsRoutes(prisma), { prefix: '/transactions' })
    fastify.register(clientsRoutes(prisma), { prefix: '/clients' })

    done()
  }
}

export default routes
