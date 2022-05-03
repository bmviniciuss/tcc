import { PaymentAmount, PaymentType } from '../common/models'

export type CardPayment = {
  id: string
  clientId: string
  amount: PaymentAmount
  paymentType: PaymentType
  cardPaymentInfo: {
    number: string
    cvv: string
    expirationMonth: string
    expirationYear: string
    cardholderName: string
  }
}

export type Payment = {
  id: string
  clientId: string
  amount: PaymentAmount
  paymentType: PaymentType
}

export interface CardProcessor {
  processCardPayment(cardPayment: CardPayment): Promise<Payment>
}

export interface CardPaymentAPI {
  process(cardPayment: CardPayment): Promise<Payment>
}

export default class CardPaymentProcessor implements CardProcessor {
  constructor (private readonly cardPaymentAPI: CardPaymentAPI) {}

  async processCardPayment (cardPayment: CardPayment): Promise<Payment> {
    return this.cardPaymentAPI.process(cardPayment)
  }
}
