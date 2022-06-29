import { FastifyInstance } from 'fastify'

import AxiosClientWalletApi from '../../adapters/client-wallet/AxiosClientWalletApi'
import GRPCClientWalletApi from '../../adapters/client-wallet/GRPCClientWalletApi'
import ClientWalletService from '../../core/client-wallet/ClientWalletService'
import logger from '../../utils/logger'
import { ENV } from '../config/env'

type ClientIdParams = {
  clientId: string
}

function buildClientWalletApi () {
  const l = logger.child({ child: 'buildClientWalletApi' })
  return (() => {
    if (ENV.ENABLE_GRPC) {
      l.info('Using GRPC Api for card generation')
      return new GRPCClientWalletApi()
    }
    l.info('Using HTTP Api for card generation')
    return new AxiosClientWalletApi()
  })()
}

export default async function clientWalletRoutes (fastify: FastifyInstance) {
  fastify.get<{ Params: ClientIdParams }>('/:clientId/transactions', async (req, res) => {
    const l = logger.child({ label: 'clientWallet.Transactions.GET.handler' })
    l.info('Process started')
    const clientWalletApi = buildClientWalletApi()

    const s = new ClientWalletService(clientWalletApi)
    return s.getClientTransactions(req.params.clientId as string)
  })

  fastify.get<{ Params: ClientIdParams }>('/:clientId/balance', async (req, res) => {
    const l = logger.child({ label: 'clientWallet.Balance.GET.handler' })
    l.info('Process started')
    const clientWalletApi = buildClientWalletApi()
    const s = new ClientWalletService(clientWalletApi)
    return s.getWalletBalance(req.params.clientId as string)
  })
}
