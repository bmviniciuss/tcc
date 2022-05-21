import { FastifyInstance } from 'fastify'

import AxiosHttpCardPaymentAPI from '../../adapters/card-payment/AxiosHttpCardPaymentAPI'
import CardPaymentService from '../../core/card-payment/CardPaymentService'
import { CreateCardPaymentHandler } from '../../handlers/payments/cardPaymentHandlers'

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
      const cardPaymentApi = new AxiosHttpCardPaymentAPI()
      const cardPaymentService = new CardPaymentService(cardPaymentApi)
      const handler = new CreateCardPaymentHandler(cardPaymentService)
      return handler.handle(req, res)
    }
  })
}
