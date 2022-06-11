import { Transaction } from './Transaction'

export type CreateTransactionInput = Pick<Transaction,
  'clientId' | 'amount' | 'type' | 'service' | 'transactionServiceId' | 'transactionDate'
  >;

export interface ITransactionService {
  create(transaction: CreateTransactionInput): Promise<Transaction>
  listByClientId(clientId: string): Promise<Transaction[]>
}

export interface ITransactionRepository {
  create(transaction: CreateTransactionInput): Promise<Transaction>
  getByClientId(clientId: string): Promise<Transaction[]>
}
