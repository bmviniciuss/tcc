import grpc from '@grpc/grpc-js'

import { Transaction as CoreTransaction } from '../../core/transaction/Transaction'
import {
  CreateTransactionInput as CoreCreateTransactionInput,
  ITransactionService
} from '../../core/transaction/transaction.interfaces'
import logger from '../../utils/logger'
import { IClientWalletServer } from '../pb/client_wallet_grpc_pb'
import { CreateTransactionInput, Transaction } from '../pb/client_wallet_pb'
import ServiceTypeGRPCParser from './utils/ServiceTypeGRPCParser'
import TransactionTypeGRPCParser from './utils/TransactionTypeGRPCParser'

export class ClientWalletServiceImpl implements IClientWalletServer {
  private readonly l = logger.child({ label: ClientWalletServiceImpl.name })

  constructor (private readonly transactionService: ITransactionService) {}

  createTransaction (call: grpc.ServerUnaryCall<CreateTransactionInput, Transaction>, callback: grpc.sendUnaryData<Transaction>): void {
    this.l.info('[gRPC] Received transaction request')
    const createTransactionInput = call.request.toObject()

    const transactionDate = call.request.getTransactionDate()

    if (!transactionDate) {
      this.l.error('[gRPC] Transaction date is required')
      callback(new Error('Transaction date is required'), null)
      return
    }

    const serviceInput = ClientWalletServiceImpl.buildCreateTransactionServiceInput(createTransactionInput, transactionDate)

    this.transactionService.create(serviceInput).then(transaction => {
      this.l.info('[gRPC] Transaction created')

      const transactionResponse = ClientWalletServiceImpl.buildTransactionResponse(transaction)

      callback(null, transactionResponse)
    }).catch(error => {
      this.l.error('[gRPC] Error creating transaction', error)
      callback(error, null)
    })
  }

  private static buildTransactionResponse (transaction: CoreTransaction) {
    return new Transaction()
      .setClientId(transaction.clientId)
      .setAmount(transaction.amount)
      .setType(TransactionTypeGRPCParser.toGRPC(transaction.type))
      .setService(ServiceTypeGRPCParser.toGRPC(transaction.service))
      .setTransactionServiceId(transaction.transactionServiceId ?? '')
      .setTransactionDate(transaction.transactionDate.toISOString())
      .setCreatedAt(transaction.createdAt.toISOString())
  }

  private static buildCreateTransactionServiceInput (createTransactionInput: CreateTransactionInput.AsObject, transactionDate: string): CoreCreateTransactionInput {
    return {
      clientId: createTransactionInput.clientId,
      amount: createTransactionInput.amount,
      type: TransactionTypeGRPCParser.toDomain(createTransactionInput.type),
      service: ServiceTypeGRPCParser.toDomain(createTransactionInput.service),
      transactionServiceId: createTransactionInput.transactionServiceId,
      transactionDate: new Date(transactionDate)
    }
  }
}
