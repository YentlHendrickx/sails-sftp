<script setup lang="ts">
import { ref, watch } from 'vue';
import { toast } from 'vue3-toastify';

// TODO: I should make some prefixes for the imports, like `go/` or `models/`
import { Connect, Disconnect } from '../../wailsjs/go/protocols/ConnectionManager';
import { types } from '../../wailsjs/go/models';
import type { ProtocolType, SailsError } from 'types/types';
import { ProtocolTypes } from 'types/types';

const lastConnectResult = ref<types.ConnectResult>({} as types.ConnectResult);
const selectedProtocol = ref<ProtocolType>('sftp');
const credentialInput = ref<types.ConnectionInput>({
  type: 'sftp',
  host: '',
  port: 22,
  username: '',
  password: ''
});

watch(selectedProtocol, (newType) => {
  credentialInput.value.type = newType;
});

const submitForm = () => {
  // TODO: Validation logic

  // TODO: This should be in a store, the call to said store should be debounced
  // This component should only be for the input + ui and exposing an event when the form is submitted
  Connect(credentialInput.value).then((result: types.ConnectResult) => {
    toast.success(`Connected successfully! Type: ${result.type}`, {
      autoClose: 5000,
      position: 'top-right',
    });

    lastConnectResult.value = result;
  }).catch((data: any) => {
    // TODO: Write a wrapper for calling Go methods
    const error = JSON.parse(data) as SailsError;

    toast.error(error.reason, {
      autoClose: 5000,
      position: 'top-right',
    });

    lastConnectResult.value = {
      is_connected: false,
    } as types.ConnectResult;
  });
};

// TODO: Again this is not for here, but i really needed to test this all first
const disconnect = () => {
  console.log('Disconnecting...');
  const uuid = lastConnectResult.value.config.uuid;

  // TODO: Build a wrapper around this so that i don't have to catch and json parse every time
  Disconnect(uuid).then(() => {
    toast.success('Disconnected successfully!', {
      autoClose: 3000,
      position: 'top-right',
    });

    lastConnectResult.value.is_connected = false;
  }).catch((data: any) => {
    const error = JSON.parse(data) as SailsError;
    toast.error(`Error disconnecting: ${error.reason}`, {
      autoClose: 5000,
      position: 'top-right',
    });
  });
};
</script>
<template>
  <h1 class="text-white text-3xl text-center">sails SFTP</h1>
  <button @click="disconnect" class="mt-4 bg-red-500 text-white p-2 rounded hover:bg-red-600"
    v-if="lastConnectResult.is_connected">
    Disconnect
  </button>
  <form class="max-w-md mx-auto mt-8 bg-white p-6 rounded-lg shadow-md" @submit.prevent="() => submitForm()">
    <div class="mb-4">
      <label for="protocol" class="block text-gray-700">Protocol</label>
      <select id="protocol" class="w-full p-2 border border-gray-300 rounded" v-model="selectedProtocol">
        <option v-for="type in [...ProtocolTypes]" :value="type">{{ type.toUpperCase() }}</option>
      </select>
    </div>
    <div class="mb-4">
      <label for="host" class="block text-gray-700">Host</label>
      <input type="text" id="host" class="w-full p-2 border border-gray-300 rounded" placeholder="Enter host"
        v-model="credentialInput.host">
    </div>
    <div class="mb-4">
      <label for="port" class="block text-gray-700">Port</label>
      <input type="number" id="port" class="w-full p-2 border border-gray-300 rounded" placeholder="Enter port"
        v-model="credentialInput.port">
    </div>
    <div class="mb-4">
      <label for="username" class="block text-gray-700">Username</label>
      <input type="text" id="username" class="w-full p-2 border border-gray-300 rounded" placeholder="Enter username"
        v-model="credentialInput.username">
    </div>
    <div class="mb-4">
      <label for="password" class="block text-gray-700">Password</label>
      <input type="password" id="password" class="w-full p-2 border border-gray-300 rounded"
        placeholder="Enter password" v-model="credentialInput.password">
    </div>
    <button type="submit" class="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600">Submit</button>
  </form>
</template>
