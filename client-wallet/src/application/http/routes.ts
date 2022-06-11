import { FastifyInstance, FastifyPluginCallback } from 'fastify'
import { FromSchema } from 'json-schema-to-ts'

import logger from '../../utils/logger'
import { TransactionType, ServiceType } from '../../core/transaction/Transaction'
import { PrismaClient } from '@prisma/client'
import TransactionService from '../../core/transaction/TransactionService'
import PrismaTransactionRepository from '../../adapters/transaction/PrismaTransactionRepository'
import { CreateTransactionInput } from '../../core/transaction/transaction.interfaces'

const TransactionRequestSchema = {
  type: 'object',
  required: ['client_id', 'amount', 'type', 'payment_id', 'service', 'transaction_date'],
  properties: {
    client_id: { type: 'string' },
    amount: { type: 'number' },
    type: { type: 'string', enum: Object.values(TransactionType) },
    payment_id: { type: 'string' },
    service: { type: 'string', enum: Object.values(ServiceType) },
    transaction_date: { type: 'string', format: 'date-time' }
  }
} as const

const routes = (prisma: PrismaClient): FastifyPluginCallback => {
  return async (fastify: FastifyInstance, _, done) => {
    const l = logger.child({ label: 'routes' })

    fastify.get('/health', async () => {
      return { ok: true }
    })

    fastify.post<{ Body: FromSchema<typeof TransactionRequestSchema>}>('/transactions', {
      schema: {
        body: TransactionRequestSchema
      }
    }, async (request, reply) => {
      try {
        l.info('Received transaction request')
        const prismaTransactionRepository = new PrismaTransactionRepository(prisma)
        const service = new TransactionService(prismaTransactionRepository)
        const createTransactionInput: CreateTransactionInput = {
          clientId: request.body.client_id,
          amount: request.body.amount,
          type: request.body.type,
          transactionServiceId: request.body.payment_id,
          service: request.body.service,
          transactionDate: new Date(request.body.transaction_date)
        }
        const transaction = await service.create(createTransactionInput)
        return reply.status(201).send(transaction)
      } catch (error: any) {
        l.error('Error creating transaction', error)
        return reply.status(500).send({ error: error?.message })
      }
    })

    done()
  }
}

export default routes
