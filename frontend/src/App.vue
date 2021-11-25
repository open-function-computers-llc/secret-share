<template>
  <div id="app">
    <div class="flex justify-content-between align-items-center">
      <h1>
        <a href="/">Super Secure Secret Share</a>
      </h1>
      <div class="logo-wrapper">
        <img alt="Open Function" src="./assets/logo.png">
      </div>
    </div>

    <router-view/>
    <p>&nbsp;</p>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {
  },
  mounted() {
    document.title = "Share a secret"
  }
}
</script>

<style>
/* global styles go here in the root app */
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

h1, h2, p {
  color: white;
}

h1 a {
  color: white;
  text-decoration: none;
}

.flex {
  display: flex;
}

.flex-wrap {
  flex-wrap: wrap;
}

.justify-content-center {
  justify-content: center;
}

.justify-content-between {
  justify-content: space-between;
}

.align-items-center {
  align-items: center;
}

.logo-wrapper {
  width: 100px;
  padding-top: 5px;
}

img {
  width: 100%;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

textarea, pre {
  width: 100%;
  border: 0;
  height: 200px;
  background-color: white;
}

</style>
