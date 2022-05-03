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

export type CardPaymentResponse = {
  id: string
  clientId: string
  amount: PaymentAmount
  paymentType: PaymentType
}

export type CardPaymentAPIRequest = {
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

export type CardPaymentAPIResponse = {
  id: string
  clientId: string
  amount: PaymentAmount
  paymentType: PaymentType
}

export interface CardProcessor {
  processCardPayment(cardPayment: CardPayment): Promise<CardPaymentResponse>
}

export interface CardPaymentAPI {
  process(cardPayment: CardPaymentAPIRequest): Promise<CardPaymentAPIResponse>
}

export default class CardPaymentProcessor implements CardProcessor {
  constructor (private readonly cardPaymentAPI: CardPaymentAPI) {}

  processCardPayment (cardPayment: CardPayment): Promise<CardPaymentResponse> {
    return this.cardPaymentAPI.process(cardPayment)
  }
}
