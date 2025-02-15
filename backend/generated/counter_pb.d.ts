import * as jspb from 'google-protobuf'



export class IncrementRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IncrementRequest.AsObject;
  static toObject(includeInstance: boolean, msg: IncrementRequest): IncrementRequest.AsObject;
  static serializeBinaryToWriter(message: IncrementRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IncrementRequest;
  static deserializeBinaryFromReader(message: IncrementRequest, reader: jspb.BinaryReader): IncrementRequest;
}

export namespace IncrementRequest {
  export type AsObject = {
  }
}

export class IncrementResponse extends jspb.Message {
  getValue(): number;
  setValue(value: number): IncrementResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IncrementResponse.AsObject;
  static toObject(includeInstance: boolean, msg: IncrementResponse): IncrementResponse.AsObject;
  static serializeBinaryToWriter(message: IncrementResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IncrementResponse;
  static deserializeBinaryFromReader(message: IncrementResponse, reader: jspb.BinaryReader): IncrementResponse;
}

export namespace IncrementResponse {
  export type AsObject = {
    value: number,
  }
}

