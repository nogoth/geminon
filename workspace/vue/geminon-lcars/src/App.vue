<template>
  <div class="lcars-container">
    <!-- Panel 1: LOC Metrics -->
    <div class="lcars-panel purple">
      <div class="lcars-elbow purple"></div>
      <div class="lcars-content">
        <h2>LOC Metrics</h2>
        <div class="scanning-text">Scanning...</div>
      </div>
    </div>

    <!-- Panel 2: Tech Hub -->
    <div class="lcars-panel blue">
      <div class="lcars-elbow blue"></div>
      <div class="lcars-content">
        <h2>Tech Stack</h2>
        <nav>
          <ul>
            <li><a href="https://go.dev" target="_blank">Go</a></li>
            <li><a href="https://python.org" target="_blank">Python</a></li>
            <li><a href="https://sqlite.org" target="_blank">SQLite</a></li>
            <li><a href="https://vuejs.org" target="_blank">Vue</a></li>
            <li><a href="https://docs.astral.sh/uv/" target="_blank">uv</a></li>
          </ul>
        </nav>
      </div>
    </div>

    <!-- Panel 3: Temporal Data -->
    <div class="lcars-panel orange">
      <div class="lcars-elbow orange"></div>
      <div class="lcars-content">
        <h2>Temporal Data</h2>
        <div class="stardate">Stardate: 1.000</div>
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
}
a:hover {
  text-decoration: underline;
}
</style>
