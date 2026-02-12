<script setup>
defineProps({
  ticket: {
    type: Object,
    required: true
  }
})

const getStatusColor = (status) => {
  switch (status.toLowerCase()) {
    case 'open': return '#4caf50';
    case 'closed': return '#f44336';
    case 'in progress': return '#ff9800';
    default: return '#9e9e9e';
  }
}
</script>

<template>
  <div class="ticket-card">
    <div class="header">
      <span class="id">#{{ ticket.id }}</span>
      <span class="status" :style="{ backgroundColor: getStatusColor(ticket.status) }">
        {{ ticket.status }}
      </span>
    </div>
    <h3>{{ ticket.title }}</h3>
    <p class="description">{{ ticket.description }}</p>
    <div class="footer">
      <div class="meta">
        <strong>Area:</strong> {{ ticket.area_of_concern }}
      </div>
      <div class="meta">
        <strong>Agent:</strong> {{ ticket.agent_name }}
      </div>
    </div>
  </div>
</template>

<style scoped>
.ticket-card {
  background: #2a2a2a;
  border: 1px solid #444;
  border-radius: 12px;
  padding: 1.5rem;
  text-align: left;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  transition: transform 0.2s, box-shadow 0.2s;
}

.ticket-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0,0,0,0.3);
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.8rem;
}

.id {
  color: #888;
  font-weight: bold;
}

.status {
  padding: 2px 8px;
  border-radius: 4px;
  color: white;
  text-transform: uppercase;
  font-weight: bold;
  font-size: 0.7rem;
}

h3 {
  margin: 0.5rem 0;
  font-size: 1.2rem;
  color: #fff;
}

.description {
  color: #ccc;
  font-size: 0.9rem;
  flex-grow: 1;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.footer {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid #444;
  font-size: 0.85rem;
  color: #aaa;
}

.meta {
  margin-bottom: 0.2rem;
}

@media (prefers-color-scheme: light) {
  .ticket-card {
    background: #fff;
    border-color: #eee;
  }
  h3 { color: #222; }
  .description { color: #555; }
  .footer { border-color: #eee; }
}
</style>
