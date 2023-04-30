/* eslint-disable */

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
