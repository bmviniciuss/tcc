import { FastifyInstance } from 'fastify'

import AxiosHttpCardAPI from '../../adapters/card/AxiosHttpCardAPI'
import GRPCCardApi from '../../adapters/card/GRPCCardAPI'
import CardService from '../../core/card/CardService'
import { CreateCardHandler } from '../../handlers/cards/cardHandlers'

export default async function cardsRoutes (fastify: FastifyInstance) {
  fastify.route({
    method: 'POST',
    url: '/',
    schema: {
      description: 'Create a new card',
      tags: ['cards'],
      body: {
        type: 'object',
        required: ['cardholder_name', 'is_credit', 'is_debit'],
        properties: {
          cardholder_name: { type: 'string' },
          is_credit: { type: 'boolean' },
          is_debit: { type: 'boolean' }
        }
      }
    },
    handler: async (req, res) => {
      const grpcCardApi = new GRPCCardApi()
      // const cardAPI = new AxiosHttpCardAPI()
      const cardService = new CardService(grpcCardApi)
      const handler = new CreateCardHandler(cardService)
      return handler.handle(req, res)
    }
  })
}
