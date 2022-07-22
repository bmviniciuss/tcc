import Fastify from 'fastify'
import { TypeBoxTypeProvider } from '@fastify/type-provider-typebox'
import { Type } from '@sinclair/typebox'

import prismaPlugin from './plugins/prisma'
import ResultSchema from './schemas/Result'

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
      id: Type.Optional(Type.String()),
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
          id: body?.id || undefined,
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

fastify.route({
  method: 'POST',
  url: '/api/results',
  schema: {
    body: ResultSchema
  },
  handler: async (req, reply) => {
    const { log, body } = req
    log.info('Processing new result')
    try {
      const result = await fastify.prisma.result.create({
        data: {
          testDuration: body.state.testRunDurationMs,
          httpReqDurationMin: body.metrics.http_req_duration.values.min,
          httpReqDurationMax: body.metrics.http_req_duration.values.max,
          httpReqDurationAvg: body.metrics.http_req_duration.values.avg,
          httpReqDurationMed: body.metrics.http_req_duration.values.med,
          iterationDurationMin: body.metrics.iteration_duration.values.min,
          iterationDurationMax: body.metrics.iteration_duration.values.max,
          iterationDurationAvg: body.metrics.iteration_duration.values.avg,
          iterationDurationMed: body.metrics.iteration_duration.values.med,
          iterationsCount: body.metrics.iterations.values.count,
          iterationsRate: body.metrics.iterations.values.rate,
          httpReqsCount: body.metrics.http_reqs.values.count,
          httpReqsRate: body.metrics.http_reqs.values.rate,
          executedAt: new Date(body.metadata.testConfig.executedAt),
          benchmark: {
            connect: {
              id: body.metadata.testConfig.id
            }
          }
        }
      })

      return reply.code(201).send(result)
    } catch (error) {
      log.error(error, 'Internal error')
      return reply.code(500).send({ message: 'Internal Error' })
    }
  }
})

export default fastify
