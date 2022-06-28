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

  async getWalletBalance (clientId: string): Promise<GetWalletBalanceResult> {
    try {
      this.logger.info('Calling client wallet api to get client wallet balance')
      const base = ENV.CLIENT_WALLET_HOST
      const url = `http://${base}/api/clients/${clientId}/balance`
      const { data } = await axios.get(url)
      return {
        balance: data?.balance ?? 0
      }
    } catch (error) {
      this.logger.error('Error in http request to get client wallet\'s balance')
      this.logger.error(error)
      let message = 'Internal error while fetching client wallet\'s balance'
      if (error instanceof Error) message = error?.message
      throw new Error(message) // TODO: Better error handling
    }
  }
}
