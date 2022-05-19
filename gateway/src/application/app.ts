import fastifySwagger from '@fastify/swagger'
import fastify from 'fastify'

import routes from './routes'

const app = fastify({ logger: false })
app.register(fastifySwagger, {
  routePrefix: '/docs',
  swagger: {
    info: {
      title: 'TCC Payment Gateway',
      description: 'Docs for the API Gateway',
      version: '0.1.0'
    },
    host: 'localhost',
    schemes: ['http'],
    consumes: ['application/json'],
    produces: ['application/json'],
    tags: [
      { name: 'cards', description: 'Cards related end-points' }
    ]
  },
  uiConfig: {
    docExpansion: 'full',
    deepLinking: true,
    filter: true
  },
  uiHooks: {
    onRequest: function (request, reply, next) { next() },
    preHandler: function (request, reply, next) { next() }
  },
  staticCSP: true,
  transformStaticCSP: (header) => header,
  exposeRoute: true
})

app.register(routes, { prefix: 'api' })

export default app
