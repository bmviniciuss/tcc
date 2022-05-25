import { Card } from './Card'

export type CreateCardInput = {
  cardholderName: string
  isDebit: boolean
  isCredit: boolean
}

export interface ICardService {
  create(input: CreateCardInput): Promise<Card>
}

export interface CardAPI {
  create(input: CreateCardInput): Promise<Card>
}
