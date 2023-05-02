/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export interface File {
  id: number;
  filename: string;
  tables: Table[];
}

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

export interface Table {
  id: number;
  fileId: number;
  title: string;
  cells: Cell[];
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
  id: number;
  row: number;
  col: number;
  text: string;
  tableId: number;
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

function createBaseFile(): File {
  return { id: 0, filename: "", tables: [] };
}

export const File = {
  encode(message: File, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.filename !== "") {
      writer.uint32(18).string(message.filename);
    }
    for (const v of message.tables) {
      Table.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): File {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFile();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.id = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.filename = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.tables.push(Table.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): File {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      filename: isSet(object.filename) ? String(object.filename) : "",
      tables: Array.isArray(object?.tables) ? object.tables.map((e: any) => Table.fromJSON(e)) : [],
    };
  },

  toJSON(message: File): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.filename !== undefined && (obj.filename = message.filename);
    if (message.tables) {
      obj.tables = message.tables.map((e) => e ? Table.toJSON(e) : undefined);
    } else {
      obj.tables = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<File>, I>>(base?: I): File {
    return File.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<File>, I>>(object: I): File {
    const message = createBaseFile();
    message.id = object.id ?? 0;
    message.filename = object.filename ?? "";
    message.tables = object.tables?.map((e) => Table.fromPartial(e)) || [];
    return message;
  },
};

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
          if (tag != 10) {
            break;
          }

          message.file = File.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.file = File.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
          if (tag != 8) {
            break;
          }

          message.ID = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.file = File.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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
      writer.uint32(8).int64(message.ID);
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

function createBaseTable(): Table {
  return { id: 0, fileId: 0, title: "", cells: [] };
}

export const Table = {
  encode(message: Table, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.fileId !== 0) {
      writer.uint32(16).int64(message.fileId);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    for (const v of message.cells) {
      Cell.encode(v!, writer.uint32(34).fork()).ldelim();
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
          if (tag != 8) {
            break;
          }

          message.id = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.fileId = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.title = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.cells.push(Cell.decode(reader, reader.uint32()));
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
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      fileId: isSet(object.fileId) ? Number(object.fileId) : 0,
      title: isSet(object.title) ? String(object.title) : "",
      cells: Array.isArray(object?.cells) ? object.cells.map((e: any) => Cell.fromJSON(e)) : [],
    };
  },

  toJSON(message: Table): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.fileId !== undefined && (obj.fileId = Math.round(message.fileId));
    message.title !== undefined && (obj.title = message.title);
    if (message.cells) {
      obj.cells = message.cells.map((e) => e ? Cell.toJSON(e) : undefined);
    } else {
      obj.cells = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Table>, I>>(base?: I): Table {
    return Table.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Table>, I>>(object: I): Table {
    const message = createBaseTable();
    message.id = object.id ?? 0;
    message.fileId = object.fileId ?? 0;
    message.title = object.title ?? "";
    message.cells = object.cells?.map((e) => Cell.fromPartial(e)) || [];
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
  return { id: 0, row: 0, col: 0, text: "", tableId: 0 };
}

export const Cell = {
  encode(message: Cell, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.row !== 0) {
      writer.uint32(16).int64(message.row);
    }
    if (message.col !== 0) {
      writer.uint32(24).int64(message.col);
    }
    if (message.text !== "") {
      writer.uint32(34).string(message.text);
    }
    if (message.tableId !== 0) {
      writer.uint32(40).int64(message.tableId);
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
          if (tag != 8) {
            break;
          }

          message.id = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.row = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.col = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.text = reader.string();
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.tableId = longToNumber(reader.int64() as Long);
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
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      row: isSet(object.row) ? Number(object.row) : 0,
      col: isSet(object.col) ? Number(object.col) : 0,
      text: isSet(object.text) ? String(object.text) : "",
      tableId: isSet(object.tableId) ? Number(object.tableId) : 0,
    };
  },

  toJSON(message: Cell): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.row !== undefined && (obj.row = Math.round(message.row));
    message.col !== undefined && (obj.col = Math.round(message.col));
    message.text !== undefined && (obj.text = message.text);
    message.tableId !== undefined && (obj.tableId = Math.round(message.tableId));
    return obj;
  },

  create<I extends Exact<DeepPartial<Cell>, I>>(base?: I): Cell {
    return Cell.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Cell>, I>>(object: I): Cell {
    const message = createBaseCell();
    message.id = object.id ?? 0;
    message.row = object.row ?? 0;
    message.col = object.col ?? 0;
    message.text = object.text ?? "";
    message.tableId = object.tableId ?? 0;
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
  /** CreateFile */
  CreateFile(request: CreateFileRequest): Promise<CreateFileResponse>;
  /** ReadFile */
  ReadFile(request: ReadFileRequest): Promise<ReadFileResponse>;
  /** UpdateFile */
  UpdateFile(request: UpdateFileRequest): Promise<UpdateFileResponse>;
  /** DeleteFile */
  DeleteFile(request: DeleteFileRequest): Promise<DeleteFileResponse>;
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
    this.CreateFile = this.CreateFile.bind(this);
    this.ReadFile = this.ReadFile.bind(this);
    this.UpdateFile = this.UpdateFile.bind(this);
    this.DeleteFile = this.DeleteFile.bind(this);
    this.CreateTable = this.CreateTable.bind(this);
    this.ReadTable = this.ReadTable.bind(this);
    this.UpdateTable = this.UpdateTable.bind(this);
    this.DeleteTable = this.DeleteTable.bind(this);
    this.CreateCell = this.CreateCell.bind(this);
    this.ReadCell = this.ReadCell.bind(this);
    this.UpdateCell = this.UpdateCell.bind(this);
    this.DeleteCell = this.DeleteCell.bind(this);
  }
  CreateFile(request: CreateFileRequest): Promise<CreateFileResponse> {
    const data = CreateFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateFile", data);
    return promise.then((data) => CreateFileResponse.decode(_m0.Reader.create(data)));
  }

  ReadFile(request: ReadFileRequest): Promise<ReadFileResponse> {
    const data = ReadFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ReadFile", data);
    return promise.then((data) => ReadFileResponse.decode(_m0.Reader.create(data)));
  }

  UpdateFile(request: UpdateFileRequest): Promise<UpdateFileResponse> {
    const data = UpdateFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateFile", data);
    return promise.then((data) => UpdateFileResponse.decode(_m0.Reader.create(data)));
  }

  DeleteFile(request: DeleteFileRequest): Promise<DeleteFileResponse> {
    const data = DeleteFileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteFile", data);
    return promise.then((data) => DeleteFileResponse.decode(_m0.Reader.create(data)));
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
