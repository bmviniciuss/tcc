import { ServiceType, TransactionType } from '../../../../core/transaction/Transaction'

export const TransactionRequestSchema = {
  type: 'object',
  required: ['client_id', 'amount', 'type', 'service', 'transaction_date'],
  properties: {
    client_id: { type: 'string' },
    amount: { type: 'number' },
    type: { type: 'string', enum: Object.values(TransactionType) },
    transaction_service_id: { type: 'string' },
    service: { type: 'string', enum: Object.values(ServiceType) },
    transaction_date: { type: 'string', format: 'date-time' }
  }
} as const
