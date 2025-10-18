import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import p5vue from "p5vue"
// import anime from 'vue-animejs'

const app = createApp(App)

app.use(router)
app.use(p5vue)
// app.use(anime)

app.mount('#app')
