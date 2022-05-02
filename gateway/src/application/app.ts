import fastify from 'fastify'

import routes from './routes/'

const app = fastify({ logger: true })

app.register(routes, { prefix: 'api' })

export default app
