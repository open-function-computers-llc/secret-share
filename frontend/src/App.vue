<template>
<div id="app">
  <div class="flex justify-content-between align-items-center">
    <h1>
      <a href="/">Super Secure Secret Share</a>
    </h1>
    <div class="logo-wrapper">
      <img alt="Open Function Computers" src="./assets/ofco-logo.svg" />
    </div>
  </div>

  <router-view />
</div>
</template>

<script>
export default {
  name: "App",
  components: {},
  mounted() {
    document.title = "Share a secret";
  },
};
</script>

<style>
/* global styles go here in the root app */

:root {
  --bg: #1e1e1e;
  --bg-2: #181818;
  --panel: #2f3238;
  --panel-2: #3b3f49;
  --text: #d4d4d4;
  --dim: #9d9d9d;
  --faint: #8a8a8a;
  --border: rgba(255, 255, 255, 0.09);
  --border-acc: rgba(144, 144, 245, 0.4);
  --brand: #5858e8;
  --brand-2: #7b7bf2;
  --brand-soft: rgba(88, 88, 232, 0.28);
  --acc: #9090f5;
  --acc-soft: rgba(144, 144, 245, 0.16);
  --ok: #27c93f;
  --dot-r: #ff5f56;
  --dot-y: #ffbd2e;
  --dot-g: #27c93f;
  --radius: 9px;
  --font-display: "Space Grotesk", ui-sans-serif, system-ui, -apple-system,
    "Segoe UI", Roboto, sans-serif;
  --font-mono: "JetBrains Mono", ui-monospace, SFMono-Regular, Menlo, Consolas,
    monospace;
}

html,
body {
  min-height: 100vh;
  margin: 0;
}

body {
  background-color: var(--bg);
  background-image: radial-gradient(1100px 600px at 12% -10%,
      rgba(88, 88, 232, 0.18),
      transparent 60%),
    radial-gradient(900px 520px at 100% 0%,
      rgba(144, 144, 245, 0.12),
      transparent 55%),
    linear-gradient(var(--bg), var(--bg-2));
  background-attachment: fixed;
  color: var(--text);
  font-family: var(--font-display);
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#app {
  text-align: center;
  max-width: 1100px;
  width: min(92%, 1100px);
  margin: 0 auto;
  padding: 28px 0 80px;
}

h1 {
  font-weight: 700;
  letter-spacing: -0.02em;
  font-size: clamp(1.6rem, 4vw, 2.3rem);
  margin: 0.2em 0 0.35em;
}

h1 a {
  color: #fff;
  text-decoration: none;
}

h2 {
  font-weight: 600;
  letter-spacing: -0.01em;
  font-size: clamp(1.3rem, 3.5vw, 1.8rem);
  color: var(--text);
}

p {
  color: var(--dim);
  line-height: 1.6;
}

/* shared button (single source of truth) */
.button {
  display: inline-block;
  font-family: var(--font-display);
  font-weight: 600;
  font-size: 1.1rem;
  line-height: 1.2;
  color: #fff;
  background: var(--brand);
  border: 1px solid transparent;
  border-radius: var(--radius);
  padding: 0.6rem 1.25rem;
  text-decoration: none;
  cursor: pointer;
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.35);
  transition: transform 0.15s ease, box-shadow 0.22s ease,
    background-color 0.22s ease, color 0.22s ease, border-color 0.22s ease;
}

.button:hover {
  background: var(--brand-2);
  box-shadow: 0 10px 30px var(--brand-soft);
  transform: translateY(-1px);
}

.button:active {
  transform: translateY(0);
  box-shadow: 0 4px 14px var(--brand-soft);
}

.button:focus-visible {
  outline: none;
  box-shadow: 0 0 0 3px var(--acc-soft), 0 10px 30px var(--brand-soft);
}

li.active .button {
  background: var(--brand);
  color: #fff;
  border-color: #fff;
  box-shadow: 0 0 18px rgba(123, 123, 242, 0.7);
}

/* copy-to-clipboard button (shared) */
.button--copy {
  position: relative;
}

.button--copy .sizer {
  display: block;
  height: 0;
  overflow: hidden;
  visibility: hidden;
  pointer-events: none;
}

.button--copy .label {
  display: block;
}

.button.copied {
  background: var(--ok);
  border-color: var(--ok);
  color: #06210d;
  box-shadow: 0 10px 30px rgba(39, 201, 63, 0.35);
  animation: pop 0.3s ease;
}

@keyframes pop {
  0% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.08);
  }

  100% {
    transform: scale(1);
  }
}

/* form controls */
textarea,
input[type="number"],
select {
  font-family: var(--font-mono);
  font-size: 1rem;
  color: var(--text);
  background: var(--panel);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 0.7rem 0.9rem;
  width: 100%;
  box-sizing: border-box;
  transition: border-color 0.2s ease, box-shadow 0.2s ease,
    background-color 0.2s ease;
}

textarea {
  height: 140px;
  resize: vertical;
  text-align: left;
  line-height: 1.5;
}

input[type="number"] {
  width: 130px;
  text-align: center;
}

select {
  width: auto;
  min-width: 190px;
  cursor: pointer;
  -webkit-appearance: none;
  appearance: none;
  padding-right: 2.4rem;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%239d9d9d' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 0.8rem center;
  background-size: 1em;
}

textarea:focus,
input[type="number"]:focus,
select:focus {
  outline: none;
  border-color: var(--border-acc);
  box-shadow: 0 0 0 3px var(--acc-soft);
  background: var(--panel-2);
}

textarea::placeholder {
  color: var(--faint);
}

option {
  background: var(--panel);
  color: var(--text);
}

/* secret output */
pre {
  font-family: var(--font-mono);
  text-align: left;
  color: var(--text);
  background: var(--panel);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 1.1rem 1.2rem;
  width: 100%;
  box-sizing: border-box;
  margin: 0 0 1rem;
  overflow-x: auto;
  white-space: pre-wrap;
  word-break: break-word;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.35);
}

/* layout utilities used by templates */
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
  width: 76px;
  height: 76px;
  padding: 12px;
  box-sizing: border-box;
  background: var(--brand);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 12px var(--brand-soft);
}

img {
  width: 100%;
  height: auto;
  display: block;
}
</style>
