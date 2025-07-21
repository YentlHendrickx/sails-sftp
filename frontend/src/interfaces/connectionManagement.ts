import { types } from "../../wailsjs/go/models";

export interface ConnectionManagement {
  connections: {
    active: types.ConnectResult[]; // Only includes currently connected connections
    inactive: types.ConnectResult[]; // Only includes currently disconnected connections
  };
}
