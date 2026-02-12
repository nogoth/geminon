<script setup>
import TicketGrid from './components/TicketGrid.vue'
</script>

<template>
  <div class="container">
    <header>
      <h1>Ticket Tracker</h1>
    </header>
    
    <main>
      <TicketGrid />
    </main>
  </div>
</template>

<style scoped>
.container {
  min-height: 100vh;
  padding: 2rem;
}

header {
  margin-bottom: 3rem;
  text-align: center;
}

h1 {
  font-size: 3rem;
  margin: 0;
  background: linear-gradient(to right, #646cff, #42b883);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>
