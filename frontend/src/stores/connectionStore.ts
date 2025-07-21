import { defineStore } from "pinia";
import { types } from "../../wailsjs/go/models";
import type { ConnectionManagement } from "@/interfaces";

export const useConnectionStore = defineStore("connection", {
  state: (): ConnectionManagement => {
    return {
      connections: {
        active: [],
        inactive: [],
      },
    };
  },
  getters: {
    activeConnections: (state) => state.connections.active,
    inactiveConnections: (state) => state.connections.inactive,
  },
  actions: {
    addNewConnection(connection: types.ConnectResult) {
      if (connection.is_connected) {
        this.addActiveConnection(connection);
      } else {
        this.addInactiveConnection(connection);
      }
    },
    addActiveConnection(connection: types.ConnectResult) {
      if (
        !this.connections.active.some(
          (c) => c.config.uuid === connection.config.uuid,
        )
      ) {
        this.connections.active.push(connection);
      }
    },
    removeActiveConnection(uuid: string) {
      this.connections.active = this.connections.active.filter(
        (c) => c.config.uuid !== uuid,
      );
    },
    addInactiveConnection(connection: types.ConnectResult) {
      if (
        !this.connections.inactive.some(
          (c) => c.config.uuid === connection.config.uuid,
        )
      ) {
        this.connections.inactive.push(connection);
      }
    },
    removeInactiveConnection(uuid: string) {
      this.connections.inactive = this.connections.inactive.filter(
        (c) => c.config.uuid !== uuid,
      );
    },
  },
});
