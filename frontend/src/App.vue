<script lang="ts" setup>
import CredentialInput from './components/CredentialInput.vue';
import Navigation from './components/Navigation.vue';
import { useMetaStore } from '@stores/metaStore';
import { storeToRefs } from 'pinia';

const store = useMetaStore();
const toggleDarkMode = store.toggleDarkMode;
const { isDarkMode } = storeToRefs(store);

</script>
<template>
  <div class="min-h-screen" :class="isDarkMode ? 'dark' : ''">
    <Navigation v-on:toggleDarkMode="toggleDarkMode" />
    <div class="xl:ml-52 mt-12 h-full z-20 border-r-1 relative">
      <!-- <RouterView /> -->
      <CredentialInput />
    </div>
  </div>
</template>
