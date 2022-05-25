// package: cards
// file: cards.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as cards_pb from "./cards_pb";

interface ICardsService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    generateCard: ICardsService_IGenerateCard;
    getCardByToken: ICardsService_IGetCardByToken;
}

interface ICardsService_IGenerateCard extends grpc.MethodDefinition<cards_pb.CreateCardRequest, cards_pb.FullCard> {
    path: "/cards.Cards/GenerateCard";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cards_pb.CreateCardRequest>;
    requestDeserialize: grpc.deserialize<cards_pb.CreateCardRequest>;
    responseSerialize: grpc.serialize<cards_pb.FullCard>;
    responseDeserialize: grpc.deserialize<cards_pb.FullCard>;
}
interface ICardsService_IGetCardByToken extends grpc.MethodDefinition<cards_pb.GetCardByTokenRequest, cards_pb.Card> {
    path: "/cards.Cards/GetCardByToken";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cards_pb.GetCardByTokenRequest>;
    requestDeserialize: grpc.deserialize<cards_pb.GetCardByTokenRequest>;
    responseSerialize: grpc.serialize<cards_pb.Card>;
    responseDeserialize: grpc.deserialize<cards_pb.Card>;
}

export const CardsService: ICardsService;

export interface ICardsServer {
    generateCard: grpc.handleUnaryCall<cards_pb.CreateCardRequest, cards_pb.FullCard>;
    getCardByToken: grpc.handleUnaryCall<cards_pb.GetCardByTokenRequest, cards_pb.Card>;
}

export interface ICardsClient {
    generateCard(request: cards_pb.CreateCardRequest, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    getCardByToken(request: cards_pb.GetCardByTokenRequest, callback: (error: grpc.ServiceError | null, response: cards_pb.Card) => void): grpc.ClientUnaryCall;
    getCardByToken(request: cards_pb.GetCardByTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cards_pb.Card) => void): grpc.ClientUnaryCall;
    getCardByToken(request: cards_pb.GetCardByTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cards_pb.Card) => void): grpc.ClientUnaryCall;
}

export class CardsClient extends grpc.Client implements ICardsClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public generateCard(request: cards_pb.CreateCardRequest, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    public generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    public generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    public getCardByToken(request: cards_pb.GetCardByTokenRequest, callback: (error: grpc.ServiceError | null, response: cards_pb.Card) => void): grpc.ClientUnaryCall;
    public getCardByToken(request: cards_pb.GetCardByTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cards_pb.Card) => void): grpc.ClientUnaryCall;
    public getCardByToken(request: cards_pb.GetCardByTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cards_pb.Card) => void): grpc.ClientUnaryCall;
}
