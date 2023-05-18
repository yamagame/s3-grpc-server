/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Table } from "./common";

export const protobufPackage = "";

export interface CreateTableRequest {
  table: Table | undefined;
}

export interface CreateTableResponse {
  ID: number;
}

export interface ReadTableRequest {
  ID: number;
}

export interface ReadTableResponse {
  ID: number;
  table: Table | undefined;
}

export interface UpdateTableRequest {
  ID: number;
  table: Table | undefined;
}

export interface UpdateTableResponse {
  ID: number;
}

export interface DeleteTableRequest {
  ID: number;
}

export interface DeleteTableResponse {
  ID: number;
}

/**  */
export interface ListTableRequest {
}

export interface ListTableResponse {
  tables: Table[];
}

function createBaseCreateTableRequest(): CreateTableRequest {
  return { table: undefined };
}

export const CreateTableRequest = {
  encode(message: CreateTableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.table !== undefined) {
      Table.encode(message.table, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTableRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTableRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.table = Table.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateTableRequest {
    return { table: isSet(object.table) ? Table.fromJSON(object.table) : undefined };
  },

  toJSON(message: CreateTableRequest): unknown {
    const obj: any = {};
    message.table !== undefined && (obj.table = message.table ? Table.toJSON(message.table) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTableRequest>, I>>(base?: I): CreateTableRequest {
    return CreateTableRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateTableRequest>, I>>(object: I): CreateTableRequest {
    const message = createBaseCreateTableRequest();
    message.table = (object.table !== undefined && object.table !== null) ? Table.fromPartial(object.table) : undefined;
    return message;
  },
};

function createBaseCreateTableResponse(): CreateTableResponse {
  return { ID: 0 };
}

export const CreateTableResponse = {
  encode(message: CreateTableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTableResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTableResponse();
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

  fromJSON(object: any): CreateTableResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: CreateTableResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTableResponse>, I>>(base?: I): CreateTableResponse {
    return CreateTableResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateTableResponse>, I>>(object: I): CreateTableResponse {
    const message = createBaseCreateTableResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadTableRequest(): ReadTableRequest {
  return { ID: 0 };
}

export const ReadTableRequest = {
  encode(message: ReadTableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadTableRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadTableRequest();
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

  fromJSON(object: any): ReadTableRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: ReadTableRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadTableRequest>, I>>(base?: I): ReadTableRequest {
    return ReadTableRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadTableRequest>, I>>(object: I): ReadTableRequest {
    const message = createBaseReadTableRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadTableResponse(): ReadTableResponse {
  return { ID: 0, table: undefined };
}

export const ReadTableResponse = {
  encode(message: ReadTableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    if (message.table !== undefined) {
      Table.encode(message.table, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadTableResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadTableResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.table = Table.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadTableResponse {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      table: isSet(object.table) ? Table.fromJSON(object.table) : undefined,
    };
  },

  toJSON(message: ReadTableResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.table !== undefined && (obj.table = message.table ? Table.toJSON(message.table) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadTableResponse>, I>>(base?: I): ReadTableResponse {
    return ReadTableResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadTableResponse>, I>>(object: I): ReadTableResponse {
    const message = createBaseReadTableResponse();
    message.ID = object.ID ?? 0;
    message.table = (object.table !== undefined && object.table !== null) ? Table.fromPartial(object.table) : undefined;
    return message;
  },
};

function createBaseUpdateTableRequest(): UpdateTableRequest {
  return { ID: 0, table: undefined };
}

export const UpdateTableRequest = {
  encode(message: UpdateTableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    if (message.table !== undefined) {
      Table.encode(message.table, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTableRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTableRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.ID = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.table = Table.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateTableRequest {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      table: isSet(object.table) ? Table.fromJSON(object.table) : undefined,
    };
  },

  toJSON(message: UpdateTableRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.table !== undefined && (obj.table = message.table ? Table.toJSON(message.table) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTableRequest>, I>>(base?: I): UpdateTableRequest {
    return UpdateTableRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateTableRequest>, I>>(object: I): UpdateTableRequest {
    const message = createBaseUpdateTableRequest();
    message.ID = object.ID ?? 0;
    message.table = (object.table !== undefined && object.table !== null) ? Table.fromPartial(object.table) : undefined;
    return message;
  },
};

function createBaseUpdateTableResponse(): UpdateTableResponse {
  return { ID: 0 };
}

export const UpdateTableResponse = {
  encode(message: UpdateTableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTableResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTableResponse();
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

  fromJSON(object: any): UpdateTableResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: UpdateTableResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTableResponse>, I>>(base?: I): UpdateTableResponse {
    return UpdateTableResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateTableResponse>, I>>(object: I): UpdateTableResponse {
    const message = createBaseUpdateTableResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteTableRequest(): DeleteTableRequest {
  return { ID: 0 };
}

export const DeleteTableRequest = {
  encode(message: DeleteTableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTableRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTableRequest();
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

  fromJSON(object: any): DeleteTableRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteTableRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTableRequest>, I>>(base?: I): DeleteTableRequest {
    return DeleteTableRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteTableRequest>, I>>(object: I): DeleteTableRequest {
    const message = createBaseDeleteTableRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteTableResponse(): DeleteTableResponse {
  return { ID: 0 };
}

export const DeleteTableResponse = {
  encode(message: DeleteTableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTableResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTableResponse();
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

  fromJSON(object: any): DeleteTableResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteTableResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTableResponse>, I>>(base?: I): DeleteTableResponse {
    return DeleteTableResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteTableResponse>, I>>(object: I): DeleteTableResponse {
    const message = createBaseDeleteTableResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseListTableRequest(): ListTableRequest {
  return {};
}

export const ListTableRequest = {
  encode(_: ListTableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListTableRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListTableRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): ListTableRequest {
    return {};
  },

  toJSON(_: ListTableRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListTableRequest>, I>>(base?: I): ListTableRequest {
    return ListTableRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListTableRequest>, I>>(_: I): ListTableRequest {
    const message = createBaseListTableRequest();
    return message;
  },
};

function createBaseListTableResponse(): ListTableResponse {
  return { tables: [] };
}

export const ListTableResponse = {
  encode(message: ListTableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tables) {
      Table.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListTableResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListTableResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tables.push(Table.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListTableResponse {
    return { tables: Array.isArray(object?.tables) ? object.tables.map((e: any) => Table.fromJSON(e)) : [] };
  },

  toJSON(message: ListTableResponse): unknown {
    const obj: any = {};
    if (message.tables) {
      obj.tables = message.tables.map((e) => e ? Table.toJSON(e) : undefined);
    } else {
      obj.tables = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListTableResponse>, I>>(base?: I): ListTableResponse {
    return ListTableResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListTableResponse>, I>>(object: I): ListTableResponse {
    const message = createBaseListTableResponse();
    message.tables = object.tables?.map((e) => Table.fromPartial(e)) || [];
    return message;
  },
};

export interface TableRepository {
  /** Create */
  Create(request: CreateTableRequest): Promise<CreateTableResponse>;
  /** Read */
  Read(request: ReadTableRequest): Promise<ReadTableResponse>;
  /** Update */
  Update(request: UpdateTableRequest): Promise<UpdateTableResponse>;
  /** Delete */
  Delete(request: DeleteTableRequest): Promise<DeleteTableResponse>;
  /** List */
  List(request: ListTableRequest): Promise<ListTableResponse>;
}

export class TableRepositoryClientImpl implements TableRepository {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "TableRepository";
    this.rpc = rpc;
    this.Create = this.Create.bind(this);
    this.Read = this.Read.bind(this);
    this.Update = this.Update.bind(this);
    this.Delete = this.Delete.bind(this);
    this.List = this.List.bind(this);
  }
  Create(request: CreateTableRequest): Promise<CreateTableResponse> {
    const data = CreateTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Create", data);
    return promise.then((data) => CreateTableResponse.decode(_m0.Reader.create(data)));
  }

  Read(request: ReadTableRequest): Promise<ReadTableResponse> {
    const data = ReadTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Read", data);
    return promise.then((data) => ReadTableResponse.decode(_m0.Reader.create(data)));
  }

  Update(request: UpdateTableRequest): Promise<UpdateTableResponse> {
    const data = UpdateTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Update", data);
    return promise.then((data) => UpdateTableResponse.decode(_m0.Reader.create(data)));
  }

  Delete(request: DeleteTableRequest): Promise<DeleteTableResponse> {
    const data = DeleteTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Delete", data);
    return promise.then((data) => DeleteTableResponse.decode(_m0.Reader.create(data)));
  }

  List(request: ListTableRequest): Promise<ListTableResponse> {
    const data = ListTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "List", data);
    return promise.then((data) => ListTableResponse.decode(_m0.Reader.create(data)));
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
