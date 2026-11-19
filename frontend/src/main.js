import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// the deployment's theme (THEME_CSS_PATH) is added after the app's own styles
// so its rules win over the built-in ones
const theme = document.createElement('link')
theme.rel = 'stylesheet'
theme.href = '/theme.css'
document.head.appendChild(theme)

createApp(App).use(router).mount('#app')
