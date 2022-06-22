import { FastifyInstance } from 'fastify'

import AxiosClientWalletApi from '../../adapters/client-wallet/AxiosClientWalletApi'
import ClientWalletService from '../../core/client-wallet/ClientWalletService'
import logger from '../../utils/logger'
import { ENV } from '../config/env'

type ClientIdParams = {
  clientId: string
}

export default async function clientWalletRoutes (fastify: FastifyInstance) {
  fastify.get<{ Params: ClientIdParams }>('/:clientId/transactions', async (req, res) => {
    const l = logger.child({ label: 'clientTransactions.GET.handler' })
    l.info('Process started')
    const clientWalletApi = (() => {
      if (ENV.ENABLE_GRPC) {
        l.info('Using GRPC Api for card generation')
        throw new Error('Not implemented')
      }
      l.info('Using HTTP Api for card generation')
      return new AxiosClientWalletApi()
    })()

    const s = new ClientWalletService(clientWalletApi)
    return s.getClientTransactions(req.params.clientId as string)
  })
}
