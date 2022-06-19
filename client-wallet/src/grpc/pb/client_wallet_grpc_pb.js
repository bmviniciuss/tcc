// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var client_wallet_pb = require('./client_wallet_pb.js');

function serialize_clientwallet_CreateTransactionInput(arg) {
  if (!(arg instanceof client_wallet_pb.CreateTransactionInput)) {
    throw new Error('Expected argument of type clientwallet.CreateTransactionInput');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_clientwallet_CreateTransactionInput(buffer_arg) {
  return client_wallet_pb.CreateTransactionInput.deserializeBinary(new Uint8Array(buffer_arg));
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
};

exports.ClientWalletClient = grpc.makeGenericClientConstructor(ClientWalletService);
