// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var card$payment_pb = require('./card-payment_pb.js');

function serialize_cardpayment_Payment(arg) {
  if (!(arg instanceof card$payment_pb.Payment)) {
    throw new Error('Expected argument of type cardpayment.Payment');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_cardpayment_Payment(buffer_arg) {
  return card$payment_pb.Payment.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_cardpayment_ProcessCardPaymentInput(arg) {
  if (!(arg instanceof card$payment_pb.ProcessCardPaymentInput)) {
    throw new Error('Expected argument of type cardpayment.ProcessCardPaymentInput');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_cardpayment_ProcessCardPaymentInput(buffer_arg) {
  return card$payment_pb.ProcessCardPaymentInput.deserializeBinary(new Uint8Array(buffer_arg));
}


var CardPaymentService = exports.CardPaymentService = {
  proccessCardPayment: {
    path: '/cardpayment.CardPayment/ProccessCardPayment',
    requestStream: false,
    responseStream: false,
    requestType: card$payment_pb.ProcessCardPaymentInput,
    responseType: card$payment_pb.Payment,
    requestSerialize: serialize_cardpayment_ProcessCardPaymentInput,
    requestDeserialize: deserialize_cardpayment_ProcessCardPaymentInput,
    responseSerialize: serialize_cardpayment_Payment,
    responseDeserialize: deserialize_cardpayment_Payment,
  },
};

exports.CardPaymentClient = grpc.makeGenericClientConstructor(CardPaymentService);
