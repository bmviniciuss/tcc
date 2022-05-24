// package: cardpayment
// file: card-payment.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class ProcessCardPaymentInput extends jspb.Message { 
    getClientid(): string;
    setClientid(value: string): ProcessCardPaymentInput;
    getPaymenttype(): PaymentTypeEnum;
    setPaymenttype(value: PaymentTypeEnum): ProcessCardPaymentInput;
    getPaymentdate(): string;
    setPaymentdate(value: string): ProcessCardPaymentInput;
    getAmount(): number;
    setAmount(value: number): ProcessCardPaymentInput;

    hasPaymentinfo(): boolean;
    clearPaymentinfo(): void;
    getPaymentinfo(): PaymentInfoInput | undefined;
    setPaymentinfo(value?: PaymentInfoInput): ProcessCardPaymentInput;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProcessCardPaymentInput.AsObject;
    static toObject(includeInstance: boolean, msg: ProcessCardPaymentInput): ProcessCardPaymentInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProcessCardPaymentInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProcessCardPaymentInput;
    static deserializeBinaryFromReader(message: ProcessCardPaymentInput, reader: jspb.BinaryReader): ProcessCardPaymentInput;
}

export namespace ProcessCardPaymentInput {
    export type AsObject = {
        clientid: string,
        paymenttype: PaymentTypeEnum,
        paymentdate: string,
        amount: number,
        paymentinfo?: PaymentInfoInput.AsObject,
    }
}

export class PaymentInfoInput extends jspb.Message { 
    getCardtoken(): string;
    setCardtoken(value: string): PaymentInfoInput;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PaymentInfoInput.AsObject;
    static toObject(includeInstance: boolean, msg: PaymentInfoInput): PaymentInfoInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PaymentInfoInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PaymentInfoInput;
    static deserializeBinaryFromReader(message: PaymentInfoInput, reader: jspb.BinaryReader): PaymentInfoInput;
}

export namespace PaymentInfoInput {
    export type AsObject = {
        cardtoken: string,
    }
}

export class Payment extends jspb.Message { 
    getId(): string;
    setId(value: string): Payment;
    getClientid(): string;
    setClientid(value: string): Payment;
    getPaymenttype(): string;
    setPaymenttype(value: string): Payment;
    getPaymentdate(): string;
    setPaymentdate(value: string): Payment;
    getAmount(): number;
    setAmount(value: number): Payment;

    hasPaymentinfo(): boolean;
    clearPaymentinfo(): void;
    getPaymentinfo(): PaymentInfo | undefined;
    setPaymentinfo(value?: PaymentInfo): Payment;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Payment.AsObject;
    static toObject(includeInstance: boolean, msg: Payment): Payment.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Payment, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Payment;
    static deserializeBinaryFromReader(message: Payment, reader: jspb.BinaryReader): Payment;
}

export namespace Payment {
    export type AsObject = {
        id: string,
        clientid: string,
        paymenttype: string,
        paymentdate: string,
        amount: number,
        paymentinfo?: PaymentInfo.AsObject,
    }
}

export class PaymentInfo extends jspb.Message { 
    getMaskednumber(): string;
    setMaskednumber(value: string): PaymentInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PaymentInfo.AsObject;
    static toObject(includeInstance: boolean, msg: PaymentInfo): PaymentInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PaymentInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PaymentInfo;
    static deserializeBinaryFromReader(message: PaymentInfo, reader: jspb.BinaryReader): PaymentInfo;
}

export namespace PaymentInfo {
    export type AsObject = {
        maskednumber: string,
    }
}

export enum PaymentTypeEnum {
    CREDIT_CARD = 0,
    DEBIT_CARD = 1,
}
