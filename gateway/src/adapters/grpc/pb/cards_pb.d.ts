// package: cards
// file: cards.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class CreateCardRequest extends jspb.Message { 
    getCardholderName(): string;
    setCardholderName(value: string): CreateCardRequest;
    getIscredit(): boolean;
    setIscredit(value: boolean): CreateCardRequest;
    getIsdebit(): boolean;
    setIsdebit(value: boolean): CreateCardRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateCardRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateCardRequest): CreateCardRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateCardRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateCardRequest;
    static deserializeBinaryFromReader(message: CreateCardRequest, reader: jspb.BinaryReader): CreateCardRequest;
}

export namespace CreateCardRequest {
    export type AsObject = {
        cardholderName: string,
        iscredit: boolean,
        isdebit: boolean,
    }
}

export class FullCard extends jspb.Message { 
    getId(): string;
    setId(value: string): FullCard;
    getCardholdername(): string;
    setCardholdername(value: string): FullCard;
    getNumber(): string;
    setNumber(value: string): FullCard;
    getCvv(): string;
    setCvv(value: string): FullCard;
    getToken(): string;
    setToken(value: string): FullCard;
    getMaskednumber(): string;
    setMaskednumber(value: string): FullCard;
    getExpirationyear(): number;
    setExpirationyear(value: number): FullCard;
    getExpirationmonth(): number;
    setExpirationmonth(value: number): FullCard;
    getActive(): boolean;
    setActive(value: boolean): FullCard;
    getIscredit(): boolean;
    setIscredit(value: boolean): FullCard;
    getIsdebit(): boolean;
    setIsdebit(value: boolean): FullCard;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FullCard.AsObject;
    static toObject(includeInstance: boolean, msg: FullCard): FullCard.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FullCard, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FullCard;
    static deserializeBinaryFromReader(message: FullCard, reader: jspb.BinaryReader): FullCard;
}

export namespace FullCard {
    export type AsObject = {
        id: string,
        cardholdername: string,
        number: string,
        cvv: string,
        token: string,
        maskednumber: string,
        expirationyear: number,
        expirationmonth: number,
        active: boolean,
        iscredit: boolean,
        isdebit: boolean,
    }
}