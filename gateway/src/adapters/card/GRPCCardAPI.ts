import { credentials } from 'grpc'

import { ENV } from '../../application/config/env'
import { Card } from '../../core/card/Card'
import { CardAPI, CreateCardInput } from '../../core/card/card.interface'
import logger from '../../utils/logger'
import { CardsClient } from '../grpc/pb/cards_grpc_pb'
import { CreateCardRequest } from '../grpc/pb/cards_pb'

export default class GRPCCardApi implements CardAPI {
  private readonly logger = logger.child({ label: GRPCCardApi.name })

  async create ({ isDebit, isCredit, cardholderName }: CreateCardInput): Promise<Card> {
    this.logger.info('Creating card through GRPC client')
    return new Promise((resolve, reject) => {
      const client = new CardsClient(ENV.CARD_GRPC_HOST, credentials.createInsecure())
      const request = new CreateCardRequest()
      request.setCardholderName(cardholderName)
      request.setIscredit(isCredit)
      request.setIsdebit(isDebit)

      client.generateCard(request, (error, res) => {
        if (error != null) {
          this.logger.error('Error in create card grpc request')
          this.logger.error(error)
          let message = 'Internal error while generating card'
          if (error instanceof Error) message = error?.message
          reject(new Error(message)) // TODO: Better error handling
        }
        return resolve({
          id: res.getId(),
          cardholderName: res.getCardholdername(),
          token: res.getToken(),
          expirationYear: res.getExpirationyear(),
          expirationMonth: res.getExpirationmonth(),
          maskedNumber: res.getMaskednumber(),
          active: res.getActive(),
          isCredit: res.getIscredit(),
          isDebit: res.getIsdebit()

        })
      })
    })
  }
}
