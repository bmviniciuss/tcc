import { Server } from '@grpc/grpc-js'

import { ClientWalletServiceImpl } from './services/ClientWalletService'
import TransactionService from '../core/transaction/TransactionService'
import { PrismaClient } from '@prisma/client'
import { ClientWalletService } from './pb/client_wallet_grpc_pb'
import PrismaTransactionRepository from '../adapters/transaction/PrismaTransactionRepository'
import logger from '../utils/logger'

export function makeGRPCServer (prismaClient: PrismaClient) {
  const l = logger.child({ label: makeGRPCServer.name })
  l.info('Creating GRPC server')
  const server = new Server({})

  const transactionRepo = new PrismaTransactionRepository(prismaClient)
  const transactionService = new TransactionService(transactionRepo)

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  server.addService(ClientWalletService, new ClientWalletServiceImpl(transactionService))

  return server
}
