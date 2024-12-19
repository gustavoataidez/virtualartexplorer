import { createWebHistory, createRouter } from 'vue-router';

import Home from '@/components/pages/Home.vue';
import MuseumsListView from '@/components/pages/MuseumsListView.vue';
import SearchPage from '@/components/SearchPage.vue';
import MuseumPage from '@/components/pages/MuseumView.vue';
import MuseumEdit from '@/components/pages/MuseumEditView.vue';
import MuseumCreate from '@/components/pages/MuseumCreateView.vue';
import ArtworkPage from '@/components/ArtworkPage.vue';
import ArtworkView from '@/components/pages/ArtworksView.vue';
import MuseumCategoryView from '@/components/pages/MuseumsCategoryView.vue'; // Importe o novo componente

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/search', component: SearchPage },
    { path: '/museums', component: MuseumsListView },
    { 
      path: '/museums/:category',
      component: MuseumCategoryView,
      props: true
    },
    { 
      path: '/museum/:id', 
      component: MuseumPage, 
      props: true
    },
    { 
      path: '/museum/edit/:id', 
      component: MuseumEdit, 
      props: true
    },
    { path: '/museum/create', component: MuseumCreate },
    { path: '/artworks', component: ArtworkView }
  ]
});
