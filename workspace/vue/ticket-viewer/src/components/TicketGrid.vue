<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import TicketCard from './TicketCard.vue'

const tickets = ref([])
const loading = ref(true)
const error = ref(null)

const fetchTickets = async () => {
  try {
    loading.value = true
    const host = window.location.hostname
    const response = await axios.get(`http://${host}:8080/tickets`)
    tickets.value = response.data
    error.value = null
  } catch (err) {
    console.error('Failed to fetch tickets:', err)
    error.value = 'Failed to load tickets. Is the server running?'
  } finally {
    loading.value = false
  }
}

onMounted(fetchTickets)

defineExpose({ fetchTickets })
</script>

<template>
  <div class="grid-container">
    <div v-if="loading" class="message">Loading tickets...</div>
    <div v-else-if="error" class="message error">{{ error }}</div>
    <div v-else-if="tickets.length === 0" class="message">No tickets found.</div>
    <div v-else class="grid">
      <TicketCard v-for="ticket in tickets" :key="ticket.id" :ticket="ticket" />
    </div>
  </div>
</template>

<style scoped>
.grid-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1.5rem;
  padding: 1rem;
}

.message {
  padding: 4rem;
  font-size: 1.2rem;
  color: #888;
}

.error {
  color: #f44336;
}

@media (max-width: 900px) {
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 600px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
</style>
