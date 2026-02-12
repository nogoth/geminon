<script setup>
import { ref } from 'vue'
import TicketGrid from './components/TicketGrid.vue'
import NewTicketForm from './components/NewTicketForm.vue'

const gridRef = ref(null)
const isCreating = ref(false)

const handleTicketCreated = () => {
  if (gridRef.value) {
    gridRef.value.fetchTickets()
  }
}
</script>

<template>
  <div class="container">
    <header :class="{ 'faded': isCreating }">
      <h1>Ticket Tracker</h1>
    </header>
    
    <main :class="{ 'faded': isCreating }">
      <TicketGrid ref="gridRef" />
    </main>

    <button 
      class="fab" 
      @click="isCreating = true"
      v-if="!isCreating"
      title="Create New Ticket"
    >
      +
    </button>

    <Transition name="fade">
      <NewTicketForm 
        v-if="isCreating" 
        @close="isCreating = false"
        @created="handleTicketCreated"
      />
    </Transition>
  </div>
</template>

<style scoped>
.container {
  min-height: 100vh;
  padding: 2rem;
  position: relative;
}

header {
  margin-bottom: 3rem;
  text-align: center;
  transition: opacity 0.3s ease, filter 0.3s ease;
}

main {
  transition: opacity 0.3s ease, filter 0.3s ease;
}

.faded {
  opacity: 0.2;
  filter: blur(4px);
  pointer-events: none;
}

h1 {
  font-size: 3rem;
  margin: 0;
  background: linear-gradient(to right, #646cff, #42b883);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.fab {
  position: fixed;
  top: 2rem;
  right: 2rem;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: #646cff;
  color: white;
  border: none;
  font-size: 2.5rem;
  font-weight: 300;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 12px rgba(100, 108, 255, 0.4);
  transition: transform 0.2s, background 0.2s;
  z-index: 100;
}

.fab:hover {
  transform: scale(1.1);
  background: #535bf2;
}

/* Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
