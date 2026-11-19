<template>
<div id="app">
  <div class="flex justify-content-between align-items-center">
    <h1>
      <a href="/">{{ theme.title }}</a>
    </h1>
    <div v-if="theme.logo" class="logo-wrapper">
      <img :alt="theme.logoAlt" :src="theme.logo" />
    </div>
  </div>

  <router-view />
</div>
</template>

<script>
// all branding comes from the server's env (see server/theme.go)
export default {
  name: "App",
  data() {
    return {
      theme: {},
    };
  },
  async mounted() {
    const response = await fetch("/api/theme");
    this.theme = await response.json();

    document.title = this.theme.title;
    if (this.theme.themeColor) {
      document
        .querySelector('meta[name="theme-color"]')
        .setAttribute("content", this.theme.themeColor);
    }
  },
};
</script>

<style>
/* global styles go here in the root app */

/* colors, fonts and sizes are CSS variables defined by the THEME_CSS_PATH
   stylesheet (see theme/default.css) */

html,
body {
  min-height: 100vh;
  margin: 0;
}

body {
  background-color: var(--bg);
  background-image: radial-gradient(1100px 600px at 12% -10%,
      color-mix(in srgb, var(--brand) 18%, transparent),
      transparent 60%),
    radial-gradient(900px 520px at 100% 0%,
      color-mix(in srgb, var(--acc) 12%, transparent),
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
  color: var(--heading);
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
  color: var(--on-brand);
  background: var(--brand);
  border: 1px solid transparent;
  border-radius: var(--radius);
  padding: 0.6rem 1.25rem;
  text-decoration: none;
  cursor: pointer;
  box-shadow: 0 6px 18px var(--shadow);
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
  color: var(--on-brand);
  border-color: var(--on-brand);
  box-shadow: 0 0 18px color-mix(in srgb, var(--brand-2) 70%, transparent);
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
  color: var(--on-ok);
  box-shadow: 0 10px 30px color-mix(in srgb, var(--ok) 35%, transparent);
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
  /* chevron drawn with two gradients so it follows the theme's --dim color */
  background-image: linear-gradient(45deg, transparent 50%, var(--dim) 50%),
    linear-gradient(135deg, var(--dim) 50%, transparent 50%);
  background-repeat: no-repeat;
  background-position: right 1.25rem center, right 0.9rem center;
  background-size: 0.35rem 0.35rem;
}

textarea:focus,
input[type="number"]:focus,
select:focus {
  outline: none;
  border-color: var(--border-acc);
  box-shadow: 0 0 0 3px var(--acc-soft);
  background-color: var(--panel-2);
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
  box-shadow: 0 10px 30px var(--shadow);
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
