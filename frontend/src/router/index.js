import { createWebHistory, createRouter } from 'vue-router';

import Home from '@/components/pages/Home.vue';
import MuseumsListView from '@/components/pages/MuseumsListView.vue';
import SearchPage from '@/components/SearchPage.vue';
import MuseumPage from '@/components/pages/MuseumView.vue';
import MuseumEdit from '@/components/pages/MuseumEditView.vue';
import MuseumCreate from '@/components/pages/MuseumCreateView.vue';
import ArtworkPage from '@/components/ArtworkPage.vue';
import LoginTest from '@/components/LoginTest.vue';

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/search', component: SearchPage },
    { path: '/museums', component: MuseumsListView },
    { 
      path: '/museum/:id', 
      component: MuseumPage, 
      props: true // Passa o ID como prop para o componente
    },
    { 
      path: '/museum/edit/:id', 
      component: MuseumEdit, 
      props: true // Passa o ID como prop para o componente
    },
    { path: '/museum/create', component: MuseumCreate },
    { path: '/artwork', component: ArtworkPage }
  ]
});
