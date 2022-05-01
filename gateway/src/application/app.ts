import fastify from 'fastify'

const app = fastify({ logger: true })

app.get('/health', async (req, res) => {
  return { alive: true }
})

export default app
