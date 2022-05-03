export type PaymentAmount = {
  currency: string
  value: number
}

export enum PaymentType {
  CREDIT_CARD = 'CREDIT_CARD',
  DEBIT_CARD = 'CREDIT_CARD',
  PIX = 'PIX'
}
