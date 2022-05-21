import { CardPayment, PaymentType } from './CardPayment'

export type CreateCardPaymentInput = {
  clientId: string
  paymentType: PaymentType
  paymentDate: string
  amount: number
  paymentInfo: {
    cardToken: string
  }
}

export interface ICardPaymentService {
  create(input: CreateCardPaymentInput): Promise<CardPayment>
}

export interface CardPaymentAPI {
  create(input: CreateCardPaymentInput): Promise<CardPayment>
}
