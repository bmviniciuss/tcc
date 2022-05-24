import { credentials } from 'grpc'

import { ENV } from '../../application/config/env'
import { CardPaymentAPI, CreateCardPaymentInput } from '../../core/card-payment/card-payment.interface'
import { CardPayment, PaymentType } from '../../core/card-payment/CardPayment'
import logger from '../../utils/logger'
import { CardPaymentClient } from '../grpc/pb/card-payment_grpc_pb'
import { PaymentInfoInput, PaymentTypeEnum, ProcessCardPaymentInput } from '../grpc/pb/card-payment_pb'

export default class GRPCCardPaymentAPI implements CardPaymentAPI {
  private readonly logger = logger.child({ label: GRPCCardPaymentAPI.name })

  async create ({ paymentDate, paymentInfo, paymentType, clientId, amount }: CreateCardPaymentInput): Promise<CardPayment> {
    return new Promise((resolve, reject) => {
      this.logger.info('Process started')

      try {
        this.logger.info(`Calling card-payment gRPC client at: ${ENV.CARD_PAYMENT_GRPC_HOST}`)
        const client = new CardPaymentClient(ENV.CARD_PAYMENT_GRPC_HOST, credentials.createInsecure())
        const request =
          new ProcessCardPaymentInput()
            .setAmount(amount)
            .setClientid(clientId)
            .setPaymentdate(paymentDate)
            .setPaymenttype(PaymentTypeEnum[paymentType])
            .setPaymentinfo(
              new PaymentInfoInput()
                .setCardtoken(paymentInfo.cardToken)
            )
        client.proccessCardPayment(request, (error, res) => {
          if (error != null) {
            this.logger.error('Error in gRPC method')
            this.logger.error(error?.message)
            throw error
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
        this.logger.error('Error to process card payment gRPC request')
        this.logger.error(error)
        let message = 'Internal error while processing card payment request'
        if (error instanceof Error) message = error?.message
        reject(message)
      }
    })
  }
}
