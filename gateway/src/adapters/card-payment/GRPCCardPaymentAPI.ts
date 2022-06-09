import { credentials, ServiceError } from '@grpc/grpc-js'

import { ENV } from '../../application/config/env'
import { CardPaymentAPI, CreateCardPaymentInput } from '../../core/card-payment/card-payment.interface'
import { CardPayment, PaymentType } from '../../core/card-payment/CardPayment'
import logger from '../../utils/logger'
import { CardPaymentClient } from '../grpc/pb/card-payment_grpc_pb'
import { Payment, PaymentInfoInput, PaymentTypeEnum, ProcessCardPaymentInput } from '../grpc/pb/card-payment_pb'

export default class GRPCCardPaymentAPI implements CardPaymentAPI {
  private readonly logger = logger.child({ label: GRPCCardPaymentAPI.name })

  async create ({ paymentDate, paymentInfo, paymentType, clientId, amount }: CreateCardPaymentInput): Promise<CardPayment> {
    return new Promise((resolve, reject) => {
      this.logger.info('Process started')

      try {
        this.logger.info(`Calling card-payment gRPC client at: ${ENV.CARD_PAYMENT_GRPC_HOST}`)
        const client = new CardPaymentClient(ENV.CARD_PAYMENT_GRPC_HOST, credentials.createInsecure())
        const request =
          GRPCCardPaymentAPI.buildProcessPaymentCardRequest(amount, clientId, paymentDate, paymentType, paymentInfo)

        client.proccessCardPayment(request, (error: ServiceError | null, res: Payment) => {
          if (error != null) {
            this.logger.error('Error while making gRPC request to create a card payment')
            this.logger.error(error)
            const fallbackMessage = 'An error occur while making a request to process the card payment'
            return reject(new Error(GRPCCardPaymentAPI.getGRPCErrorMessage(error, fallbackMessage)))
          }

          this.logger.info('Successfully called gRPC process payment')

          return resolve({
            id: res.getId(),
            clientId: res.getClientid(),
            amount: res.getAmount(),
            paymentDate: res.getPaymentdate(),
            paymentInfo: {
              maskedNumber: res.getPaymentinfo()?.getMaskednumber() ?? ''
            },
            paymentType: res.getPaymenttype() as PaymentType
          })
        })
      } catch (error) {
        this.logger.error('Error while building the process card payment gRPC request')
        this.logger.error(error)
        reject(new Error('Internal error while processing card payment request'))
      }
    })
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
