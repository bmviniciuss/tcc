import axios from 'axios'

import { ENV } from '../../application/config/env'
import {
  GetClientTransactionsResult,
  GetWalletBalanceResult,
  IClientWalletAPI
} from '../../core/client-wallet/client-wallet.interface'
import logger from '../../utils/logger'

export default class AxiosClientWalletApi implements IClientWalletAPI {
  private readonly logger = logger.child({ label: AxiosClientWalletApi.name })

  async getClientTransactions (clientId: string): Promise<GetClientTransactionsResult> {
    try {
      this.logger.info('Calling client wallet api to get client transactions')
      const base = ENV.CLIENT_WALLET_HOST
      const { data } = await axios.get(`http://${base}/api/clients/${clientId}/transactions`)

      return {
        transactions: data
      }
    } catch (error:any) {
      this.logger.error('Error in create card http request')
      this.logger.error(error)
      let message = 'Internal error while fetching client transactions'
      if (error instanceof Error) message = error?.message
      throw new Error(message) // TODO: Better error handling
    }
  }

  getWalletBalance (clientId: string): Promise<GetWalletBalanceResult> {
    throw new Error('Not Implemented')
  }
}
