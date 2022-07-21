import fastify from './app'

async function main () {
  await fastify.listen({
    port: 3000
  })
}

main()
