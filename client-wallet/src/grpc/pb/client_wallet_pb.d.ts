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

export enum TransactionTypeEnum {
    CREDIT_CARD_PAYMENT = 0,
    DEBIT_CARD_PAYMENT = 1,
    WITHDRAWAL = 2,
}

export enum ServiceTypeEnum {
    CARD_PAYMENT = 0,
    INTERNAL = 1,
}
