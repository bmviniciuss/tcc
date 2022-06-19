import logger from '../utils/logger'
import { GRPC_ENABLED, PORT } from '../config/env'
import makeHttpApp from './http/app'
import { PrismaClient } from '@prisma/client'
import { makeGRPCServer } from '../grpc/grpcServer'
import { ServerCredentials } from '@grpc/grpc-js'

const prisma = new PrismaClient()
console.log('PROCESS: ', process.env?.GRPC_ENABLED, process.env?.GRPC_ENABLED === 'true')
console.log('GRPC_ENABLED: ', GRPC_ENABLED)
function runGRPC () {
  const l = logger.child({ label: 'runGRPC' })
  const grpcServer = makeGRPCServer(prisma)
  grpcServer.bindAsync(`localhost:${PORT}`, ServerCredentials.createInsecure(), (err) => {
    if (err) {
      l.error(err)
      process.exit(1)
    }
    grpcServer.start()
    l.info(`GRPC server listening on port ${PORT}`)
  })
}

function runHTTP () {
  const l = logger.child({ label: 'runHTTP' })
  const port = PORT
  const app = makeHttpApp(prisma)

  app.listen({ port, host: '0.0.0.0' }, (err) => {
    if (err) {
      l.error(`Error listening on port ${port}`, err)
      process.exit(1)
    }
    l.info(`HTTP Server listening on port ${port}`)
  })
}

async function main () {
  const l = logger.child({ label: 'main' })
  l.info('Initializing Client wallet app')
  if (GRPC_ENABLED) {
    l.info('gRPC enabled')
    runGRPC()
  } else {
    l.info('HTTP enabled')
    runHTTP()
  }
}

main().finally(async () => {
  await prisma.$disconnect()
})
