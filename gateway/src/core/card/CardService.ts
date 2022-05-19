import { Card } from './Card'
import { CardAPI, CreateCardInput, ICardService } from './card.interface'

export default class CardService implements ICardService {
  constructor (private readonly cardAPI: CardAPI) {}

  async create (input: CreateCardInput): Promise<Card> {
    return await this.cardAPI.create(input)
  }
}
