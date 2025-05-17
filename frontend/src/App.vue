<script setup>
import { ref, onMounted } from 'vue';

const backendMessage = ref('Loading message from backend...');
const errorMessage = ref('');

async function fetchMessage() {
  try {
    // The vite.config.js proxy (see below) will forward this to http://localhost:8080/api/message
    const response = await fetch('/api/hello'); // Using a relative path due to proxy
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    backendMessage.value = data.text;
  } catch (error) {
    console.error("Failed to fetch message:", error);
    errorMessage.value = `Failed to load message: ${error.message}. Make sure the backend is running and CORS is configured if not using a proxy.`;
    backendMessage.value = ''; // Clear loading message
  }
}

onMounted(() => {
  fetchMessage();
});
</script>

<template>
  <div>
    <h1>Vue.js Frontend</h1>
    <p v-if="backendMessage">{{ backendMessage }}</p>
    <p v-if="errorMessage" style="color: red;">{{ errorMessage }}</p>
    <router-view></router-view>
  </div>
</template>

<style scoped>
/* Your styles */
</style>
