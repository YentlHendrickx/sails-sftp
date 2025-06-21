<script lang="ts" setup>
const emit = defineEmits(['toggleDarkMode']);

</script>
<template>
  <nav class="w-full bg-white dark:bg-gray-500 h-12 flex items-center">
    <div class="container mx-auto px-4">
      <h1 class="text-2xl font-bold">sails SFTP</h1>
      <!-- Button to toggle Tailwind CSS dark mode -->
      <button class="ml-auto bg-blue-500 text-white rounded hover:bg-blue-600" @click="() => $emit('toggleDarkMode')">
        Toggle Dark Mode
      </button>
    </div>
  </nav>
</template>
