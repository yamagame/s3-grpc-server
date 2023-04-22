/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export enum result {
  UNDEFINED = 0,
  OK = 1,
  ERR = 2,
  UNRECOGNIZED = -1,
}

export function resultFromJSON(object: any): result {
  switch (object) {
    case 0:
    case "UNDEFINED":
      return result.UNDEFINED;
    case 1:
    case "OK":
      return result.OK;
    case 2:
    case "ERR":
      return result.ERR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return result.UNRECOGNIZED;
  }
}

export function resultToJSON(object: result): string {
  switch (object) {
    case result.UNDEFINED:
      return "UNDEFINED";
    case result.OK:
      return "OK";
    case result.ERR:
      return "ERR";
    case result.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface CreateBucketRequest {
}

export interface CreateBucketResponse {
  result: result;
}

export interface ListBucketsRequest {
}

export interface ListBucketsResponse {
  result: result;
  buckets: ListBucketsResponse_bucket[];
}

export interface ListBucketsResponse_bucket {
  name: string;
}

export interface PutObjectRequest {
  key: string;
  content: string;
}

export interface PutObjectResponse {
  result: result;
  key: string;
}

export interface GetObjectRequest {
  key: string;
}

export interface GetObjectResponse {
  result: result;
  key: string;
  content: string;
}

export interface DeleteObjectRequest {
  key: string;
}

export interface DeleteObjectResponse {
  result: result;
  key: string;
}

export interface ListObjectsRequest {
  prefix: string;
  next?: string | undefined;
}

export interface ListObjectsResponse {
  result: result;
  prefix: string;
  keys: string[];
  next?: string | undefined;
}

function createBaseCreateBucketRequest(): CreateBucketRequest {
  return {};
}

export const CreateBucketRequest = {
  encode(_: CreateBucketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBucketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBucketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CreateBucketRequest {
    return {};
  },

  toJSON(_: CreateBucketRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBucketRequest>, I>>(base?: I): CreateBucketRequest {
    return CreateBucketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateBucketRequest>, I>>(_: I): CreateBucketRequest {
    const message = createBaseCreateBucketRequest();
    return message;
  },
};

function createBaseCreateBucketResponse(): CreateBucketResponse {
  return { result: 0 };
}

export const CreateBucketResponse = {
  encode(message: CreateBucketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== 0) {
      writer.uint32(8).int32(message.result);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBucketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBucketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.result = reader.int32() as any;
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateBucketResponse {
    return { result: isSet(object.result) ? resultFromJSON(object.result) : 0 };
  },

  toJSON(message: CreateBucketResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = resultToJSON(message.result));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBucketResponse>, I>>(base?: I): CreateBucketResponse {
    return CreateBucketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateBucketResponse>, I>>(object: I): CreateBucketResponse {
    const message = createBaseCreateBucketResponse();
    message.result = object.result ?? 0;
    return message;
  },
};

function createBaseListBucketsRequest(): ListBucketsRequest {
  return {};
}

export const ListBucketsRequest = {
  encode(_: ListBucketsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListBucketsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListBucketsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): ListBucketsRequest {
    return {};
  },

  toJSON(_: ListBucketsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListBucketsRequest>, I>>(base?: I): ListBucketsRequest {
    return ListBucketsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListBucketsRequest>, I>>(_: I): ListBucketsRequest {
    const message = createBaseListBucketsRequest();
    return message;
  },
};

function createBaseListBucketsResponse(): ListBucketsResponse {
  return { result: 0, buckets: [] };
}

export const ListBucketsResponse = {
  encode(message: ListBucketsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== 0) {
      writer.uint32(8).int32(message.result);
    }
    for (const v of message.buckets) {
      ListBucketsResponse_bucket.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListBucketsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListBucketsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.result = reader.int32() as any;
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.buckets.push(ListBucketsResponse_bucket.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListBucketsResponse {
    return {
      result: isSet(object.result) ? resultFromJSON(object.result) : 0,
      buckets: Array.isArray(object?.buckets)
        ? object.buckets.map((e: any) => ListBucketsResponse_bucket.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListBucketsResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = resultToJSON(message.result));
    if (message.buckets) {
      obj.buckets = message.buckets.map((e) => e ? ListBucketsResponse_bucket.toJSON(e) : undefined);
    } else {
      obj.buckets = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListBucketsResponse>, I>>(base?: I): ListBucketsResponse {
    return ListBucketsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListBucketsResponse>, I>>(object: I): ListBucketsResponse {
    const message = createBaseListBucketsResponse();
    message.result = object.result ?? 0;
    message.buckets = object.buckets?.map((e) => ListBucketsResponse_bucket.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListBucketsResponse_bucket(): ListBucketsResponse_bucket {
  return { name: "" };
}

export const ListBucketsResponse_bucket = {
  encode(message: ListBucketsResponse_bucket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListBucketsResponse_bucket {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListBucketsResponse_bucket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListBucketsResponse_bucket {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: ListBucketsResponse_bucket): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListBucketsResponse_bucket>, I>>(base?: I): ListBucketsResponse_bucket {
    return ListBucketsResponse_bucket.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListBucketsResponse_bucket>, I>>(object: I): ListBucketsResponse_bucket {
    const message = createBaseListBucketsResponse_bucket();
    message.name = object.name ?? "";
    return message;
  },
};

function createBasePutObjectRequest(): PutObjectRequest {
  return { key: "", content: "" };
}

export const PutObjectRequest = {
  encode(message: PutObjectRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PutObjectRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePutObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.content = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PutObjectRequest {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      content: isSet(object.content) ? String(object.content) : "",
    };
  },

  toJSON(message: PutObjectRequest): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  create<I extends Exact<DeepPartial<PutObjectRequest>, I>>(base?: I): PutObjectRequest {
    return PutObjectRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PutObjectRequest>, I>>(object: I): PutObjectRequest {
    const message = createBasePutObjectRequest();
    message.key = object.key ?? "";
    message.content = object.content ?? "";
    return message;
  },
};

function createBasePutObjectResponse(): PutObjectResponse {
  return { result: 0, key: "" };
}

export const PutObjectResponse = {
  encode(message: PutObjectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== 0) {
      writer.uint32(8).int32(message.result);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PutObjectResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePutObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.result = reader.int32() as any;
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.key = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PutObjectResponse {
    return {
      result: isSet(object.result) ? resultFromJSON(object.result) : 0,
      key: isSet(object.key) ? String(object.key) : "",
    };
  },

  toJSON(message: PutObjectResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = resultToJSON(message.result));
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  create<I extends Exact<DeepPartial<PutObjectResponse>, I>>(base?: I): PutObjectResponse {
    return PutObjectResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PutObjectResponse>, I>>(object: I): PutObjectResponse {
    const message = createBasePutObjectResponse();
    message.result = object.result ?? 0;
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseGetObjectRequest(): GetObjectRequest {
  return { key: "" };
}

export const GetObjectRequest = {
  encode(message: GetObjectRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetObjectRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.key = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetObjectRequest {
    return { key: isSet(object.key) ? String(object.key) : "" };
  },

  toJSON(message: GetObjectRequest): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetObjectRequest>, I>>(base?: I): GetObjectRequest {
    return GetObjectRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetObjectRequest>, I>>(object: I): GetObjectRequest {
    const message = createBaseGetObjectRequest();
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseGetObjectResponse(): GetObjectResponse {
  return { result: 0, key: "", content: "" };
}

export const GetObjectResponse = {
  encode(message: GetObjectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== 0) {
      writer.uint32(8).int32(message.result);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.content !== "") {
      writer.uint32(26).string(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetObjectResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.result = reader.int32() as any;
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.key = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.content = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetObjectResponse {
    return {
      result: isSet(object.result) ? resultFromJSON(object.result) : 0,
      key: isSet(object.key) ? String(object.key) : "",
      content: isSet(object.content) ? String(object.content) : "",
    };
  },

  toJSON(message: GetObjectResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = resultToJSON(message.result));
    message.key !== undefined && (obj.key = message.key);
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetObjectResponse>, I>>(base?: I): GetObjectResponse {
    return GetObjectResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetObjectResponse>, I>>(object: I): GetObjectResponse {
    const message = createBaseGetObjectResponse();
    message.result = object.result ?? 0;
    message.key = object.key ?? "";
    message.content = object.content ?? "";
    return message;
  },
};

function createBaseDeleteObjectRequest(): DeleteObjectRequest {
  return { key: "" };
}

export const DeleteObjectRequest = {
  encode(message: DeleteObjectRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteObjectRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.key = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteObjectRequest {
    return { key: isSet(object.key) ? String(object.key) : "" };
  },

  toJSON(message: DeleteObjectRequest): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteObjectRequest>, I>>(base?: I): DeleteObjectRequest {
    return DeleteObjectRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteObjectRequest>, I>>(object: I): DeleteObjectRequest {
    const message = createBaseDeleteObjectRequest();
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseDeleteObjectResponse(): DeleteObjectResponse {
  return { result: 0, key: "" };
}

export const DeleteObjectResponse = {
  encode(message: DeleteObjectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== 0) {
      writer.uint32(8).int32(message.result);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteObjectResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.result = reader.int32() as any;
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.key = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteObjectResponse {
    return {
      result: isSet(object.result) ? resultFromJSON(object.result) : 0,
      key: isSet(object.key) ? String(object.key) : "",
    };
  },

  toJSON(message: DeleteObjectResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = resultToJSON(message.result));
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteObjectResponse>, I>>(base?: I): DeleteObjectResponse {
    return DeleteObjectResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteObjectResponse>, I>>(object: I): DeleteObjectResponse {
    const message = createBaseDeleteObjectResponse();
    message.result = object.result ?? 0;
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseListObjectsRequest(): ListObjectsRequest {
  return { prefix: "", next: undefined };
}

export const ListObjectsRequest = {
  encode(message: ListObjectsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.prefix !== "") {
      writer.uint32(10).string(message.prefix);
    }
    if (message.next !== undefined) {
      writer.uint32(18).string(message.next);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListObjectsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListObjectsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.prefix = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.next = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListObjectsRequest {
    return {
      prefix: isSet(object.prefix) ? String(object.prefix) : "",
      next: isSet(object.next) ? String(object.next) : undefined,
    };
  },

  toJSON(message: ListObjectsRequest): unknown {
    const obj: any = {};
    message.prefix !== undefined && (obj.prefix = message.prefix);
    message.next !== undefined && (obj.next = message.next);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListObjectsRequest>, I>>(base?: I): ListObjectsRequest {
    return ListObjectsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListObjectsRequest>, I>>(object: I): ListObjectsRequest {
    const message = createBaseListObjectsRequest();
    message.prefix = object.prefix ?? "";
    message.next = object.next ?? undefined;
    return message;
  },
};

function createBaseListObjectsResponse(): ListObjectsResponse {
  return { result: 0, prefix: "", keys: [], next: undefined };
}

export const ListObjectsResponse = {
  encode(message: ListObjectsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result !== 0) {
      writer.uint32(8).int32(message.result);
    }
    if (message.prefix !== "") {
      writer.uint32(18).string(message.prefix);
    }
    for (const v of message.keys) {
      writer.uint32(26).string(v!);
    }
    if (message.next !== undefined) {
      writer.uint32(34).string(message.next);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListObjectsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListObjectsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.result = reader.int32() as any;
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.prefix = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.keys.push(reader.string());
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.next = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListObjectsResponse {
    return {
      result: isSet(object.result) ? resultFromJSON(object.result) : 0,
      prefix: isSet(object.prefix) ? String(object.prefix) : "",
      keys: Array.isArray(object?.keys) ? object.keys.map((e: any) => String(e)) : [],
      next: isSet(object.next) ? String(object.next) : undefined,
    };
  },

  toJSON(message: ListObjectsResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = resultToJSON(message.result));
    message.prefix !== undefined && (obj.prefix = message.prefix);
    if (message.keys) {
      obj.keys = message.keys.map((e) => e);
    } else {
      obj.keys = [];
    }
    message.next !== undefined && (obj.next = message.next);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListObjectsResponse>, I>>(base?: I): ListObjectsResponse {
    return ListObjectsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListObjectsResponse>, I>>(object: I): ListObjectsResponse {
    const message = createBaseListObjectsResponse();
    message.result = object.result ?? 0;
    message.prefix = object.prefix ?? "";
    message.keys = object.keys?.map((e) => e) || [];
    message.next = object.next ?? undefined;
    return message;
  },
};

/** The greeting service definition. */
export interface aws {
  /** CreateBucket */
  CreateBucket(request: CreateBucketRequest): Promise<CreateBucketResponse>;
  /** ListBuckets */
  ListBuckets(request: ListBucketsRequest): Promise<ListBucketsResponse>;
  /** PutObject */
  PutObject(request: PutObjectRequest): Promise<PutObjectResponse>;
  /** GetObject */
  GetObject(request: GetObjectRequest): Promise<GetObjectResponse>;
  /** DeleteObject */
  DeleteObject(request: DeleteObjectRequest): Promise<DeleteObjectResponse>;
  /** ListObjects */
  ListObjects(request: ListObjectsRequest): Promise<ListObjectsResponse>;
}

export class awsClientImpl implements aws {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "aws";
    this.rpc = rpc;
    this.CreateBucket = this.CreateBucket.bind(this);
    this.ListBuckets = this.ListBuckets.bind(this);
    this.PutObject = this.PutObject.bind(this);
    this.GetObject = this.GetObject.bind(this);
    this.DeleteObject = this.DeleteObject.bind(this);
    this.ListObjects = this.ListObjects.bind(this);
  }
  CreateBucket(request: CreateBucketRequest): Promise<CreateBucketResponse> {
    const data = CreateBucketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateBucket", data);
    return promise.then((data) => CreateBucketResponse.decode(_m0.Reader.create(data)));
  }

  ListBuckets(request: ListBucketsRequest): Promise<ListBucketsResponse> {
    const data = ListBucketsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListBuckets", data);
    return promise.then((data) => ListBucketsResponse.decode(_m0.Reader.create(data)));
  }

  PutObject(request: PutObjectRequest): Promise<PutObjectResponse> {
    const data = PutObjectRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PutObject", data);
    return promise.then((data) => PutObjectResponse.decode(_m0.Reader.create(data)));
  }

  GetObject(request: GetObjectRequest): Promise<GetObjectResponse> {
    const data = GetObjectRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetObject", data);
    return promise.then((data) => GetObjectResponse.decode(_m0.Reader.create(data)));
  }

  DeleteObject(request: DeleteObjectRequest): Promise<DeleteObjectResponse> {
    const data = DeleteObjectRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteObject", data);
    return promise.then((data) => DeleteObjectResponse.decode(_m0.Reader.create(data)));
  }

  ListObjects(request: ListObjectsRequest): Promise<ListObjectsResponse> {
    const data = ListObjectsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListObjects", data);
    return promise.then((data) => ListObjectsResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
