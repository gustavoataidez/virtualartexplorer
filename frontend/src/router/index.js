import { createWebHistory, createRouter } from 'vue-router';

import Home from '@/components/Home.vue';
import SearchPage from '@/components/SearchPage.vue';
import MuseumPage from '@/components/MuseumPage.vue';
import MuseumEdit from '@/components/MuseumEdit.vue';
import MuseumCreate from '@/components/MuseumCreate.vue';
import ArtworkPage from '@/components/ArtworkPage.vue';
import LoginTest from '@/components/LoginTest.vue';

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/search', component: SearchPage },
    { 
      path: '/museum/:id', 
      component: MuseumPage, 
      props: true // Passa o ID como prop para o componente
    },
    { path: '/museum/edit', component: MuseumEdit },
    { path: '/museum/create', component: MuseumCreate },
    { path: '/artwork', component: ArtworkPage }
  ]
});
