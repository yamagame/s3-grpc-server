/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export interface FileInfo {
  filename: string;
}

export interface CreateFileInfoRequest {
  file: FileInfo | undefined;
}

export interface CreateFileInfoResponse {
  ID: number;
}

export interface ReadFileInfoRequest {
  ID: number;
}

export interface ReadFileInfoResponse {
  ID: number;
  file: FileInfo | undefined;
}

export interface UpdateFileInfoRequest {
  ID: number;
  file: FileInfo | undefined;
}

export interface UpdateFileInfoResponse {
  ID: number;
}

export interface DeleteFileInfoRequest {
  ID: number;
}

export interface DeleteFileInfoResponse {
  ID: number;
}

function createBaseFileInfo(): FileInfo {
  return { filename: "" };
}

export const FileInfo = {
  encode(message: FileInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.filename !== "") {
      writer.uint32(10).string(message.filename);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FileInfo {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFileInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.filename = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FileInfo {
    return { filename: isSet(object.filename) ? String(object.filename) : "" };
  },

  toJSON(message: FileInfo): unknown {
    const obj: any = {};
    message.filename !== undefined && (obj.filename = message.filename);
    return obj;
  },

  create<I extends Exact<DeepPartial<FileInfo>, I>>(base?: I): FileInfo {
    return FileInfo.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FileInfo>, I>>(object: I): FileInfo {
    const message = createBaseFileInfo();
    message.filename = object.filename ?? "";
    return message;
  },
};

function createBaseCreateFileInfoRequest(): CreateFileInfoRequest {
  return { file: undefined };
}

export const CreateFileInfoRequest = {
  encode(message: CreateFileInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.file !== undefined) {
      FileInfo.encode(message.file, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFileInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFileInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.file = FileInfo.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateFileInfoRequest {
    return { file: isSet(object.file) ? FileInfo.fromJSON(object.file) : undefined };
  },

  toJSON(message: CreateFileInfoRequest): unknown {
    const obj: any = {};
    message.file !== undefined && (obj.file = message.file ? FileInfo.toJSON(message.file) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateFileInfoRequest>, I>>(base?: I): CreateFileInfoRequest {
    return CreateFileInfoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateFileInfoRequest>, I>>(object: I): CreateFileInfoRequest {
    const message = createBaseCreateFileInfoRequest();
    message.file = (object.file !== undefined && object.file !== null) ? FileInfo.fromPartial(object.file) : undefined;
    return message;
  },
};

function createBaseCreateFileInfoResponse(): CreateFileInfoResponse {
  return { ID: 0 };
}

export const CreateFileInfoResponse = {
  encode(message: CreateFileInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFileInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFileInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateFileInfoResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: CreateFileInfoResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateFileInfoResponse>, I>>(base?: I): CreateFileInfoResponse {
    return CreateFileInfoResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateFileInfoResponse>, I>>(object: I): CreateFileInfoResponse {
    const message = createBaseCreateFileInfoResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadFileInfoRequest(): ReadFileInfoRequest {
  return { ID: 0 };
}

export const ReadFileInfoRequest = {
  encode(message: ReadFileInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadFileInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadFileInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadFileInfoRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: ReadFileInfoRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadFileInfoRequest>, I>>(base?: I): ReadFileInfoRequest {
    return ReadFileInfoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadFileInfoRequest>, I>>(object: I): ReadFileInfoRequest {
    const message = createBaseReadFileInfoRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseReadFileInfoResponse(): ReadFileInfoResponse {
  return { ID: 0, file: undefined };
}

export const ReadFileInfoResponse = {
  encode(message: ReadFileInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    if (message.file !== undefined) {
      FileInfo.encode(message.file, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadFileInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadFileInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.file = FileInfo.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadFileInfoResponse {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      file: isSet(object.file) ? FileInfo.fromJSON(object.file) : undefined,
    };
  },

  toJSON(message: ReadFileInfoResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.file !== undefined && (obj.file = message.file ? FileInfo.toJSON(message.file) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadFileInfoResponse>, I>>(base?: I): ReadFileInfoResponse {
    return ReadFileInfoResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadFileInfoResponse>, I>>(object: I): ReadFileInfoResponse {
    const message = createBaseReadFileInfoResponse();
    message.ID = object.ID ?? 0;
    message.file = (object.file !== undefined && object.file !== null) ? FileInfo.fromPartial(object.file) : undefined;
    return message;
  },
};

function createBaseUpdateFileInfoRequest(): UpdateFileInfoRequest {
  return { ID: 0, file: undefined };
}

export const UpdateFileInfoRequest = {
  encode(message: UpdateFileInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    if (message.file !== undefined) {
      FileInfo.encode(message.file, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateFileInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateFileInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.file = FileInfo.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateFileInfoRequest {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      file: isSet(object.file) ? FileInfo.fromJSON(object.file) : undefined,
    };
  },

  toJSON(message: UpdateFileInfoRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.file !== undefined && (obj.file = message.file ? FileInfo.toJSON(message.file) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateFileInfoRequest>, I>>(base?: I): UpdateFileInfoRequest {
    return UpdateFileInfoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateFileInfoRequest>, I>>(object: I): UpdateFileInfoRequest {
    const message = createBaseUpdateFileInfoRequest();
    message.ID = object.ID ?? 0;
    message.file = (object.file !== undefined && object.file !== null) ? FileInfo.fromPartial(object.file) : undefined;
    return message;
  },
};

function createBaseUpdateFileInfoResponse(): UpdateFileInfoResponse {
  return { ID: 0 };
}

export const UpdateFileInfoResponse = {
  encode(message: UpdateFileInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateFileInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateFileInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateFileInfoResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: UpdateFileInfoResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateFileInfoResponse>, I>>(base?: I): UpdateFileInfoResponse {
    return UpdateFileInfoResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateFileInfoResponse>, I>>(object: I): UpdateFileInfoResponse {
    const message = createBaseUpdateFileInfoResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteFileInfoRequest(): DeleteFileInfoRequest {
  return { ID: 0 };
}

export const DeleteFileInfoRequest = {
  encode(message: DeleteFileInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFileInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFileInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteFileInfoRequest {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteFileInfoRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFileInfoRequest>, I>>(base?: I): DeleteFileInfoRequest {
    return DeleteFileInfoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFileInfoRequest>, I>>(object: I): DeleteFileInfoRequest {
    const message = createBaseDeleteFileInfoRequest();
    message.ID = object.ID ?? 0;
    return message;
  },
};

function createBaseDeleteFileInfoResponse(): DeleteFileInfoResponse {
  return { ID: 0 };
}

export const DeleteFileInfoResponse = {
  encode(message: DeleteFileInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFileInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFileInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteFileInfoResponse {
    return { ID: isSet(object.ID) ? Number(object.ID) : 0 };
  },

  toJSON(message: DeleteFileInfoResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFileInfoResponse>, I>>(base?: I): DeleteFileInfoResponse {
    return DeleteFileInfoResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFileInfoResponse>, I>>(object: I): DeleteFileInfoResponse {
    const message = createBaseDeleteFileInfoResponse();
    message.ID = object.ID ?? 0;
    return message;
  },
};

/** The greeting service definition. */
export interface repository {
  /** CreateFileInfo */
  CreateFileInfo(request: CreateFileInfoRequest): Promise<CreateFileInfoResponse>;
  /** ReadFileInfo */
  ReadFileInfo(request: ReadFileInfoRequest): Promise<ReadFileInfoResponse>;
  /** UpdateFileInfo */
  UpdateFileInfo(request: UpdateFileInfoRequest): Promise<UpdateFileInfoResponse>;
  /** DeleteFileInfo */
  DeleteFileInfo(request: DeleteFileInfoRequest): Promise<DeleteFileInfoResponse>;
}

export class repositoryClientImpl implements repository {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "repository";
    this.rpc = rpc;
    this.CreateFileInfo = this.CreateFileInfo.bind(this);
    this.ReadFileInfo = this.ReadFileInfo.bind(this);
    this.UpdateFileInfo = this.UpdateFileInfo.bind(this);
    this.DeleteFileInfo = this.DeleteFileInfo.bind(this);
  }
  CreateFileInfo(request: CreateFileInfoRequest): Promise<CreateFileInfoResponse> {
    const data = CreateFileInfoRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateFileInfo", data);
    return promise.then((data) => CreateFileInfoResponse.decode(_m0.Reader.create(data)));
  }

  ReadFileInfo(request: ReadFileInfoRequest): Promise<ReadFileInfoResponse> {
    const data = ReadFileInfoRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ReadFileInfo", data);
    return promise.then((data) => ReadFileInfoResponse.decode(_m0.Reader.create(data)));
  }

  UpdateFileInfo(request: UpdateFileInfoRequest): Promise<UpdateFileInfoResponse> {
    const data = UpdateFileInfoRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateFileInfo", data);
    return promise.then((data) => UpdateFileInfoResponse.decode(_m0.Reader.create(data)));
  }

  DeleteFileInfo(request: DeleteFileInfoRequest): Promise<DeleteFileInfoResponse> {
    const data = DeleteFileInfoRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteFileInfo", data);
    return promise.then((data) => DeleteFileInfoResponse.decode(_m0.Reader.create(data)));
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
