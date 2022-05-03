import { FastifyInstance, FastifyRequest, FastifyReply } from 'fastify'

import HttpCardPaymentAPI from '../../adapters/card/httpCardPaymentAPI'
import CardPaymentProcessor, { CardPayment, CardProcessor } from '../../core/cardPayment/cardPayment'

export default async function paymentRoutes (fastify: FastifyInstance) {
  fastify.route({
    method: 'POST',
    url: '/card',
    handler: async (req, res) => {
      const cardPaymentAPI = new HttpCardPaymentAPI()
      const cardCaymentService = new CardPaymentProcessor(cardPaymentAPI)
      const controller = new ProcessCardPaymentController(cardCaymentService)
      return controller.handle(req, res)
    }
  })
}

class ProcessCardPaymentController {
  constructor (private readonly cardPaymentService: CardProcessor) {}

  async handle (req: FastifyRequest, res: FastifyReply) {
    try {
      const cardPaymentDTO: CardPayment = req.body as CardPayment
      const payment = await this.cardPaymentService.processCardPayment(cardPaymentDTO)
      return payment
    } catch (error) {
      return res.status(500).send({
        message: 'Internal Error while processing card payment'
      })
    }
  }
}
