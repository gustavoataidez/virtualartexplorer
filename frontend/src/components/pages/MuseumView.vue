<template>
  <div>
    <Header></Header>
    <div class="museu-page" v-if="museum">
      <div class="sec-resume">
        <div 
          class="capa-museu" 
          :style="{ backgroundImage: `url(${museum.image})` }"
        ></div>
        <div class="resume-desc">
          <div class="d-flex flex-wrap">
  <div v-for="(category, index) in museum.category" :key="index" class="mb-2 mx-1">
    <span 
      v-if="category" 
      class="badge rounded-pill text-bg-warning text-capitalize fs-6"
    >{{ category }}
    </span>
  </div>
</div>
          <h1>{{ museum.title }}</h1>
          <p>{{ museum.description }}</p>
          <h6 style="font-weight: 600; font-size: 1rem;">Mais informações</h6>
          <p style="line-height: 1.5rem;">{{ museum.information }}</p>
          <div id="endereco">
            <p style="line-height: 1.5rem; margin-bottom: 0; font-weight: 500; color: var(--vt-c-brown);">
              {{ museum.address }}<br>
              <span id="city-museum" class="address-page">{{ museum.city }}</span> - 
              <span id="state-museum" class="address-page">{{ museum.state }}</span>
            </p>
          </div>
        </div>
      </div>
      <div class="sec-obras">
        <span style="font-weight: 600; color: var(--vt-c-brown); line-height: 1rem;">Explore mais</span>
        <h2 style="line-height: 1.5rem; font-weight: 600;">Obras do Museu</h2>
        <div class="grade-obras">
          <!-- Exibe as obras se existirem, caso contrário, exibe a mensagem -->
          <div v-if="artworks.length === 0">
            <p style="font-weight: 400; color: var(--vt-c-brown); font-size:1rem; opacity: 0.9;">Nenhuma obra cadastrada para este museu.</p>
          </div>
          <BoxArtwork
            v-for="artwork in artworks"
            :key="artwork.id"
            :artwork="artwork"
          />
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "../Header.vue";
import Footer from "../Footer.vue";
import BoxArtwork from "../BoxArtwork.vue";
import axios from "axios";
import { API_URL } from "@/config";

export default {
  components: { BoxArtwork, Header, Footer },
  data() {
    return {
      museum: null, // Detalhes do museu
      artworks: [], // Lista de obras associadas
    };
  },
  async created() {
    const museumId = this.$route.params.id; // Obtém o ID do museu da rota
    await this.fetchMuseum(museumId); // Busca informações do museu
    await this.fetchArtworks(museumId); // Busca obras associadas ao museu
  },
  methods: {
    async fetchMuseum(id) {
      try {
        const response = await axios.get(`${API_URL}/museums/${id}`);
        console.log("Dados do museu recebidos:", response.data); // Log para depuração
        this.museum = response.data || null; // Armazena diretamente os dados do museu
      } catch (error) {
        console.error("Erro ao buscar o museu:", error);
        this.museum = null; // Define como null em caso de erro
      }
    },
    async fetchArtworks(museumId) {
      try {
        const response = await axios.get(`${API_URL}/artworks/museum/id/${museumId}`);
        this.artworks = response.data || []; // Trata diretamente como lista
      } catch (error) {
        console.error("Erro ao buscar as obras:", error);
        this.artworks = []; // Define uma lista vazia em caso de erro
      }
    },
  },
};
</script>



    
  <style>
  .address-page {
    color: var(--vt-c-brown);
  }
  .museu-page {
    font-family: 'Poppins';
    width: 90vw;
    margin-left: 50%;
    transform: translate(-50%, 0);
    padding: 50px;
    max-width: 1200px;
    color: var(--vt-c-brown);
    display: flex;
    flex-direction: column;
    gap: 50px;
  }
  .sec-resume {
    display: flex;
    flex-direction: row;
    justify-content: start;
    align-items: center;
    gap: 50px;
  }
  .resume-desc {
    width: 50%;
  }
  .resume-desc p {
    font-size: 1.1rem;
  }
  .resume-desc h1{
    font-size: 3rem;
    font-weight: 600;
    margin-bottom: 20px;
  }
  .capa-museu {
    width: 50%;
    height: 400px;
    border-radius: 20px;
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
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
  