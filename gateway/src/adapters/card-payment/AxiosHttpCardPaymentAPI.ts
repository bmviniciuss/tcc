import axios from 'axios'

import { CardPaymentAPI, CreateCardPaymentInput } from '../../core/card-payment/card-payment.interface'
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

type CreateCardPaymentResponse = {
  id: string
  client_id: string
  amount: number
  payment_type: string
  payment_date: string
  payment_info: { masked_number: string }
}

export default class AxiosHttpCardPaymentAPI implements CardPaymentAPI {
  private readonly logger = logger.child({ label: AxiosHttpCardPaymentAPI.name })

  async create (input: CreateCardPaymentInput): Promise<CardPayment> {
    try {
      const requestPayload: CreateCardPaymentRequest = {
        client_id: input.clientId,
        amount: input.amount,
        payment_date: input.paymentDate,
        payment_info: {
          card_token: input.paymentInfo.cardToken
        },
        payment_type: input.paymentType
      }
      const URL = 'http://localhost:5555/api/payment' // TODO: move to env variables

      const { data } = await axios.post<CreateCardPaymentResponse>(URL, requestPayload)

      return {
        id: data.id,
        paymentDate: data.payment_date,
        paymentInfo: {
          maskedNumber: data.payment_info.masked_number
        },
        paymentType: data.payment_type as PaymentType,
        amount: data.amount,
        clientId: data.client_id
      }
    } catch (error) {
      this.logger.error('Error in create card http request')
      this.logger.error(error)
      let message = 'Internal error while creating card'
      if (error instanceof Error) message = error?.message
      throw new Error(message) // TODO: Better error handling
    }
  }
}
