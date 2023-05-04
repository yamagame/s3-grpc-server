/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Cell } from "./common";

export const protobufPackage = "";

export interface CreateCellRequest {
  cell: Cell | undefined;
}

export interface CreateCellResponse {
  ID: number;
}

export interface ReadCellRequest {
  ID: number;
}

export interface ReadCellResponse {
  cell: Cell | undefined;
}

export interface UpdateCellRequest {
  cell: Cell | undefined;
}

export interface UpdateCellResponse {
  ID: number;
}

export interface DeleteCellRequest {
  ID: number;
}

export interface DeleteCellResponse {
  ID: number;
}

function createBaseCreateCellRequest(): CreateCellRequest {
  return { cell: undefined };
}

export const CreateCellRequest = {
  encode(message: CreateCellRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cell !== undefined) {
      Cell.encode(message.cell, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateCellRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateCellRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.cell = Cell.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateCellRequest {
    return { cell: isSet(object.cell) ? Cell.fromJSON(object.cell) : undefined };
  },

  toJSON(message: CreateCellRequest): unknown {
    const obj: any = {};
    message.cell !== undefined && (obj.cell = message.cell ? Cell.toJSON(message.cell) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCellRequest>, I>>(base?: I): CreateCellRequest {
    return CreateCellRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateCellRequest>, I>>(object: I): CreateCellRequest {
    const message = createBaseCreateCellRequest();
    message.cell = (object.cell !== undefined && object.cell !== null) ? Cell.fromPartial(object.cell) : undefined;
    return message;
  },
};

function createBaseCreateCellResponse(): CreateCellResponse {
  return { ID: 0 };
}

export const CreateCellResponse = {
  encode(message: CreateCellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateCellResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateCellResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateCellResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: CreateCellResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCellResponse>, I>>(base?: I): CreateCellResponse {
    return CreateCellResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateCellResponse>, I>>(object: I): CreateCellResponse {
    const message = createBaseCreateCellResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadCellRequest(): ReadCellRequest {
  return { ID: 0 };
}

export const ReadCellRequest = {
  encode(message: ReadCellRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadCellRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadCellRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadCellRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: ReadCellRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadCellRequest>, I>>(base?: I): ReadCellRequest {
    return ReadCellRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadCellRequest>, I>>(object: I): ReadCellRequest {
    const message = createBaseReadCellRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadCellResponse(): ReadCellResponse {
  return { cell: undefined };
}

export const ReadCellResponse = {
  encode(message: ReadCellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cell !== undefined) {
      Cell.encode(message.cell, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadCellResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadCellResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.cell = Cell.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadCellResponse {
    return { cell: isSet(object.cell) ? Cell.fromJSON(object.cell) : undefined };
  },

  toJSON(message: ReadCellResponse): unknown {
    const obj: any = {};
    message.cell !== undefined && (obj.cell = message.cell ? Cell.toJSON(message.cell) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadCellResponse>, I>>(base?: I): ReadCellResponse {
    return ReadCellResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadCellResponse>, I>>(object: I): ReadCellResponse {
    const message = createBaseReadCellResponse();
    message.cell = (object.cell !== undefined && object.cell !== null) ? Cell.fromPartial(object.cell) : undefined;
    return message;
  },
};

function createBaseUpdateCellRequest(): UpdateCellRequest {
  return { cell: undefined };
}

export const UpdateCellRequest = {
  encode(message: UpdateCellRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cell !== undefined) {
      Cell.encode(message.cell, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateCellRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateCellRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.cell = Cell.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateCellRequest {
    return { cell: isSet(object.cell) ? Cell.fromJSON(object.cell) : undefined };
  },

  toJSON(message: UpdateCellRequest): unknown {
    const obj: any = {};
    message.cell !== undefined && (obj.cell = message.cell ? Cell.toJSON(message.cell) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateCellRequest>, I>>(base?: I): UpdateCellRequest {
    return UpdateCellRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateCellRequest>, I>>(object: I): UpdateCellRequest {
    const message = createBaseUpdateCellRequest();
    message.cell = (object.cell !== undefined && object.cell !== null) ? Cell.fromPartial(object.cell) : undefined;
    return message;
  },
};

function createBaseUpdateCellResponse(): UpdateCellResponse {
  return { ID: 0 };
}

export const UpdateCellResponse = {
  encode(message: UpdateCellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateCellResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateCellResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateCellResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: UpdateCellResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateCellResponse>, I>>(base?: I): UpdateCellResponse {
    return UpdateCellResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateCellResponse>, I>>(object: I): UpdateCellResponse {
    const message = createBaseUpdateCellResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteCellRequest(): DeleteCellRequest {
  return { ID: 0 };
}

export const DeleteCellRequest = {
  encode(message: DeleteCellRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteCellRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteCellRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteCellRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteCellRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteCellRequest>, I>>(base?: I): DeleteCellRequest {
    return DeleteCellRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteCellRequest>, I>>(object: I): DeleteCellRequest {
    const message = createBaseDeleteCellRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteCellResponse(): DeleteCellResponse {
  return { ID: 0 };
}

export const DeleteCellResponse = {
  encode(message: DeleteCellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteCellResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteCellResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteCellResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteCellResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteCellResponse>, I>>(base?: I): DeleteCellResponse {
    return DeleteCellResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteCellResponse>, I>>(object: I): DeleteCellResponse {
    const message = createBaseDeleteCellResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

export interface CellRepository {
  /** Create */
  Create(request: CreateCellRequest): Promise<CreateCellResponse>;
  /** Read */
  Read(request: ReadCellRequest): Promise<ReadCellResponse>;
  /** Update */
  Update(request: UpdateCellRequest): Promise<UpdateCellResponse>;
  /** Delete */
  Delete(request: DeleteCellRequest): Promise<DeleteCellResponse>;
}

export class CellRepositoryClientImpl implements CellRepository {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "CellRepository";
    this.rpc = rpc;
    this.Create = this.Create.bind(this);
    this.Read = this.Read.bind(this);
    this.Update = this.Update.bind(this);
    this.Delete = this.Delete.bind(this);
  }
  Create(request: CreateCellRequest): Promise<CreateCellResponse> {
    const data = CreateCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Create", data);
    return promise.then((data) => CreateCellResponse.decode(_m0.Reader.create(data)));
  }

  Read(request: ReadCellRequest): Promise<ReadCellResponse> {
    const data = ReadCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Read", data);
    return promise.then((data) => ReadCellResponse.decode(_m0.Reader.create(data)));
  }

  Update(request: UpdateCellRequest): Promise<UpdateCellResponse> {
    const data = UpdateCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Update", data);
    return promise.then((data) => UpdateCellResponse.decode(_m0.Reader.create(data)));
  }

  Delete(request: DeleteCellRequest): Promise<DeleteCellResponse> {
    const data = DeleteCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Delete", data);
    return promise.then((data) => DeleteCellResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
