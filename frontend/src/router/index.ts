import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import TroubleshootingView from '../views/TroubleshootingView.vue'
import CorrosionAIView from '../views/CorrosionAIView.vue'
import SafetyAdvisorView from '../views/SafetyAdvisorView.vue'
import VCRAView from '../views/VCRAView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/troubleshooting',
      name: 'troubleshooting',
      component: TroubleshootingView
    },
    {
      path: '/corrosionai',
      name: 'corrosion-ai',
      component: CorrosionAIView,
    },
    {
      path: '/safetyadvisor',
      name: 'safety-advisor',
      component: SafetyAdvisorView
    },
    {
      path: '/vcra',
      name: 'vcra',
      component: VCRAView
    },
  ],
})

export default router
