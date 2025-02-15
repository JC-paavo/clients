import {createApp} from 'vue'
import App from './App.vue'

import router from './renderer/router'
import ElementUI from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import './renderer/assets/fonts/font-awesome.min.css';

const app = createApp(App)

app.use(ElementUI)
app.use(router)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.mount('#app')
