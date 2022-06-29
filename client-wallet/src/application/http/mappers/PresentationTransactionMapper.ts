import { FromSchema } from 'json-schema-to-ts'

import { Transaction } from '../../../core/transaction/Transaction'
import { CreateTransactionInput } from '../../../core/transaction/transaction.interfaces'
import { PresentationTransaction } from '../routes/models/transaction'
import { TransactionRequestSchema } from '../routes/schemas/transaction'

export default class PresentationTransactionMapper {
  static mapTransactionToPresentation (transaction: Transaction): PresentationTransaction {
    return {
      id: transaction.id,
      client_id: transaction.clientId,
      amount: transaction.amount,
      type: transaction.type,
      transaction_service_id: transaction.transactionServiceId,
      service: transaction.service,
      transaction_date: transaction.transactionDate,
      created_at: transaction.createdAt
    }
  }

  static mapRequestInputToCreateTransactionInput (body: FromSchema<typeof TransactionRequestSchema>): CreateTransactionInput {
    return {
      clientId: body.client_id,
      amount: body.amount,
      type: body.type,
      transactionServiceId: body?.transaction_service_id ?? undefined,
      service: body.service,
      transactionDate: new Date(body.transaction_date)
    }
  }
}
