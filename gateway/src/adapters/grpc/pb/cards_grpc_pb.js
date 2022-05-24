// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var cards_pb = require('./cards_pb.js');

function serialize_cards_Card(arg) {
  if (!(arg instanceof cards_pb.Card)) {
    throw new Error('Expected argument of type cards.Card');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_cards_Card(buffer_arg) {
  return cards_pb.Card.deserializeBinary(new Uint8Array(buffer_arg));
}

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

function serialize_cards_GetCardByTokenRequest(arg) {
  if (!(arg instanceof cards_pb.GetCardByTokenRequest)) {
    throw new Error('Expected argument of type cards.GetCardByTokenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_cards_GetCardByTokenRequest(buffer_arg) {
  return cards_pb.GetCardByTokenRequest.deserializeBinary(new Uint8Array(buffer_arg));
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
  getCardByToken: {
    path: '/cards.Cards/GetCardByToken',
    requestStream: false,
    responseStream: false,
    requestType: cards_pb.GetCardByTokenRequest,
    responseType: cards_pb.Card,
    requestSerialize: serialize_cards_GetCardByTokenRequest,
    requestDeserialize: deserialize_cards_GetCardByTokenRequest,
    responseSerialize: serialize_cards_Card,
    responseDeserialize: deserialize_cards_Card,
  },
};

exports.CardsClient = grpc.makeGenericClientConstructor(CardsService);
