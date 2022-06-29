import { ServiceType, TransactionType } from '../../../../core/transaction/Transaction'

export type PresentationTransaction = {
  id: string
  client_id: string
  amount: number
  type: TransactionType
  transaction_service_id?: string
  service: ServiceType
  transaction_date: Date
  created_at: Date
}
