// package: clientwallet
// file: client_wallet.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as client_wallet_pb from "./client_wallet_pb";

interface IClientWalletService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createTransaction: IClientWalletService_ICreateTransaction;
    getClientBalance: IClientWalletService_IGetClientBalance;
    getClientTransactions: IClientWalletService_IGetClientTransactions;
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
interface IClientWalletService_IGetClientTransactions extends grpc.MethodDefinition<client_wallet_pb.GetClientTransactionsInput, client_wallet_pb.ClientTransactionsReturn> {
    path: "/clientwallet.ClientWallet/GetClientTransactions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<client_wallet_pb.GetClientTransactionsInput>;
    requestDeserialize: grpc.deserialize<client_wallet_pb.GetClientTransactionsInput>;
    responseSerialize: grpc.serialize<client_wallet_pb.ClientTransactionsReturn>;
    responseDeserialize: grpc.deserialize<client_wallet_pb.ClientTransactionsReturn>;
}

export const ClientWalletService: IClientWalletService;

export interface IClientWalletServer {
    createTransaction: grpc.handleUnaryCall<client_wallet_pb.CreateTransactionInput, client_wallet_pb.Transaction>;
    getClientBalance: grpc.handleUnaryCall<client_wallet_pb.GetBalanceInput, client_wallet_pb.BalanceReturn>;
    getClientTransactions: grpc.handleUnaryCall<client_wallet_pb.GetClientTransactionsInput, client_wallet_pb.ClientTransactionsReturn>;
}

export interface IClientWalletClient {
    createTransaction(request: client_wallet_pb.CreateTransactionInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    getClientBalance(request: client_wallet_pb.GetBalanceInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    getClientTransactions(request: client_wallet_pb.GetClientTransactionsInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.ClientTransactionsReturn) => void): grpc.ClientUnaryCall;
    getClientTransactions(request: client_wallet_pb.GetClientTransactionsInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.ClientTransactionsReturn) => void): grpc.ClientUnaryCall;
    getClientTransactions(request: client_wallet_pb.GetClientTransactionsInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.ClientTransactionsReturn) => void): grpc.ClientUnaryCall;
}

export class ClientWalletClient extends grpc.Client implements IClientWalletClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createTransaction(request: client_wallet_pb.CreateTransactionInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    public createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    public createTransaction(request: client_wallet_pb.CreateTransactionInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.Transaction) => void): grpc.ClientUnaryCall;
    public getClientBalance(request: client_wallet_pb.GetBalanceInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    public getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    public getClientBalance(request: client_wallet_pb.GetBalanceInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.BalanceReturn) => void): grpc.ClientUnaryCall;
    public getClientTransactions(request: client_wallet_pb.GetClientTransactionsInput, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.ClientTransactionsReturn) => void): grpc.ClientUnaryCall;
    public getClientTransactions(request: client_wallet_pb.GetClientTransactionsInput, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.ClientTransactionsReturn) => void): grpc.ClientUnaryCall;
    public getClientTransactions(request: client_wallet_pb.GetClientTransactionsInput, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: client_wallet_pb.ClientTransactionsReturn) => void): grpc.ClientUnaryCall;
}
