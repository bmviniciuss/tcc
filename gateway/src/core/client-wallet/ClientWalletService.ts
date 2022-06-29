import logger from '../../utils/logger'
import {
  GetClientTransactionsResult,
  GetWalletBalanceResult,
  IClientWalletAPI,
  IClientWalletService
} from './client-wallet.interface'

export default class ClientWalletService implements IClientWalletService {
  private readonly logger = logger.child({ name: ClientWalletService.name })

  constructor (private readonly clientWalletApi: IClientWalletAPI) {}

  getClientTransactions (clientId: string): Promise<GetClientTransactionsResult> {
    this.logger.info('Service called')
    return this.clientWalletApi.getClientTransactions(clientId)
  }

  getWalletBalance (clientId: string): Promise<GetWalletBalanceResult> {
    this.logger.info('Process started')
    return this.clientWalletApi.getWalletBalance(clientId)
  }
}
