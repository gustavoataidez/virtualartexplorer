<template>
  <div>
    <HeaderPage></HeaderPage>

    <div class="page container-sm px-5 my-5">
      <h2 class="title-h2 mb-4">Obras de Arte</h2>
      
      <div class="grade-obras">
        <!-- Renderiza as obras de arte -->
        <BoxArtwork
          v-for="artwork in artworks"
          :key="artwork.id"
          :artwork="artwork"
        />
      </div>
    </div>

    <Footer></Footer>
  </div>
</template>

<script>
import HeaderPage from "../Header.vue";
import Footer from "../Footer.vue";
import BoxArtwork from "../BoxArtwork.vue";
import axios from "axios";
import { API_URL } from "@/config";

export default {
  components: {
    HeaderPage,
    Footer,
    BoxArtwork,
  },
  data() {
    return {
      artworks: [], // Armazena as obras de arte
    };
  },
  async created() {
    await this.fetchArtworks(); // Busca as obras ao carregar a página
  },
  methods: {
    async fetchArtworks() {
      try {
        const response = await axios.get(`${API_URL}/artworks`);
        // Adapta os dados retornados pela API para o formato esperado
        this.artworks = response.data.map((artwork) => ({
          id: artwork.ID, // Mapeia `ID` para `id`
          museumId: artwork.museum_id,
          name: artwork.name,
          description: artwork.description,
          image: artwork.image,
          author: artwork.author,
          active: artwork.active,
        }));
      } catch (error) {
        console.error("Erro ao buscar as obras de arte:", error);
      }
    },
  },
};
</script>

  
  <style scoped>
  .page {
    font-family: 'Poppins';
    color: var(--vt-c-brown);
  }
  
  .grade-obras {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: center;
    gap: 30px;
    margin: 40px 0;
  }
  </style>
  