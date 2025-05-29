export type ProtocolType = "sftp" | "ftp";

// Neat way to allow spread operator and iteration over ProtocolTypes :)
export const ProtocolTypes = {
  SFTP: "sftp" as ProtocolType,
  FTP: "ftp" as ProtocolType,

  [Symbol.iterator]() {
    return Object.values(ProtocolTypes)[Symbol.iterator]();
  },
};

export interface SailsError {
  operation: string;
  reason: string;
  code: number;
}
