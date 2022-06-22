import { ClientWalletTransaction } from './ClientWalletTransaction'

export type GetClientTransactionsResult = {
  transactions: ClientWalletTransaction[]
}

export interface IClientWalletService {
  getClientTransactions(clientId: string): Promise<GetClientTransactionsResult>
}

export interface IClientWalletAPI {
  getClientTransactions(clientId: string): Promise<GetClientTransactionsResult>
}
