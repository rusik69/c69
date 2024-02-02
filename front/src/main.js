import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueVnc from "vue-vnc";


const app = createApp(App)

app.use(router)
app.use(VueVnc);

app.mount('#app')
