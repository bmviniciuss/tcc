export enum TransactionType {
  CREDIT_CARD_PAYMENT = 'CREDIT_CARD_PAYMENT',
  DEBIT_CARD_PAYMENT = 'DEBIT_CARD_PAYMENT',
  WITHDRAWAL = 'WITHDRAWAL'
}

export enum ServiceType {
  CARD_PAYMENT = 'CARD_PAYMENT',
  INTERNAL = 'INTERNAL'
}

export type Transaction = {
  id: string
  clientId: string
  amount: number
  type: TransactionType
  transactionServiceId?: string
  service: ServiceType
  transactionDate: Date
  createdAt: Date
}
