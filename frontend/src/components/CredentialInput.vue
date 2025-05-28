<script setup lang="ts">
import { ref } from 'vue';
import { toast } from 'vue3-toastify';
import { Connect } from '../../wailsjs/go/protocols/ConnectionManager';
import { types } from '../../wailsjs/go/models';

interface Credential {
  protocol: 'ftp' | 'sftp';
  host: string;
  port: number;
  username: string;
  password: string;
}

const credential = ref<Credential>({
  protocol: 'sftp',
  host: '',
  port: 22,
  username: '',
  password: ''
});

const connectString = ref<string>('');
const connectData = ref<string>('');

const submitForm = () => {
  console.log('Submitted Credential:', credential.value);
  // Here you can handle the form submission, e.g., send the data to an API
  toast.success('Credential submitted successfully!', {
    autoClose: 3000,
    position: 'top-right',
  });

  Connect().then((result: types.ConnectResult) => {
    console.log("CONNECTION TYPE", result.Type);
    console.log("IS CONNECTED", result.IsConnected);
    console.log("LIST DIR", result.Data);
    connectString.value = `Connected to ${result.Type} at ${credential.value.host}:${credential.value.port}`;
    connectData.value = result.Data;
    toast.success(`Connected successfully! Type: ${result.Type}`, {
      autoClose: 5000,
      position: 'top-right',
    });
  }).catch((error: Error) => {
    toast.error(`Error connecting: ${error.message}`, {
      autoClose: 5000,
      position: 'top-right',
    });
  });
};

</script>
<template>
  <h1 class="text-white text-3xl text-center">Credential Input Component</h1>
  <p>{{ connectString }}</p>
  <p>{{ connectData }}</p>
  <!-- Form with: Select (FTP/SFTP), Host, Port, Username, Password -->
  <form class="max-w-md mx-auto mt-8 bg-white p-6 rounded-lg shadow-md" @submit.prevent="() => submitForm()">
    <div class="mb-4">
      <label for="protocol" class="block text-gray-700">Protocol</label>
      <select id="protocol" class="w-full p-2 border border-gray-300 rounded" v-model="credential.protocol">
        <option value="ftp">FTP</option>
        <option value="sftp">SFTP</option>
      </select>
    </div>
    <div class="mb-4">
      <label for="host" class="block text-gray-700">Host</label>
      <input type="text" id="host" class="w-full p-2 border border-gray-300 rounded" placeholder="Enter host"
        v-model="credential.host">
    </div>
    <div class="mb-4">
      <label for="port" class="block text-gray-700">Port</label>
      <input type="number" id="port" class="w-full p-2 border border-gray-300 rounded" placeholder="Enter port"
        v-model="credential.port">
    </div>
    <div class="mb-4">
      <label for="username" class="block text-gray-700">Username</label>
      <input type="text" id="username" class="w-full p-2 border border-gray-300 rounded" placeholder="Enter username"
        v-model="credential.username">
    </div>
    <div class="mb-4">
      <label for="password" class="block text-gray-700">Password</label>
      <input type="password" id="password" class="w-full p-2 border border-gray-300 rounded"
        placeholder="Enter password" v-model="credential.password">
    </div>
    <button type="submit" class="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600">Submit</button>
  </form>
</template>
