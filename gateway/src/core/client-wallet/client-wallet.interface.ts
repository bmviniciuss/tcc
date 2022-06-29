import { ClientWalletTransaction } from './ClientWalletTransaction'

export type GetClientTransactionsResult = {
  transactions: ClientWalletTransaction[]
}

export type GetWalletBalanceResult = {
  balance: number
}

export interface IClientWalletService {
  getClientTransactions(clientId: string): Promise<GetClientTransactionsResult>
  getWalletBalance(clientId: string): Promise<GetWalletBalanceResult>
}

export interface IClientWalletAPI {
  getClientTransactions(clientId: string): Promise<GetClientTransactionsResult>
  getWalletBalance(clientId: string): Promise<GetWalletBalanceResult>
}
