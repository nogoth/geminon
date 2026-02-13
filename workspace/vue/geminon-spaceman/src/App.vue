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
    <!-- Sidebar Launcher -->
    <aside class="side-launcher">
      <button 
        class="nav-btn purple" 
        :class="{ active: activePanel === 'metrics' }"
        @click="togglePanel('metrics')"
        title="LOC METRICS"
      ></button>
      <button 
        class="nav-btn blue" 
        :class="{ active: activePanel === 'tech' }"
        @click="togglePanel('tech')"
        title="TECH HUB"
      ></button>
      <button 
        class="nav-btn orange" 
        :class="{ active: activePanel === 'temporal' }"
        @click="togglePanel('temporal')"
        title="TEMPORAL DATA"
      ></button>
    </aside>

    <!-- Main Console Stage -->
    <main class="console-stage">
      <Transition name="panel-slide">
        <!-- Panel 1: LOC Metrics -->
        <div v-if="activePanel === 'metrics'" class="spaceman-panel purple">
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
      </Transition>

      <Transition name="panel-slide">
        <!-- Panel 2: Tech Hub -->
        <div v-if="activePanel === 'tech'" class="spaceman-panel blue">
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
      </Transition>

      <Transition name="panel-slide">
        <!-- Panel 3: Temporal Data -->
        <div v-if="activePanel === 'temporal'" class="spaceman-panel orange">
          <div class="spaceman-elbow orange"></div>
          <div class="spaceman-content">
            <h2>Temporal Data</h2>
            <div class="stardate-container">
              <div class="label">STARDATE</div>
              <div class="value">{{ stardate.toFixed(3) }}</div>
            </div>
          </div>
        </div>
      </Transition>
    </main>
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

/* Panel Slide Transitions */
.panel-slide-enter-active,
.panel-slide-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.panel-slide-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.panel-slide-leave-to {
  opacity: 0;
  transform: translateX(-30px);
  position: absolute;
}

@media (max-width: 768px) {
  .value {
    font-size: 1.5rem;
  }
}
</style>
