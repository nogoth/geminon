<script setup>
import { ref, onMounted } from 'vue'

const stardate = ref(1.000)
const totalLoc = ref(0)
const isScanning = ref(true)
const activePanel = ref('metrics') // Default to metrics

const techStack = [
  { name: 'Go', url: 'https://go.dev' },
  { name: 'Python', url: 'https://python.org' },
  { name: 'SQLite', url: 'https://sqlite.org' },
  { name: 'Vue', url: 'https://vuejs.org' },
  { name: 'uv', url: 'https://docs.astral.sh/uv/' }
]

const togglePanel = (panel) => {
  activePanel.value = activePanel.value === panel ? null : panel
}

onMounted(() => {
  setInterval(() => {
    stardate.value = parseFloat((stardate.value + 0.001).toFixed(3))
  }, 100)

  setTimeout(() => {
    isScanning.value = false
    totalLoc.value = 12402
  }, 3000)
})
</script>

<template>
  <div class="spaceman-interface">
    <!-- Left Sidebar: Primary Controls -->
    <aside class="side-launcher left">
      <button 
        class="nav-btn action" 
        :class="{ active: activePanel === 'metrics' }"
        @click="togglePanel('metrics')"
        title="METRICS"
      ></button>
      <button 
        class="nav-btn variation" 
        :class="{ active: activePanel === 'tech' }"
        @click="togglePanel('tech')"
        title="TECH HUB"
      ></button>
      <button 
        class="nav-btn primary" 
        :class="{ active: activePanel === 'temporal' }"
        @click="togglePanel('temporal')"
        title="TEMPORAL"
      ></button>
    </aside>

    <!-- Main Console Stage -->
    <main class="console-stage">
      <Transition name="panel-slide">
        <!-- Panel 1: LOC Metrics -->
        <div v-if="activePanel === 'metrics'" class="spaceman-panel action">
          <div class="spaceman-elbow action"></div>
          <div class="spaceman-content">
            <h2>CODE METRICS</h2>
            <div v-if="isScanning" class="scanning-text">INITIALIZING SURVEY...</div>
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
      </Transition>

      <Transition name="panel-slide">
        <!-- Panel 2: Tech Hub -->
        <div v-if="activePanel === 'tech'" class="spaceman-panel variation">
          <div class="spaceman-elbow variation"></div>
          <div class="spaceman-content">
            <h2>LOGISTICS HUB</h2>
            <nav>
              <ul>
                <li v-for="tech in techStack" :key="tech.name">
                  <a :href="tech.url" target="_blank">{{ tech.name }}</a>
                </li>
              </ul>
            </nav>
          </div>
        </div>
      </Transition>

      <Transition name="panel-slide">
        <!-- Panel 3: Temporal Data -->
        <div v-if="activePanel === 'temporal'" class="spaceman-panel primary">
          <div class="spaceman-elbow primary"></div>
          <div class="spaceman-content">
            <h2>CHRONO DATA</h2>
            <div class="stardate-container">
              <div class="label">STARDATE</div>
              <div class="value">{{ stardate.toFixed(3) }}</div>
            </div>
          </div>
        </div>
      </Transition>
    </main>

    <!-- Right Sidebar: Utility Controls -->
    <aside class="side-launcher right">
      <button class="nav-btn action" title="UTILITY ALPHA"></button>
      <button class="nav-btn variation" title="COMMAND BETA"></button>
      <button class="nav-btn primary" title="SYSTEM GAMMA"></button>
    </aside>
  </div>
</template>

<style scoped>
.label {
  display: block;
  font-size: 0.9rem;
  margin-bottom: 5px;
}
.value {
  font-family: 'Courier New', Courier, monospace;
}

/* Panel Slide Transitions */
.panel-slide-enter-active,
.panel-slide-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.panel-slide-enter-from {
  opacity: 0;
  transform: scale(0.95) translateY(20px);
}

.panel-slide-leave-to {
  opacity: 0;
  transform: scale(1.05) translateY(-20px);
  position: absolute;
}

@media (max-width: 768px) {
  .value {
    font-size: 1.8rem;
  }
}
</style>
