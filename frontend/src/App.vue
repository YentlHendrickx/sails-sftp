<script lang="ts" setup>
import CredentialInput from './components/CredentialInput.vue';
</script>
<template>
  <div class="min-h-screen">
    <!-- <Navigation v-if="showNavigation" :standardHide="hideNavigation" /> -->
    <div class="xl:ml-52 mt-12 h-full z-20 border-r-1 relative">
      <!-- <RouterView /> -->
      <CredentialInput />
    </div>
  </div>
</template>
