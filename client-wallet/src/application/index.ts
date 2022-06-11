import logger from '../utils/logger'
import { GRPC_ENABLED, PORT } from '../config/env'
import makeHttpApp from './http/app'
import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

async function main () {
  const l = logger.child({ label: 'main' })
  l.info('Initializing Client wallet app')
  if (GRPC_ENABLED) {
    throw new Error('gRPC not implemented')
  }
  l.info('Initializing HTTP server')

  const port = PORT
  const app = makeHttpApp(prisma)

  app.listen({ port, host: '0.0.0.0' }, (err, address) => {
    if (err) {
      l.error(`Error listening on port ${port}`, err)
      process.exit(1)
    }
    l.info(`HTTP Server listening on port ${port}`)
  })
}

main().finally(async () => {
  await prisma.$disconnect()
})
