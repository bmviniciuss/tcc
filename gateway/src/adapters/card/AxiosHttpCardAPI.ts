import axios from 'axios'

import { ENV } from '../../application/config/env'
import { Card } from '../../core/card/Card'
import { CardAPI, CreateCardInput } from '../../core/card/card.interface'
import logger from '../../utils/logger'

type CreateCardRequest = {
  cardholder_name: string
  is_debit: boolean
  is_credit: boolean
}

type CreateCardResponse = {
  id: string
  cardholder_name: string
  token: string
  masked_number: string
  expiration_year: number
  expiration_month: number
  active: boolean
  is_credit: boolean
  is_debit: boolean
}

export default class AxiosHttpCardAPI implements CardAPI {
  private readonly logger = logger.child({ label: AxiosHttpCardAPI.name })

  async create ({ isDebit, isCredit, cardholderName }: CreateCardInput): Promise<Card> {
    try {
      const request: CreateCardRequest = {
        cardholder_name: cardholderName,
        is_credit: isCredit,
        is_debit: isDebit
      }
      const base = `http://${ENV.CARD_HOST}`
      const URL = `${base}/api/cards`
      const { data } = await axios.post<CreateCardResponse>(URL, request)
      return {
        id: data.id,
        cardholderName: data.cardholder_name,
        token: data.token,
        expirationYear: data.expiration_year,
        expirationMonth: data.expiration_month,
        maskedNumber: data.masked_number,
        active: data.active,
        isCredit: data.is_credit,
        isDebit: data.is_debit
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
