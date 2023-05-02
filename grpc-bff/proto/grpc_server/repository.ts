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

export interface Table {
  title: string;
}

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

export interface Cell {
  text: string;
}

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
  ID: number;
  cell: Cell | undefined;
}

export interface UpdateCellRequest {
  ID: number;
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

function createBaseTable(): Table {
  return { title: "" };
}

export const Table = {
  encode(message: Table, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Table {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTable();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.title = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Table {
    return { title: isSet(object.title) ? String(object.title) : "" };
  },

  toJSON(message: Table): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    return obj;
  },

  create<I extends Exact<DeepPartial<Table>, I>>(base?: I): Table {
    return Table.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Table>, I>>(object: I): Table {
    const message = createBaseTable();
    message.title = object.title ?? "";
    return message;
  },
};

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
          if (tag != 10) {
            break;
          }

          message.table = Table.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.table = Table.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.table = Table.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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

function createBaseCell(): Cell {
  return { text: "" };
}

export const Cell = {
  encode(message: Cell, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Cell {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCell();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Cell {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: Cell): unknown {
    const obj: any = {};
    message.text !== undefined && (obj.text = message.text);
    return obj;
  },

  create<I extends Exact<DeepPartial<Cell>, I>>(base?: I): Cell {
    return Cell.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Cell>, I>>(object: I): Cell {
    const message = createBaseCell();
    message.text = object.text ?? "";
    return message;
  },
};

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
          if (tag != 10) {
            break;
          }

          message.cell = Cell.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
  return { ID: 0, cell: undefined };
}

export const ReadCellResponse = {
  encode(message: ReadCellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    if (message.cell !== undefined) {
      Cell.encode(message.cell, writer.uint32(18).fork()).ldelim();
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
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.cell = Cell.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadCellResponse {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      cell: isSet(object.cell) ? Cell.fromJSON(object.cell) : undefined,
    };
  },

  toJSON(message: ReadCellResponse): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.cell !== undefined && (obj.cell = message.cell ? Cell.toJSON(message.cell) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadCellResponse>, I>>(base?: I): ReadCellResponse {
    return ReadCellResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadCellResponse>, I>>(object: I): ReadCellResponse {
    const message = createBaseReadCellResponse();
    message.ID = object.ID ?? 0;
    message.cell = (object.cell !== undefined && object.cell !== null) ? Cell.fromPartial(object.cell) : undefined;
    return message;
  },
};

function createBaseUpdateCellRequest(): UpdateCellRequest {
  return { ID: 0, cell: undefined };
}

export const UpdateCellRequest = {
  encode(message: UpdateCellRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ID !== 0) {
      writer.uint32(8).int64(message.ID);
    }
    if (message.cell !== undefined) {
      Cell.encode(message.cell, writer.uint32(18).fork()).ldelim();
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
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.cell = Cell.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateCellRequest {
    return {
      ID: isSet(object.ID) ? Number(object.ID) : 0,
      cell: isSet(object.cell) ? Cell.fromJSON(object.cell) : undefined,
    };
  },

  toJSON(message: UpdateCellRequest): unknown {
    const obj: any = {};
    message.ID !== undefined && (obj.ID = Math.round(message.ID));
    message.cell !== undefined && (obj.cell = message.cell ? Cell.toJSON(message.cell) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateCellRequest>, I>>(base?: I): UpdateCellRequest {
    return UpdateCellRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateCellRequest>, I>>(object: I): UpdateCellRequest {
    const message = createBaseUpdateCellRequest();
    message.ID = object.ID ?? 0;
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
  /** CreateTable */
  CreateTable(request: CreateTableRequest): Promise<CreateTableResponse>;
  /** ReadTable */
  ReadTable(request: ReadTableRequest): Promise<ReadTableResponse>;
  /** UpdateTable */
  UpdateTable(request: UpdateTableRequest): Promise<UpdateTableResponse>;
  /** DeleteTable */
  DeleteTable(request: DeleteTableRequest): Promise<DeleteTableResponse>;
  /** CreateCell */
  CreateCell(request: CreateCellRequest): Promise<CreateCellResponse>;
  /** ReadCell */
  ReadCell(request: ReadCellRequest): Promise<ReadCellResponse>;
  /** UpdateCell */
  UpdateCell(request: UpdateCellRequest): Promise<UpdateCellResponse>;
  /** DeleteCell */
  DeleteCell(request: DeleteCellRequest): Promise<DeleteCellResponse>;
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
    this.CreateTable = this.CreateTable.bind(this);
    this.ReadTable = this.ReadTable.bind(this);
    this.UpdateTable = this.UpdateTable.bind(this);
    this.DeleteTable = this.DeleteTable.bind(this);
    this.CreateCell = this.CreateCell.bind(this);
    this.ReadCell = this.ReadCell.bind(this);
    this.UpdateCell = this.UpdateCell.bind(this);
    this.DeleteCell = this.DeleteCell.bind(this);
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

  CreateTable(request: CreateTableRequest): Promise<CreateTableResponse> {
    const data = CreateTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateTable", data);
    return promise.then((data) => CreateTableResponse.decode(_m0.Reader.create(data)));
  }

  ReadTable(request: ReadTableRequest): Promise<ReadTableResponse> {
    const data = ReadTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ReadTable", data);
    return promise.then((data) => ReadTableResponse.decode(_m0.Reader.create(data)));
  }

  UpdateTable(request: UpdateTableRequest): Promise<UpdateTableResponse> {
    const data = UpdateTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateTable", data);
    return promise.then((data) => UpdateTableResponse.decode(_m0.Reader.create(data)));
  }

  DeleteTable(request: DeleteTableRequest): Promise<DeleteTableResponse> {
    const data = DeleteTableRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteTable", data);
    return promise.then((data) => DeleteTableResponse.decode(_m0.Reader.create(data)));
  }

  CreateCell(request: CreateCellRequest): Promise<CreateCellResponse> {
    const data = CreateCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateCell", data);
    return promise.then((data) => CreateCellResponse.decode(_m0.Reader.create(data)));
  }

  ReadCell(request: ReadCellRequest): Promise<ReadCellResponse> {
    const data = ReadCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ReadCell", data);
    return promise.then((data) => ReadCellResponse.decode(_m0.Reader.create(data)));
  }

  UpdateCell(request: UpdateCellRequest): Promise<UpdateCellResponse> {
    const data = UpdateCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateCell", data);
    return promise.then((data) => UpdateCellResponse.decode(_m0.Reader.create(data)));
  }

  DeleteCell(request: DeleteCellRequest): Promise<DeleteCellResponse> {
    const data = DeleteCellRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteCell", data);
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
