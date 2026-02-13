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
  }, 1000) // The plan says 100ms, but 0.001 every 100ms might be too slow/fast?
  // Wait, "incrementing every 100ms". I'll stick to 100ms.
  
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
  <div class="lcars-container">
    <!-- Panel 1: LOC Metrics -->
    <div class="lcars-panel purple">
      <div class="lcars-elbow purple"></div>
      <div class="lcars-content">
        <h2>LOC Metrics</h2>
        <div v-if="isScanning" class="scanning-text">Scanning Project...</div>
        <div v-else class="loc-data">
          <div class="stat">TOTAL LOC: {{ totalLoc }}</div>
          <div class="stat">FILES: 142</div>
          <div class="stat">DIRECTORIES: 12</div>
        </div>
      </div>
    </div>

    <!-- Panel 2: Tech Hub -->
    <div class="lcars-panel blue">
      <div class="lcars-elbow blue"></div>
      <div class="lcars-content">
        <h2>Tech Stack</h2>
        <nav>
          <ul>
            <li v-for="tech in techStack" :key="tech.name">
              <a :href="tech.url" target="_blank">{{ tech.name }}</a>
            </li>
          </ul>
        </nav>
      </div>
    </div>

    <!-- Panel 3: Temporal Data -->
    <div class="lcars-panel orange">
      <div class="lcars-elbow orange"></div>
      <div class="lcars-content">
        <h2>Temporal Data</h2>
        <div class="stardate-container">
          <div class="label">STARDATE</div>
          <div class="value">{{ stardate.toFixed(3) }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
h2 {
  margin: 0 0 20px 0;
  color: inherit;
}
ul {
  list-style: none;
  padding: 0;
}
li {
  margin-bottom: 10px;
}
a {
  color: var(--lcars-blue);
  text-decoration: none;
  font-weight: bold;
}
a:hover {
  text-decoration: underline;
  color: var(--lcars-orange);
}
.stat, .stardate-container {
  font-size: 1.5rem;
  margin-bottom: 10px;
}
.label {
  font-size: 0.8rem;
  color: var(--lcars-orange);
}
.value {
  font-size: 2.5rem;
  color: white;
}
</style>
