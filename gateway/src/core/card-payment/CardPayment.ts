
export enum PaymentType {
  CREDIT_CARD ='CREDIT_CARD',
  DEBIT_CARD = 'DEBIT_CARD'
}

export type CardPayment = {
  clientId: string
  paymentType: PaymentType
  paymentDate: string
  amount: number
  paymentInfo: {
    maskedNumber: string
  }
}
