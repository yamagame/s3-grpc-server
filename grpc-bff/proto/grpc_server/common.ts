/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export enum Result {
  UNDEFINED = 0,
  OK = 1,
  ERR = 2,
  UNRECOGNIZED = -1,
}

export function resultFromJSON(object: any): Result {
  switch (object) {
    case 0:
    case "UNDEFINED":
      return Result.UNDEFINED;
    case 1:
    case "OK":
      return Result.OK;
    case 2:
    case "ERR":
      return Result.ERR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Result.UNRECOGNIZED;
  }
}

export function resultToJSON(object: Result): string {
  switch (object) {
    case Result.UNDEFINED:
      return "UNDEFINED";
    case Result.OK:
      return "OK";
    case Result.ERR:
      return "ERR";
    case Result.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface File {
  id: number;
  filename: string;
  tables: Table[];
}

export interface Table {
  id: number;
  fileId: number;
  title: string;
  cells: Cell[];
}

export interface Cell {
  id: number;
  row: number;
  col: number;
  text: string;
  tableId: number;
}

function createBaseFile(): File {
  return { id: 0, filename: "", tables: [] };
}

export const File = {
  encode(message: File, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.filename = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
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

function createBaseTable(): Table {
  return { id: 0, fileId: 0, title: "", cells: [] };
}

export const Table = {
  encode(message: Table, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.fileId !== 0) {
      writer.uint32(16).uint64(message.fileId);
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
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.fileId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.title = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.cells.push(Cell.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
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

function createBaseCell(): Cell {
  return { id: 0, row: 0, col: 0, text: "", tableId: 0 };
}

export const Cell = {
  encode(message: Cell, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.row !== 0) {
      writer.uint32(16).uint64(message.row);
    }
    if (message.col !== 0) {
      writer.uint32(24).uint64(message.col);
    }
    if (message.text !== "") {
      writer.uint32(34).string(message.text);
    }
    if (message.tableId !== 0) {
      writer.uint32(40).uint64(message.tableId);
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
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.row = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.col = longToNumber(reader.uint64() as Long);
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.text = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.tableId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
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
