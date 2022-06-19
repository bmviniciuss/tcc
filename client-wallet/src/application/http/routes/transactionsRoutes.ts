import { PrismaClient } from '@prisma/client'
import { FastifyPluginCallback } from 'fastify'
import { FromSchema } from 'json-schema-to-ts'

import PrismaTransactionRepository from '../../../adapters/transaction/PrismaTransactionRepository'
import { CreateTransactionInput } from '../../../core/transaction/transaction.interfaces'
import TransactionService from '../../../core/transaction/TransactionService'
import logger from '../../../utils/logger'
import PresentationTransactionMapper from '../mappers/PresentationTransactionMapper'
import { TransactionRequestSchema } from './schemas/transaction'

export const transactionsRoutes = (prisma: PrismaClient): FastifyPluginCallback => {
  return async (fastify) => {
    const l = logger.child({ label: transactionsRoutes.name })

    fastify.post<{ Body: FromSchema<typeof TransactionRequestSchema>}>('/', {
      schema: {
        body: TransactionRequestSchema
      }
    }, async (request, reply) => {
      try {
        l.info('Received transaction request')
        const prismaTransactionRepository = new PrismaTransactionRepository(prisma)
        const service = new TransactionService(prismaTransactionRepository)
        const createTransactionInput: CreateTransactionInput = PresentationTransactionMapper.mapRequestInputToCreateTransactionInput(request.body)
        const transaction = await service.create(createTransactionInput)
        const presentationTransaction = PresentationTransactionMapper.mapTransactionToPresentation(transaction)
        return reply.status(201).send(presentationTransaction)
      } catch (error: any) {
        l.error('Error creating transaction', error)
        return reply.status(500).send({ error: error?.message })
      }
    })
  }
}
