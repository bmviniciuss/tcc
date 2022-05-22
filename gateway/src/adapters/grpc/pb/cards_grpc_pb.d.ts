// package: cards
// file: cards.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as cards_pb from "./cards_pb";

interface ICardsService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    generateCard: ICardsService_IGenerateCard;
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

export const CardsService: ICardsService;

export interface ICardsServer {
    generateCard: grpc.handleUnaryCall<cards_pb.CreateCardRequest, cards_pb.FullCard>;
}

export interface ICardsClient {
    generateCard(request: cards_pb.CreateCardRequest, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
}

export class CardsClient extends grpc.Client implements ICardsClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public generateCard(request: cards_pb.CreateCardRequest, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    public generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
    public generateCard(request: cards_pb.CreateCardRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cards_pb.FullCard) => void): grpc.ClientUnaryCall;
}
