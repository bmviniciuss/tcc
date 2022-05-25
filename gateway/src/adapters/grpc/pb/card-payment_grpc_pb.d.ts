// package: cardpayment
// file: card-payment.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as card_payment_pb from "./card-payment_pb";

interface ICardPaymentService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    proccessCardPayment: ICardPaymentService_IProccessCardPayment;
}

interface ICardPaymentService_IProccessCardPayment extends grpc.MethodDefinition<card_payment_pb.ProcessCardPaymentInput, card_payment_pb.Payment> {
    path: "/cardpayment.CardPayment/ProccessCardPayment";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<card_payment_pb.ProcessCardPaymentInput>;
    requestDeserialize: grpc.deserialize<card_payment_pb.ProcessCardPaymentInput>;
    responseSerialize: grpc.serialize<card_payment_pb.Payment>;
    responseDeserialize: grpc.deserialize<card_payment_pb.Payment>;
}

export const CardPaymentService: ICardPaymentService;

export interface ICardPaymentServer {
    proccessCardPayment: grpc.handleUnaryCall<card_payment_pb.ProcessCardPaymentInput, card_payment_pb.Payment>;
}

export interface ICardPaymentClient {
    proccessCardPayment(request: card_payment_pb.ProcessCardPaymentInput, callback: (error: grpc.ServiceError | null, response: card_payment_pb.Payment) => void): grpc.ClientUnaryCall;
    proccessCardPayment(request: card_payment_pb.ProcessCardPaymentInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: card_payment_pb.Payment) => void): grpc.ClientUnaryCall;
    proccessCardPayment(request: card_payment_pb.ProcessCardPaymentInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: card_payment_pb.Payment) => void): grpc.ClientUnaryCall;
}

export class CardPaymentClient extends grpc.Client implements ICardPaymentClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public proccessCardPayment(request: card_payment_pb.ProcessCardPaymentInput, callback: (error: grpc.ServiceError | null, response: card_payment_pb.Payment) => void): grpc.ClientUnaryCall;
    public proccessCardPayment(request: card_payment_pb.ProcessCardPaymentInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: card_payment_pb.Payment) => void): grpc.ClientUnaryCall;
    public proccessCardPayment(request: card_payment_pb.ProcessCardPaymentInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: card_payment_pb.Payment) => void): grpc.ClientUnaryCall;
}
