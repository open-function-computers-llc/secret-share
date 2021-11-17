<template>
  <div id="app">
    <div class="flex justify-content-between align-items-center">
      <h1>Super Secure Secret Share</h1>
      <div class="logo-wrapper">
        <img alt="Open Function" src="./assets/logo.png">
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {
  },
  mounted() {
    document.title = "Share a secret"
    this.$api.get("test")
  }
}
</script>

<style>
body, html {
  min-height: 100vh;
  margin: 0;
  background: #252525;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  max-width: 80%;
  margin: 0 auto;
  width: 1200px;
}

h1 {
  color: white;
}

.flex {
  display: flex;
}

.justify-content-between {
  justify-content: space-between;
}

.align-items-center {
  align-items: center;
}

.logo-wrapper {
  width: 100px;
}

img {
  width: 100%;
}
</style>
