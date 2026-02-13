<script setup>
import { ref, onMounted } from 'vue'

const stardate = ref(1.000)
const totalLoc = ref(0)
const isScanning = ref(true)

const techStack = [
  { name: 'Go', url: 'https://go.dev' },
  { name: 'Python', url: 'https://python.org' },
  { name: 'SQLite', url: 'https://sqlite.org' },
  { name: 'Vue', url: 'https://vuejs.org' },
  { name: 'uv', url: 'https://docs.astral.sh/uv/' }
]

onMounted(() => {
  // Stardate counter: increments every 100ms
  setInterval(() => {
    stardate.value = parseFloat((stardate.value + 0.001).toFixed(3))
  }, 100)

  // Simulate scanning for LOC
  setTimeout(() => {
    isScanning.value = false
    totalLoc.value = 12402 // Placeholder for real metrics
  }, 3000)
})
</script>

<template>
  <div class="spaceman-container">
    <!-- Panel 1: LOC Metrics -->
    <div class="spaceman-panel purple">
      <div class="spaceman-elbow purple"></div>
      <div class="spaceman-content">
        <h2>LOC Metrics</h2>
        <div v-if="isScanning" class="scanning-text">SCANNING SECTOR...</div>
        <div v-else class="loc-data">
          <div class="stat">
            <span class="label">TOTAL LOC:</span>
            <span class="value">{{ totalLoc }}</span>
          </div>
          <div class="stat">
            <span class="label">FILES:</span>
            <span class="value">142</span>
          </div>
          <div class="stat">
            <span class="label">DIRECTORIES:</span>
            <span class="value">12</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Panel 2: Tech Hub -->
    <div class="spaceman-panel blue">
      <div class="spaceman-elbow blue"></div>
      <div class="spaceman-content">
        <h2>Tech Hub</h2>
        <nav>
          <ul>
            <li v-for="tech in techStack" :key="tech.name">
              <a :href="tech.url" target="_blank">> {{ tech.name }}</a>
            </li>
          </ul>
        </nav>
      </div>
    </div>

    <!-- Panel 3: Temporal Data -->
    <div class="spaceman-panel orange">
      <div class="spaceman-elbow orange"></div>
      <div class="spaceman-content">
        <h2>Temporal Data</h2>
        <div class="stardate-container">
          <div class="label">STARDATE</div>
          <div class="value">{{ stardate.toFixed(3) }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.label {
  display: block;
  font-size: 0.8rem;
  color: #888;
  margin-bottom: 5px;
}
.value {
  font-size: 2rem;
  font-weight: bold;
}
@media (max-width: 768px) {
  .value {
    font-size: 1.5rem;
  }
}
</style>
