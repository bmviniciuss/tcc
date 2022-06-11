import { CreateTransactionInput, ITransactionRepository, ITransactionService } from './transaction.interfaces'
import { Transaction } from './Transaction'
import logger from '../../utils/logger'

export default class TransactionService implements ITransactionService {
  private readonly logger = logger.child({ label: TransactionService.name })
  constructor (private readonly transactionRepository: ITransactionRepository) {}

  async create (transaction: CreateTransactionInput): Promise<Transaction> {
    this.logger.info('Creating transaction')
    return this.transactionRepository.create(transaction)
  }

  listByClientId (clientId: string): Promise<Transaction[]> {
    this.logger.info('Listing transactions by clientId')
    return this.transactionRepository.getByClientId(clientId)
  }

  getBalanceByClientId (clientId: string): Promise<number> {
    this.logger.info('Getting balance by clientId')
    return this.transactionRepository.getBalanceByClientId(clientId)
  }
}
