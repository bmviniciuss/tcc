import axios from 'axios'

import { CardPayment, CardPaymentAPI, Payment } from '../../core/cardPayment/cardPayment'

export default class HttpCardPaymentAPI implements CardPaymentAPI {
  async process (cardPayment: CardPayment): Promise<Payment> {
    try {
      const { data } = await axios.post<Payment>('https://run.mocky.io/v3/c6c36610-b5a4-4833-97dc-b53a7083de59', cardPayment)
      return data
    } catch (error) {
      console.log(error)
      throw new Error('Error while processing card payment from api')
    }
  }
}
