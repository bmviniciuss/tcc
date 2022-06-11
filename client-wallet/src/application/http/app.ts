import fastify from 'fastify'
import routes from './routes'
import { PrismaClient } from '@prisma/client'

function makeApp (prismaClient: PrismaClient) {
  const app = fastify({ logger: false })
  app.register(routes(prismaClient), { prefix: '/api' })
  return app
}

export default makeApp
