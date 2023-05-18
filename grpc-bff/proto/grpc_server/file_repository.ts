/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { File } from "./common";

export const protobufPackage = "";

export interface CreateFileRequest {
  file: File | undefined;
}

export interface CreateFileResponse {
  ID: number;
}

export interface ReadFileRequest {
  ID: number;
}

export interface ReadFileResponse {
  ID: number;
  file: File | undefined;
}

export interface UpdateFileRequest {
  ID: number;
  file: File | undefined;
}

export interface UpdateFileResponse {
  ID: number;
}

export interface DeleteFileRequest {
  ID: number;
}

export interface DeleteFileResponse {
  ID: number;
}

/**  */
export interface ListFileRequest {
}

export interface ListFileResponse {
  files: File[];
}

function createBaseCreateFileRequest(): CreateFileRequest {
  return { file: undefined };
}

export const CreateFileRequest = {
  encode(message: CreateFileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.file !== undefined) {
      File.encode(message.file, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.file = File.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateFileRequest {
    return { file: isSet(object.file) ? File.fromJSON(object.file) : undefined };
  },

  toJSON(message: CreateFileRequest): unknown {
    const obj: any = {};
    message.file !== undefined && (obj.file = message.file ? File.toJSON(message.file) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateFileRequest>, I>>(base?: I): CreateFileRequest {
    return CreateFileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateFileRequest>, I>>(object: I): CreateFileRequest {
    const message = createBaseCreateFileRequest();
    message.file = (object.file !== undefined && object.file !== null) ? File.fromPartial(object.file) : undefined;
    return message;
  },
};

function createBaseCreateFileResponse(): CreateFileResponse {
  return { ID: 0 };
}

export const CreateFileResponse = {
  encode(message: CreateFileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFileResponse();
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

  fromJSON(object: any): CreateFileResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: CreateFileResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateFileResponse>, I>>(base?: I): CreateFileResponse {
    return CreateFileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateFileResponse>, I>>(object: I): CreateFileResponse {
    const message = createBaseCreateFileResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadFileRequest(): ReadFileRequest {
  return { ID: 0 };
}

export const ReadFileRequest = {
  encode(message: ReadFileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadFileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadFileRequest();
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

  fromJSON(object: any): ReadFileRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: ReadFileRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadFileRequest>, I>>(base?: I): ReadFileRequest {
    return ReadFileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadFileRequest>, I>>(object: I): ReadFileRequest {
    const message = createBaseReadFileRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadFileResponse(): ReadFileResponse {
  return { ID: 0, file: undefined };
}

export const ReadFileResponse = {
  encode(message: ReadFileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    if (message.file !== undefined) {
      File.encode(message.file, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadFileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadFileResponse();
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

          message.file = File.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadFileResponse {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      file: isSet(object.file) ? File.fromJSON(object.file) : undefined,
    };
  },

  toJSON(message: ReadFileResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.file !== undefined && (obj.file = message.file ? File.toJSON(message.file) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadFileResponse>, I>>(base?: I): ReadFileResponse {
    return ReadFileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadFileResponse>, I>>(object: I): ReadFileResponse {
    const message = createBaseReadFileResponse();
    message.ID = object.ID ?? 0;
    message.file = (object.file !== undefined && object.file !== null) ? File.fromPartial(object.file) : undefined;
    return message;
  },
};

function createBaseUpdateFileRequest(): UpdateFileRequest {
  return { ID: 0, file: undefined };
}

export const UpdateFileRequest = {
  encode(message: UpdateFileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    if (message.file !== undefined) {
      File.encode(message.file, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateFileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateFileRequest();
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

          message.file = File.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateFileRequest {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      file: isSet(object.file) ? File.fromJSON(object.file) : undefined,
    };
  },

  toJSON(message: UpdateFileRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.file !== undefined && (obj.file = message.file ? File.toJSON(message.file) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateFileRequest>, I>>(base?: I): UpdateFileRequest {
    return UpdateFileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateFileRequest>, I>>(object: I): UpdateFileRequest {
    const message = createBaseUpdateFileRequest();
    message.ID = object.ID ?? 0;
    message.file = (object.file !== undefined && object.file !== null) ? File.fromPartial(object.file) : undefined;
    return message;
  },
};

function createBaseUpdateFileResponse(): UpdateFileResponse {
  return { ID: 0 };
}

export const UpdateFileResponse = {
  encode(message: UpdateFileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateFileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateFileResponse();
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

  fromJSON(object: any): UpdateFileResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: UpdateFileResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateFileResponse>, I>>(base?: I): UpdateFileResponse {
    return UpdateFileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateFileResponse>, I>>(object: I): UpdateFileResponse {
    const message = createBaseUpdateFileResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteFileRequest(): DeleteFileRequest {
  return { ID: 0 };
}

export const DeleteFileRequest = {
  encode(message: DeleteFileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFileRequest();
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

  fromJSON(object: any): DeleteFileRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteFileRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFileRequest>, I>>(base?: I): DeleteFileRequest {
    return DeleteFileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFileRequest>, I>>(object: I): DeleteFileRequest {
    const message = createBaseDeleteFileRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteFileResponse(): DeleteFileResponse {
  return { ID: 0 };
}

export const DeleteFileResponse = {
  encode(message: DeleteFileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).uint64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFileResponse();
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

  fromJSON(object: any): DeleteFileResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteFileResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFileResponse>, I>>(base?: I): DeleteFileResponse {
    return DeleteFileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFileResponse>, I>>(object: I): DeleteFileResponse {
    const message = createBaseDeleteFileResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseListFileRequest(): ListFileRequest {
  return {};
}

export const ListFileRequest = {
  encode(_: ListFileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFileRequest();
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

  fromJSON(_: any): ListFileRequest {
    return {};
  },

  toJSON(_: ListFileRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFileRequest>, I>>(base?: I): ListFileRequest {
    return ListFileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFileRequest>, I>>(_: I): ListFileRequest {
    const message = createBaseListFileRequest();
    return message;
  },
};

function createBaseListFileResponse(): ListFileResponse {
  return { files: [] };
}

export const ListFileResponse = {
  encode(message: ListFileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.files) {
      File.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFileResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.files.push(File.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListFileResponse {
    return { files: Array.isArray(object?.files) ? object.files.map((e: any) => File.fromJSON(e)) : [] };
  },

  toJSON(message: ListFileResponse): unknown {
    const obj: any = {};
    if (message.files) {
      obj.files = message.files.map((e) => e ? File.toJSON(e) : undefined);
    } else {
      obj.files = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFileResponse>, I>>(base?: I): ListFileResponse {
    return ListFileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFileResponse>, I>>(object: I): ListFileResponse {
    const message = createBaseListFileResponse();
    message.files = object.files?.map((e) => File.fromPartial(e)) || [];
    return message;
  },
};

export interface FileRepository {
  /** Create */
  Create(request: CreateFileRequest): Promise<CreateFileResponse>;
  /** Read */
  Read(request: ReadFileRequest): Promise<ReadFileResponse>;
  /** Update */
  Update(request: UpdateFileRequest): Promise<UpdateFileResponse>;
  /** Delete */
  Delete(request: DeleteFileRequest): Promise<DeleteFileResponse>;
  /** List */
  List(request: ListFileRequest): Promise<ListFileResponse>;
}

export class FileRepositoryClientImpl implements FileRepository {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "FileRepository";
    this.rpc = rpc;
    this.Create = this.Create.bind(this);
    this.Read = this.Read.bind(this);
    this.Update = this.Update.bind(this);
    this.Delete = this.Delete.bind(this);
    this.List = this.List.bind(this);
  }
  Create(request: CreateFileRequest): Promise<CreateFileResponse> {
    const data = CreateFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Create", data);
    return promise.then((data) => CreateFileResponse.decode(_m0.Reader.create(data)));
  }

  Read(request: ReadFileRequest): Promise<ReadFileResponse> {
    const data = ReadFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Read", data);
    return promise.then((data) => ReadFileResponse.decode(_m0.Reader.create(data)));
  }

  Update(request: UpdateFileRequest): Promise<UpdateFileResponse> {
    const data = UpdateFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Update", data);
    return promise.then((data) => UpdateFileResponse.decode(_m0.Reader.create(data)));
  }

  Delete(request: DeleteFileRequest): Promise<DeleteFileResponse> {
    const data = DeleteFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Delete", data);
    return promise.then((data) => DeleteFileResponse.decode(_m0.Reader.create(data)));
  }

  List(request: ListFileRequest): Promise<ListFileResponse> {
    const data = ListFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "List", data);
    return promise.then((data) => ListFileResponse.decode(_m0.Reader.create(data)));
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
