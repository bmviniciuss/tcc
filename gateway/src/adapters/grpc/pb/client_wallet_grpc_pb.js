// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var client_wallet_pb = require('./client_wallet_pb.js');

function serialize_clientwallet_BalanceReturn(arg) {
  if (!(arg instanceof client_wallet_pb.BalanceReturn)) {
    throw new Error('Expected argument of type clientwallet.BalanceReturn');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_BalanceReturn(buffer_arg) {
  return client_wallet_pb.BalanceReturn.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_clientwallet_ClientTransactionsReturn(arg) {
  if (!(arg instanceof client_wallet_pb.ClientTransactionsReturn)) {
    throw new Error('Expected argument of type clientwallet.ClientTransactionsReturn');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_ClientTransactionsReturn(buffer_arg) {
  return client_wallet_pb.ClientTransactionsReturn.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_clientwallet_CreateTransactionInput(arg) {
  if (!(arg instanceof client_wallet_pb.CreateTransactionInput)) {
    throw new Error('Expected argument of type clientwallet.CreateTransactionInput');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_CreateTransactionInput(buffer_arg) {
  return client_wallet_pb.CreateTransactionInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_clientwallet_GetBalanceInput(arg) {
  if (!(arg instanceof client_wallet_pb.GetBalanceInput)) {
    throw new Error('Expected argument of type clientwallet.GetBalanceInput');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_GetBalanceInput(buffer_arg) {
  return client_wallet_pb.GetBalanceInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_clientwallet_GetClientTransactionsInput(arg) {
  if (!(arg instanceof client_wallet_pb.GetClientTransactionsInput)) {
    throw new Error('Expected argument of type clientwallet.GetClientTransactionsInput');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_GetClientTransactionsInput(buffer_arg) {
  return client_wallet_pb.GetClientTransactionsInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_clientwallet_Transaction(arg) {
  if (!(arg instanceof client_wallet_pb.Transaction)) {
    throw new Error('Expected argument of type clientwallet.Transaction');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_Transaction(buffer_arg) {
  return client_wallet_pb.Transaction.deserializeBinary(new Uint8Array(buffer_arg));
}


var ClientWalletService = exports.ClientWalletService = {
  createTransaction: {
    path: '/clientwallet.ClientWallet/CreateTransaction',
    requestStream: false,
    responseStream: false,
    requestType: client_wallet_pb.CreateTransactionInput,
    responseType: client_wallet_pb.Transaction,
    requestSerialize: serialize_clientwallet_CreateTransactionInput,
    requestDeserialize: deserialize_clientwallet_CreateTransactionInput,
    responseSerialize: serialize_clientwallet_Transaction,
    responseDeserialize: deserialize_clientwallet_Transaction,
  },
  getClientBalance: {
    path: '/clientwallet.ClientWallet/GetClientBalance',
    requestStream: false,
    responseStream: false,
    requestType: client_wallet_pb.GetBalanceInput,
    responseType: client_wallet_pb.BalanceReturn,
    requestSerialize: serialize_clientwallet_GetBalanceInput,
    requestDeserialize: deserialize_clientwallet_GetBalanceInput,
    responseSerialize: serialize_clientwallet_BalanceReturn,
    responseDeserialize: deserialize_clientwallet_BalanceReturn,
  },
  getClientTransactions: {
    path: '/clientwallet.ClientWallet/GetClientTransactions',
    requestStream: false,
    responseStream: false,
    requestType: client_wallet_pb.GetClientTransactionsInput,
    responseType: client_wallet_pb.ClientTransactionsReturn,
    requestSerialize: serialize_clientwallet_GetClientTransactionsInput,
    requestDeserialize: deserialize_clientwallet_GetClientTransactionsInput,
    responseSerialize: serialize_clientwallet_ClientTransactionsReturn,
    responseDeserialize: deserialize_clientwallet_ClientTransactionsReturn,
  },
};

exports.ClientWalletClient = grpc.makeGenericClientConstructor(ClientWalletService);
