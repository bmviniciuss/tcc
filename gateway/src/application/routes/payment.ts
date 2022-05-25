import { FastifyInstance } from 'fastify'

import AxiosHttpCardPaymentAPI from '../../adapters/card-payment/AxiosHttpCardPaymentAPI'
import GRPCCardPaymentAPI from '../../adapters/card-payment/GRPCCardPaymentAPI'
import CardPaymentService from '../../core/card-payment/CardPaymentService'
import { CreateCardPaymentHandler } from '../../handlers/payments/CreateCardPaymentHandler'
import logger from '../../utils/logger'
import { ENV } from '../config/env'

export default async function paymentRoutes (fastify: FastifyInstance) {
  fastify.route({
    method: 'POST',
    url: '/card',
    schema: {
      description: 'Create Card Payment',
      tags: ['cards'],
      body: {
        type: 'object',
        required: ['client_id', 'amount', 'payment_type', 'payment_date', 'payment_info'],
        properties: {
          client_id: { type: 'string' },
          amount: { type: 'number' },
          payment_type: { type: 'string', enum: ['CREDIT_CARD', 'DEBIT_CARD'] },
          payment_date: { type: 'string' },
          payment_info: {
            type: 'object',
            required: ['card_token'],
            properties: {
              card_token: { type: 'string' }
            }
          }
        }
      }
    },
    handler: async (req, res) => {
      const l = logger.child({ label: 'paymentRoutes.POST.handler' })
      l.info('Process started')
      l.info(ENV)
      const api = (() => {
        if (ENV.ENABLE_GRPC) {
          l.info('Using gRPC Api for payment processing')
          return new GRPCCardPaymentAPI()
        }
        l.info('Using gRPC Api for payment processing')
        return new AxiosHttpCardPaymentAPI()
      })()
      const cardPaymentService = new CardPaymentService(api)
      const handler = new CreateCardPaymentHandler(cardPaymentService)
      return handler.handle(req, res)
    }
  })
}
