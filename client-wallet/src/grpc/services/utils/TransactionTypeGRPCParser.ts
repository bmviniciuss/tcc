import { TransactionType } from '../../../core/transaction/Transaction'
import { TransactionTypeEnum } from '../../pb/client_wallet_pb'

export default class TransactionTypeGRPCParser {
  public static toDomain (type: TransactionTypeEnum): TransactionType {
    switch (type) {
      case TransactionTypeEnum.CREDIT_CARD_PAYMENT:
        return TransactionType.CREDIT_CARD_PAYMENT
      case TransactionTypeEnum.DEBIT_CARD_PAYMENT:
        return TransactionType.DEBIT_CARD_PAYMENT
      case TransactionTypeEnum.WITHDRAWAL:
        return TransactionType.WITHDRAWAL
      default:
        throw new Error(`Unknown type: ${type}`)
    }
  }

  public static toGRPC (type: TransactionType): TransactionTypeEnum {
    switch (type) {
      case TransactionType.CREDIT_CARD_PAYMENT:
        return TransactionTypeEnum.CREDIT_CARD_PAYMENT
      case TransactionType.DEBIT_CARD_PAYMENT:
        return TransactionTypeEnum.DEBIT_CARD_PAYMENT
      case TransactionType.WITHDRAWAL:
        return TransactionTypeEnum.WITHDRAWAL
      default:
        throw new Error(`Unknown type: ${type}`)
    }
  }
}
