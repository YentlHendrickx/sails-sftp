<script lang="ts" setup>
import CredentialInput from './components/CredentialInput.vue';
import Navigation from './components/Navigation.vue';
import { useMetaStore } from '@stores/metaStore';
import { storeToRefs } from 'pinia';
import { RouterView } from 'vue-router';

const store = useMetaStore();
const toggleDarkMode = store.toggleDarkMode;
const { isDarkMode } = storeToRefs(store);

// Make sure t

</script>
<template>
  <div class="min-h-screen" :class="isDarkMode ? 'dark' : ''">
    <Navigation v-on:toggleDarkMode="toggleDarkMode" />
    <main class="min-h-screen">
      <RouterView />
    </main>
  </div>
</template>
