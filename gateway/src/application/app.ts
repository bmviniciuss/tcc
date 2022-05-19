import fastify from 'fastify'

import routes from './routes'

const app = fastify({ logger: false })

app.register(routes, { prefix: 'api' })

export default app
