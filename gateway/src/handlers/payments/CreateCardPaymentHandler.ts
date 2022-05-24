import { FastifyReply, FastifyRequest } from 'fastify'

import { CreateCardPaymentInput, ICardPaymentService } from '../../core/card-payment/card-payment.interface'
import { CardPayment, PaymentType } from '../../core/card-payment/CardPayment'
import logger from '../../utils/logger'

export type CreateCardPaymentRequest = {
  client_id: string
  payment_type: PaymentType
  payment_date: string
  amount: number
  payment_info: {
    card_token: string
  }
}

export type CreateCardPaymentResponse = {
  client_id: string
  payment_type: PaymentType
  payment_date: string
  amount: number
  payment_info: {
    masked_number: string
  }
}

export class CreateCardPaymentHandler {
  private readonly logger = logger.child({ label: CreateCardPaymentHandler.name })

  constructor (private readonly cardPaymentService: ICardPaymentService) {}

  async handle (req: FastifyRequest, res: FastifyReply) {
    try {
      this.logger.info('Process started')
      const body = req.body as CreateCardPaymentRequest
      const createCardInput = CreateCardPaymentHandler.mapToCreateCardPaymentInput(body)
      const cardPayment = await this.cardPaymentService.create(createCardInput)
      return CreateCardPaymentHandler.mapToPresentationPayment(cardPayment)
    } catch (error) {
      console.error(error)
      return res.status(500).send({
        message: 'Internal Error while processing card payment'
      })
    }
  }

  static mapToCreateCardPaymentInput ({ client_id, payment_date, payment_info, payment_type, amount }: CreateCardPaymentRequest): CreateCardPaymentInput {
    return {
      clientId: client_id,
      amount,
      paymentType: payment_type,
      paymentDate: payment_date,
      paymentInfo: {
        cardToken: payment_info.card_token
      }
    }
  }

  static mapToPresentationPayment ({ paymentType, paymentInfo, paymentDate, amount, clientId }: CardPayment): CreateCardPaymentResponse {
    return {
      client_id: clientId,
      amount,
      payment_type: paymentType,
      payment_date: paymentDate,
      payment_info: {
        masked_number: paymentInfo.maskedNumber
      }
    }
  }
}
