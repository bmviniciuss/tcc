import { FastifyInstance } from 'fastify'

import AxiosHttpCardAPI from '../../adapters/card/AxiosHttpCardAPI'
import GRPCCardApi from '../../adapters/card/GRPCCardAPI'
import CardService from '../../core/card/CardService'
import { CreateCardHandler } from '../../handlers/cards/cardHandlers'
import logger from '../../utils/logger'
import { ENV } from '../config/env'

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
      const l = logger.child({ label: 'createCard.POST.handler' })
      l.info('Process started')
      const cardAPI = (() => {
        if (ENV.ENABLE_GRPC) {
          l.info('Using GRPC Api for card generation')
          return new GRPCCardApi()
        }
        l.info('Using HTTP Api for card generation')
        return new AxiosHttpCardAPI()
      })()

      const cardService = new CardService(cardAPI)
      const handler = new CreateCardHandler(cardService)
      return handler.handle(req, res)
    }
  })
}
