// package: clientwallet
// file: client_wallet.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Transaction extends jspb.Message { 
    getId(): string;
    setId(value: string): Transaction;
    getClientId(): string;
    setClientId(value: string): Transaction;
    getAmount(): number;
    setAmount(value: number): Transaction;
    getType(): TransactionTypeEnum;
    setType(value: TransactionTypeEnum): Transaction;
    getTransactionServiceId(): string;
    setTransactionServiceId(value: string): Transaction;
    getService(): ServiceTypeEnum;
    setService(value: ServiceTypeEnum): Transaction;
    getTransactionDate(): string;
    setTransactionDate(value: string): Transaction;
    getCreatedAt(): string;
    setCreatedAt(value: string): Transaction;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Transaction.AsObject;
    static toObject(includeInstance: boolean, msg: Transaction): Transaction.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Transaction, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Transaction;
    static deserializeBinaryFromReader(message: Transaction, reader: jspb.BinaryReader): Transaction;
}

export namespace Transaction {
    export type AsObject = {
        id: string,
        clientId: string,
        amount: number,
        type: TransactionTypeEnum,
        transactionServiceId: string,
        service: ServiceTypeEnum,
        transactionDate: string,
        createdAt: string,
    }
}

export class CreateTransactionInput extends jspb.Message { 
    getClientId(): string;
    setClientId(value: string): CreateTransactionInput;
    getAmount(): number;
    setAmount(value: number): CreateTransactionInput;
    getType(): TransactionTypeEnum;
    setType(value: TransactionTypeEnum): CreateTransactionInput;
    getService(): ServiceTypeEnum;
    setService(value: ServiceTypeEnum): CreateTransactionInput;
    getTransactionServiceId(): string;
    setTransactionServiceId(value: string): CreateTransactionInput;
    getTransactionDate(): string;
    setTransactionDate(value: string): CreateTransactionInput;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateTransactionInput.AsObject;
    static toObject(includeInstance: boolean, msg: CreateTransactionInput): CreateTransactionInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateTransactionInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateTransactionInput;
    static deserializeBinaryFromReader(message: CreateTransactionInput, reader: jspb.BinaryReader): CreateTransactionInput;
}

export namespace CreateTransactionInput {
    export type AsObject = {
        clientId: string,
        amount: number,
        type: TransactionTypeEnum,
        service: ServiceTypeEnum,
        transactionServiceId: string,
        transactionDate: string,
    }
}

export class GetBalanceInput extends jspb.Message { 
    getClientId(): string;
    setClientId(value: string): GetBalanceInput;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetBalanceInput.AsObject;
    static toObject(includeInstance: boolean, msg: GetBalanceInput): GetBalanceInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetBalanceInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetBalanceInput;
    static deserializeBinaryFromReader(message: GetBalanceInput, reader: jspb.BinaryReader): GetBalanceInput;
}

export namespace GetBalanceInput {
    export type AsObject = {
        clientId: string,
    }
}

export class BalanceReturn extends jspb.Message { 
    getBalance(): number;
    setBalance(value: number): BalanceReturn;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BalanceReturn.AsObject;
    static toObject(includeInstance: boolean, msg: BalanceReturn): BalanceReturn.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BalanceReturn, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BalanceReturn;
    static deserializeBinaryFromReader(message: BalanceReturn, reader: jspb.BinaryReader): BalanceReturn;
}

export namespace BalanceReturn {
    export type AsObject = {
        balance: number,
    }
}

export class GetClientTransactionsInput extends jspb.Message { 
    getClientId(): string;
    setClientId(value: string): GetClientTransactionsInput;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetClientTransactionsInput.AsObject;
    static toObject(includeInstance: boolean, msg: GetClientTransactionsInput): GetClientTransactionsInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetClientTransactionsInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetClientTransactionsInput;
    static deserializeBinaryFromReader(message: GetClientTransactionsInput, reader: jspb.BinaryReader): GetClientTransactionsInput;
}

export namespace GetClientTransactionsInput {
    export type AsObject = {
        clientId: string,
    }
}

export class ClientTransactionsReturn extends jspb.Message { 
    clearTransactionsList(): void;
    getTransactionsList(): Array<Transaction>;
    setTransactionsList(value: Array<Transaction>): ClientTransactionsReturn;
    addTransactions(value?: Transaction, index?: number): Transaction;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClientTransactionsReturn.AsObject;
    static toObject(includeInstance: boolean, msg: ClientTransactionsReturn): ClientTransactionsReturn.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClientTransactionsReturn, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClientTransactionsReturn;
    static deserializeBinaryFromReader(message: ClientTransactionsReturn, reader: jspb.BinaryReader): ClientTransactionsReturn;
}

export namespace ClientTransactionsReturn {
    export type AsObject = {
        transactionsList: Array<Transaction.AsObject>,
    }
}

export enum TransactionTypeEnum {
    CREDIT_CARD_PAYMENT = 0,
    DEBIT_CARD_PAYMENT = 1,
    WITHDRAWAL = 2,
}

export enum ServiceTypeEnum {
    CARD_PAYMENT = 0,
    INTERNAL = 1,
}
