<script setup>
import { reactive, ref } from 'vue'
import axios from 'axios'

const emit = defineEmits(['close', 'created'])

const form = reactive({
  title: '',
  description: '',
  area_of_concern: '',
  agent_name: '',
  status: 'Open'
})

const submitting = ref(false)
const error = ref(null)

const submitForm = async () => {
  if (!form.title) {
    error.value = 'Title is required'
    return
  }
  
  try {
    submitting.value = true
    error.value = null
    const host = window.location.hostname
    await axios.post(`http://${host}:8080/tickets`, form)
    emit('created')
    emit('close')
  } catch (err) {
    console.error('Failed to create ticket:', err)
    error.value = 'Failed to create ticket. Please try again.'
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div class="form-overlay" @click.self="$emit('close')">
    <div class="form-card">
      <header>
        <h2>Create New Ticket</h2>
        <button class="close-btn" @click="$emit('close')">&times;</button>
      </header>

      <form @submit.prevent="submitForm">
        <div class="field">
          <label>Title</label>
          <input v-model="form.title" placeholder="Summary of the issue" required />
        </div>

        <div class="field">
          <label>Description</label>
          <textarea v-model="form.description" placeholder="Detailed explanation..." rows="4"></textarea>
        </div>

        <div class="row">
          <div class="field">
            <label>Area of Concern</label>
            <input v-model="form.area_of_concern" placeholder="e.g. IT, HR" />
          </div>
          <div class="field">
            <label>Agent Name</label>
            <input v-model="form.agent_name" placeholder="Assigned to" />
          </div>
        </div>

        <div v-if="error" class="error-msg">{{ error }}</div>

        <div class="actions">
          <button type="button" class="cancel-btn" @click="$emit('close')">Cancel</button>
          <button type="submit" class="submit-btn" :disabled="submitting">
            {{ submitting ? 'Creating...' : 'Create Ticket' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
.form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  padding: 1rem;
}

.form-card {
  background: #2a2a2a;
  border: 1px solid #444;
  border-radius: 16px;
  width: 100%;
  max-width: 600px;
  padding: 2rem;
  box-shadow: 0 20px 40px rgba(0,0,0,0.5);
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

h2 {
  margin: 0;
  font-size: 1.5rem;
}

.close-btn {
  background: transparent;
  border: none;
  font-size: 2rem;
  color: #888;
  cursor: pointer;
  padding: 0;
  line-height: 1;
}

form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

label {
  font-size: 0.9rem;
  font-weight: bold;
  color: #888;
}

input, textarea {
  background: #1a1a1a;
  border: 1px solid #444;
  border-radius: 8px;
  padding: 0.8rem;
  color: white;
  font-family: inherit;
}

input:focus, textarea:focus {
  outline: none;
  border-color: #646cff;
}

.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.error-msg {
  color: #f44336;
  font-size: 0.9rem;
}

.actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1rem;
}

button {
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
  font-family: inherit;
}

.cancel-btn {
  background: transparent;
  border: 1px solid #444;
  color: #888;
}

.submit-btn {
  background: #646cff;
  border: none;
  color: white;
}

.submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

@media (prefers-color-scheme: light) {
  .form-card { background: white; border-color: #eee; }
  input, textarea { background: #f9f9f9; border-color: #ddd; color: #222; }
  .cancel-btn { border-color: #ddd; }
}
</style>
