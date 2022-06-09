import { CardPaymentAPI, CreateCardPaymentInput, ICardPaymentService } from './card-payment.interface'
import { CardPayment } from './CardPayment'

export default class CardPaymentService implements ICardPaymentService {
  constructor (private readonly cardPaymentAPI: CardPaymentAPI) {}

  async create (input: CreateCardPaymentInput): Promise<CardPayment> {
    return this.cardPaymentAPI.create(input)
  }
}
