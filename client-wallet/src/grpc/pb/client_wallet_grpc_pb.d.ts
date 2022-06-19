// package: clientwallet
// file: client_wallet.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as client_wallet_pb from "./client_wallet_pb";

interface IClientWalletService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createTransaction: IClientWalletService_ICreateTransaction;
    getClientBalance: IClientWalletService_IGetClientBalance;
}

interface IClientWalletService_ICreateTransaction extends grpc.MethodDefinition<client_wallet_pb.CreateTransactionInput, client_wallet_pb.Transaction> {
    path: "/clientwallet.ClientWallet/CreateTransaction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<client_wallet_pb.CreateTransactionInput>;
    requestDeserialize: grpc.deserialize<client_wallet_pb.CreateTransactionInput>;
    responseSerialize: grpc.serialize<client_wallet_pb.Transaction>;
    responseDeserialize: grpc.deserialize<client_wallet_pb.Transaction>;
}
interface IClientWalletService_IGetClientBalance extends grpc.MethodDefinition<client_wallet_pb.GetBalanceInput, client_wallet_pb.BalanceReturn> {
    path: "/clientwallet.ClientWallet/GetClientBalance";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<client_wallet_pb.GetBalanceInput>;
    requestDeserialize: grpc.deserialize<client_wallet_pb.GetBalanceInput>;
    responseSerialize: grpc.serialize<client_wallet_pb.BalanceReturn>;
    responseDeserialize: grpc.deserialize<client_wallet_pb.BalanceReturn>;
}

export const ClientWalletService: IClientWalletService;

export interface IClientWalletServer {
    createTransaction: grpc.handleUnaryCall<client_wallet_pb.CreateTransactionInput, client_wallet_pb.Transaction>;
    getClientBalance: grpc.handleUnaryCall<client_wallet_pb.GetBalanceInput, client_wallet_pb.BalanceReturn>;
}

export interface IClientWalletClient {
    createTransaction(request: client_wallet_pb.CreateTransactionInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    getClientBalance(request: client_wallet_pb.GetBalanceInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
}

export class ClientWalletClient extends grpc.Client implements IClientWalletClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createTransaction(request: client_wallet_pb.CreateTransactionInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    public createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    public createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    public getClientBalance(request: client_wallet_pb.GetBalanceInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    public getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    public getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
}
