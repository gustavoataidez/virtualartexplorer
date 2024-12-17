<template>
    <Header></Header>
    <div class="museu-page" v-if="museum">
        <div class="sec-resume">
            <div 
  class="capa-museu" 
  :style="{ backgroundImage: `url(${museum.image})` }"
>
</div>

            <div class="resume-desc">
                <span id="category-museum" class="badge rounded-pill text-bg-warning mb-2">
                    {{ museum.category1 }}, {{ museum.category2 }}
                </span>
                <h1>{{ museum.title }}</h1>
                <p>{{ museum.description }}</p>
                <h6 style="font-weight: 600; font-size: 1rem;">Mais informações</h6>
                <p style="line-height: 1.5rem;">{{ museum.information }}</p>
                <div id="endereco" class="d-flex flex-row align-items-center">
                    <img style="width: 60px; margin-right: 10px;" src="../assets/location.png" alt="Localização">
                    <p style="line-height: 1.5rem; margin-bottom: 0rem; font-weight: 500; color: var(--vt-c-brown);">
                        {{ museum.address }}<br>
                        <span id="city-museum" class="address-page">{{ museum.city }}</span> -
                        <span id="state-museum" class="address-page">{{ museum.state }}</span>
                    </p>
                </div>
            </div>
        </div>
        <div class="sec-obras">
            <span style="font-weight: 600; color: var(--vt-c-brown); line-height: 1rem;">Explore mais</span>
            <h2 style="line-height: 1.5rem;">Obras do Museu</h2>
            <div class="grade-obras">
                <BoxArtwork
                    v-for="artwork in artworks"
                    :key="artwork.id"
                    :artwork="artwork"
                />
            </div>
        </div>
    </div>
    <Footer></Footer>
</template>

<script>
import Header from "../HeaderPage.vue";
import Footer from "../Footer.vue";
import BoxArtwork from "../BoxArtwork.vue";
import axios from "axios";
import { API_URL } from '@/config'; // Importe a constante da configuração

export default {
    components: { BoxArtwork, Header, Footer },
    data() {
        return {
            museum: null, // Dados do museu
            artworks: [] // Obras do museu
        };
    },
    async created() {
        const museumId = this.$route.params.id; // Obtém o ID do museu da URL
        await this.fetchMuseum(museumId);
        await this.fetchArtworks(museumId);
    },
    methods: {
        async fetchMuseum(id) {
            try {
                const response = await axios.get(`${API_URL}/museums/${id}`);
                this.museum = response.data;
            } catch (error) {
                console.error("Erro ao buscar o museu:", error);
            }
        },
        async fetchArtworks(museumId) {
            try {
                const response = await axios.get(`${API_URL}/artworks?museum_id=${museumId}`);
                this.artworks = response.data;
            } catch (error) {
                console.error("Erro ao buscar as obras:", error);
            }
        }
    }
};
</script>

<style>
.address-page{
    color: var(--vt-c-brown);
}
.museu-page{
    font-family: 'Poppins';
    width: 90vw;
    margin-left: 50%;
    transform: translate(-50%,0);
    padding: 50px;
    max-width: 1200px;
    color: var(--vt-c-brown);
    display: flex;
    flex-direction: column;
    gap: 50px;
}
.sec-resume{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    gap: 50px;
}
.resume-desc{
    width: 50%;
}
.resume-desc p{
    font-size: 0.9rem;
}
.capa-museu {
    width: 50%;
    height: 400px;
    border-radius: 50px;
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
}

.grade-obras{
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: center;
    gap: 30px;
    margin: 40px 0;
}
</style>