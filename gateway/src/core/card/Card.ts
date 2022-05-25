export type Card = {
  id: string
  cardholderName: string
  token: string
  maskedNumber: string
  expirationYear: number
  expirationMonth: number
  active: boolean
  isCredit: boolean
  isDebit: boolean
}
