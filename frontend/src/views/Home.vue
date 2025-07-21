<script setup lang="ts">
import { ref } from 'vue';
import CredentialInput from '@/components/CredentialInput.vue';

const newConnection = ref(false);
</script>
<template>
  <div class="bg-white dark:bg-gray-500 p-2 min-h-screen">
    <h1 class="text-2xl font-bold">Welcome to Sails!</h1>
    <button @click="newConnection = true"
      class="bg-blue-500 hover:bg-blue-600 hover:cursor-pointer p-2 text-white rounded-md">New Connection</button>

    <Teleport to="body">
      <div @click.self="newConnection = false"
        class="fixed top-0 left-0 w-full h-full bg-gray-500/30 flex flex-col gap-2 items-center justify-center"
        v-if="newConnection">
        <CredentialInput />

        <button @click="newConnection = false"
          class="bg-red-500 hover:bg-red-600 hover:cursor-pointer p-2 text-white rounded-md">Cancel</button>
      </div>
    </Teleport>
  </div>
</template>
