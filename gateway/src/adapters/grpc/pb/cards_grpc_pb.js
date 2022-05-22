// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var cards_pb = require('./cards_pb.js');

function serialize_cards_CreateCardRequest(arg) {
  if (!(arg instanceof cards_pb.CreateCardRequest)) {
    throw new Error('Expected argument of type cards.CreateCardRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_cards_CreateCardRequest(buffer_arg) {
  return cards_pb.CreateCardRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_cards_FullCard(arg) {
  if (!(arg instanceof cards_pb.FullCard)) {
    throw new Error('Expected argument of type cards.FullCard');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_cards_FullCard(buffer_arg) {
  return cards_pb.FullCard.deserializeBinary(new Uint8Array(buffer_arg));
}


var CardsService = exports.CardsService = {
  generateCard: {
    path: '/cards.Cards/GenerateCard',
    requestStream: false,
    responseStream: false,
    requestType: cards_pb.CreateCardRequest,
    responseType: cards_pb.FullCard,
    requestSerialize: serialize_cards_CreateCardRequest,
    requestDeserialize: deserialize_cards_CreateCardRequest,
    responseSerialize: serialize_cards_FullCard,
    responseDeserialize: deserialize_cards_FullCard,
  },
};

exports.CardsClient = grpc.makeGenericClientConstructor(CardsService);
