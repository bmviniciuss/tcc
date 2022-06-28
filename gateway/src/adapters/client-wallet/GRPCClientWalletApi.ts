import { credentials, ServiceError } from '@grpc/grpc-js'

import { ENV } from '../../application/config/env'
import { GetClientTransactionsResult, IClientWalletAPI } from '../../core/client-wallet/client-wallet.interface'
import { ClientWalletTransaction } from '../../core/client-wallet/ClientWalletTransaction'
import logger from '../../utils/logger'
import { ClientWalletClient } from '../grpc/pb/client_wallet_grpc_pb'
import { GetClientTransactionsInput, ServiceTypeEnum, TransactionTypeEnum } from '../grpc/pb/client_wallet_pb'

export default class GRPCClientWalletApi implements IClientWalletAPI {
  private readonly logger = logger.child({ label: GRPCClientWalletApi.name })

  async getClientTransactions (clientId: string): Promise<GetClientTransactionsResult> {
    return new Promise((resolve, reject) => {
      this.logger.info('Process started')
      try {
        this.logger.info(`Calling client-wallet gRPC client at: ${ENV.CLIENT_WALLET_HOST}`)
        const client = new ClientWalletClient(ENV.CLIENT_WALLET_HOST, credentials.createInsecure())
        const request = new GetClientTransactionsInput().setClientId(clientId)
        client.getClientTransactions(request, (error: ServiceError | null, res) => {
          if (error !== null) {
            this.logger.error('Error while making request')
            this.logger.error(error)
            const fallbackMessage = 'An error occur while fetching client transactions'
            return reject(new Error(GRPCClientWalletApi.getGRPCErrorMessage(error, fallbackMessage)))
          }

          this.logger.info('Successfully fetched client transactions from gRPC service')
          const transactions = res.getTransactionsList().map(t => {
            const tt: ClientWalletTransaction = {
              id: t.getId(),
              clientId: t.getClientId(),
              amount: t.getAmount(),
              createdAt: new Date(t.getCreatedAt()),
              service: ServiceTypeEnum[t.getService()],
              transactionDate: new Date(t.getTransactionDate()),
              type: TransactionTypeEnum[t.getType()]
            }
            return tt
          })
          const result: GetClientTransactionsResult = {
            transactions
          }
          return resolve(result)
        })
      } catch (error: any) {
        this.logger.error('Error while building the get client transactions gRPC request')
        this.logger.error(error)
        reject(new Error('Internal error while processing get client transactions api call'))
      }
    })
  }

  private static getGRPCErrorMessage (error: any, fallback: string): string {
    if (error?.details) {
      return error.details
    }

    if (error?.message) {
      return error?.message
    }

    return fallback
  }
}
