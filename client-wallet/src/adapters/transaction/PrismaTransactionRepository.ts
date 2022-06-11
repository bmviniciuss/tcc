import { CreateTransactionInput, ITransactionRepository } from '../../core/transaction/transaction.interfaces'
import { ServiceType, Transaction, TransactionType } from '../../core/transaction/Transaction'
import { PrismaClient, PrismaServiceType, PrismaTransactionType, PrismaTransaction } from '@prisma/client'
import logger from '../../utils/logger'

export default class PrismaTransactionRepository implements ITransactionRepository {
  private readonly logger = logger.child({ label: PrismaTransactionRepository.name })

  constructor (private readonly prisma: PrismaClient) {}

  async create (transaction: CreateTransactionInput): Promise<Transaction> {
    this.logger.info('Saving transaction to database')
    try {
      const data = await this.prisma.prismaTransaction.create({
        data: {
          clientId: transaction.clientId,
          amount: transaction.amount,
          type: transaction.type,
          transactionServiceId: transaction.transactionServiceId,
          service: transaction.service,
          transactionDate: transaction.transactionDate
        }
      })
      this.logger.info('Transaction saved to database')
      return this.mapPrismaTransactionToCoreTransaction(data)
    } catch (error: any) {
      this.logger.error('Error saving transaction to database')
      this.logger.error(error?.message)
      throw new Error('Error while saving transaction')
    }
  }

  private mapPrismaTransactionToCoreTransaction (data: PrismaTransaction): Transaction {
    return {
      id: data.id,
      clientId: data.clientId,
      amount: data.amount,
      type: PrismaTransactionRepository.mapTransactionType(data.type),
      transactionServiceId: data.transactionServiceId ?? undefined,
      transactionDate: data.transactionDate,
      createdAt: data.createdAt,
      service: PrismaTransactionRepository.mapServiceType(data.service)
    }
  }

  async getByClientId (clientId: string): Promise<Transaction[]> {
    try {
      const data = await this.prisma.prismaTransaction.findMany({
        where: { clientId }
      })
      return data.map(this.mapPrismaTransactionToCoreTransaction)
    } catch (error: any) {
      this.logger.error('Error getting transactions by clientId')
      this.logger.error(error?.message)
      throw new Error('Error while getting transactions by clientId')
    }
  }

  private static mapTransactionType (type: PrismaTransactionType): TransactionType {
    switch (type) {
      case 'CREDIT_CARD_PAYMENT':
        return TransactionType.CREDIT_CARD_PAYMENT
      case 'DEBIT_CARD_PAYMENT':
        return TransactionType.DEBIT_CARD_PAYMENT
      case 'WITHDRAWAL':
        return TransactionType.WITHDRAWAL
      default:
        throw new Error('Invalid transaction type')
    }
  }

  private static mapServiceType (service: PrismaServiceType): ServiceType {
    switch (service) {
      case 'CARD_PAYMENT':
        return ServiceType.CARD_PAYMENT
      case 'INTERNAL':
        return ServiceType.INTERNAL
      default:
        throw new Error('Invalid service type')
    }
  }
}
