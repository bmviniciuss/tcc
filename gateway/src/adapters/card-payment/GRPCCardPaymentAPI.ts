import { credentials, ServiceError } from '@grpc/grpc-js'

import { ENV } from '../../application/config/env'
import { CardPaymentAPI, CreateCardPaymentInput } from '../../core/card-payment/card-payment.interface'
import { CardPayment, PaymentType } from '../../core/card-payment/CardPayment'
import logger from '../../utils/logger'
import { CardPaymentClient } from '../grpc/pb/card-payment_grpc_pb'
import { Payment, PaymentInfoInput, PaymentTypeEnum, ProcessCardPaymentInput } from '../grpc/pb/card-payment_pb'

export default class GRPCCardPaymentAPI implements CardPaymentAPI {
  private readonly logger = logger.child({ label: GRPCCardPaymentAPI.name })
  private readonly client: CardPaymentClient

  constructor () {
    this.client = new CardPaymentClient(ENV.CARD_PAYMENT_HOST, credentials.createInsecure())
  }

  async create ({ paymentDate, paymentInfo, paymentType, clientId, amount }: CreateCardPaymentInput): Promise<CardPayment> {
    this.logger.info('Process started')

    try {
      const request =
        GRPCCardPaymentAPI.buildProcessPaymentCardRequest(amount, clientId, paymentDate, paymentType, paymentInfo)

      const res: Payment = await new Promise((resolve, reject) => {
        this.client.proccessCardPayment(request, (error: ServiceError | null, res: Payment) => {
          if (error !== null) return reject(error)
          return resolve(res)
        })
      })

      return {
        id: res.getId(),
        clientId: res.getClientid(),
        amount: res.getAmount(),
        paymentDate: res.getPaymentdate(),
        paymentInfo: {
          maskedNumber: res.getPaymentinfo()?.getMaskednumber() ?? ''
        },
        paymentType: res.getPaymenttype() as PaymentType
      }
    } catch (error) {
      this.logger.error('Error while building the process card payment gRPC request')
      this.logger.error(error)
      throw new Error('Internal error while processing card payment request')
    }
  }

  private static buildProcessPaymentCardRequest (amount: number, clientId: string, paymentDate: string, paymentType: PaymentType, paymentInfo: { cardToken: string }) {
    return new ProcessCardPaymentInput()
      .setAmount(amount)
      .setClientid(clientId)
      .setPaymentdate(paymentDate)
      .setPaymenttype(PaymentTypeEnum[paymentType])
      .setPaymentinfo(
        new PaymentInfoInput()
          .setCardtoken(paymentInfo.cardToken)
      )
  }

  private static getGRPCErrorMessage (error: any, fallback: string): string {
    if (error?.details) {
      return error.details
    }

    if (error?.message) {
      return error?.message
    }

    return fallback
  }
}
