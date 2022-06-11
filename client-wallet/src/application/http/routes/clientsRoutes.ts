import { PrismaClient } from '@prisma/client'
import { FastifyPluginCallback } from 'fastify'

import PrismaTransactionRepository from '../../../adapters/transaction/PrismaTransactionRepository'
import TransactionService from '../../../core/transaction/TransactionService'
import logger from '../../../utils/logger'
import PresentationTransactionMapper from '../mappers/PresentationTransactionMapper'

export const clientsRoutes = (prisma: PrismaClient): FastifyPluginCallback => {
  return async (fastify, _, done) => {
    const l = logger.child({ label: clientsRoutes.name })

    fastify.get<{ Params: {clientId: string}}>('/:clientId/transactions', {
      schema: {
        params: {
          clientId: {
            type: 'string'
          }
        }
      }
    }, async (request, reply) => {
      try {
        l.info('Received transaction request')
        const prismaTransactionRepository = new PrismaTransactionRepository(prisma)
        const service = new TransactionService(prismaTransactionRepository)
        const transactions = await service.listByClientId(request.params.clientId)
        const presentation = transactions.map(PresentationTransactionMapper.mapTransactionToPresentation)
        return reply.status(200).send(presentation)
      } catch (error: any) {
        l.error('Error creating transaction', error)
        return reply.status(500).send({ error: error?.message })
      }
    })
  }
}
