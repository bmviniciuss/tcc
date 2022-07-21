import Fastify from 'fastify'
import { TypeBoxTypeProvider } from '@fastify/type-provider-typebox'
import { Type } from '@sinclair/typebox'

import prismaPlugin from './plugins/prisma'

const fastify = Fastify({
  logger: {
    transport: {
      target: 'pino-pretty',
      options: {
        colorize: true
      }
    }
  },
  ajv: {
    customOptions: {
      strict: 'log',
      keywords: ['kind', 'modifier']
    }
  }
}).withTypeProvider<TypeBoxTypeProvider>()

fastify.register(prismaPlugin)

fastify.route({
  method: 'POST',
  url: '/api/benchmarks',
  schema: {
    body: Type.Object({
      name: Type.String(),
      vus: Type.Integer({ minimum: 0 }),
      duration: Type.Optional(Type.String()),
      iterations: Type.Optional(Type.Integer())
    })
  },
  handler: async (request, reply) => {
    const { body, log } = request
    log.info('Creating benchmark')
    try {
      const benchmark = await fastify.prisma.benchmark.create({
        data: {
          name: body.name,
          vus: body.vus,
          duration: body?.duration ?? null,
          iterations: body?.iterations ?? null
        }
      })
      return reply.code(201).send(benchmark)
    } catch (error) {
      log.error(error, 'Internal error')
      return reply.code(500).send({ message: 'Internal Error' })
    }
  }
})

fastify.route({
  method: 'GET',
  url: '/api/benchmarks',
  handler: async (request, reply) => {
    const { log } = request
    log.info('List all benchmarks')
    try {
      const benchmarks = await fastify.prisma.benchmark.findMany({
        orderBy: {
          createdAt: 'asc'
        }
      })
      return reply.code(200).send({
        content: benchmarks
      })
    } catch (error) {
      log.error(error, 'Internal error')
      return reply.code(500).send({ message: 'Internal Error' })
    }
  }
})

export default fastify
