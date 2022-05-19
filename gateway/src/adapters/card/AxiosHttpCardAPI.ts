import axios from 'axios'

import { Card } from '../../core/card/Card'
import { CardAPI, CreateCardInput } from '../../core/card/card.interface'

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
  async create ({ isDebit, isCredit, cardholderName }: CreateCardInput): Promise<Card> {
    try {
      const request: CreateCardRequest = {
        cardholder_name: cardholderName,
        is_credit: isCredit,
        is_debit: isDebit
      }
      const URL = 'http://localhost:3333/api/cards'
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
      console.error('Error in create card http request', error)
      throw new Error('HttpCreateCardError') // TODO: Better error handling
    }
  }
}
