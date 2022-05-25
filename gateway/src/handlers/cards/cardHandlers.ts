import { FastifyReply, FastifyRequest } from 'fastify'

import { CreateCardInput, ICardService } from '../../core/card/card.interface'
import logger from '../../utils/logger'

export type CreateCardRequestBody = {
  cardholder_name: string
  is_debit: boolean
  is_credit: boolean
}

export class CreateCardHandler {
  private readonly logger = logger.child({ label: CreateCardHandler.name })

  constructor (private readonly cardService: ICardService) {}

  async handle (req: FastifyRequest, res: FastifyReply) {
    try {
      this.logger.info('Process started')
      const { is_credit, is_debit, cardholder_name } = req.body as CreateCardRequestBody
      const createCardInput: CreateCardInput = {
        cardholderName: cardholder_name,
        isCredit: is_credit,
        isDebit: is_debit
      }
      return await this.cardService.create(createCardInput)
    } catch (error) {
      console.error(error)
      return res.status(500).send({
        message: 'Internal Error while processing card payment'
      })
    }
  }
}
