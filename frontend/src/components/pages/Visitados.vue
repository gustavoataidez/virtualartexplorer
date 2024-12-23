<template>
  <div class="mt-4 container p-0" style="width: 100%;">
    <div class="grid">
      <div
        v-for="(museum, index) in items"
        :key="museum.id"
        class="box-1"
      >
        <div class="card-container position-relative">
          <div class="img" :style="{ backgroundImage: `url(${randomImage()})` }">
          </div>
          <div class="card-overlay d-flex flex-column align-items-center justify-content-center">
            <div class="content">
              <h5 class="card-title">{{ truncateTitle(museum.title) }}</h5>
              <span class="text-location">{{ museum.city }}, {{ museum.state }}</span>
              <p class="text-description">{{ museum.category1 }} | {{ museum.category2 }}</p>
              <button 
                class="btn1 btn-primary mt-2"
                @click="goToMuseum(museum.id)"
              >
                Saber Mais
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { API_URL } from '@/config';

export default {
  props: {
    urlAPI: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      items: [] // Lista de museus
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    truncateTitle(title) {
      if (title.length > 30) {
        return title.substring(0, 30) + '...';
      }
      return title;
    },
    async fetchData() {
  const url = `${API_URL}/${this.urlAPI}`;

  try {
    const response = await axios.get(url);

    if (response.data) {
      if (Array.isArray(response.data.museums)) {
        // Caso a resposta tenha a chave "museums"
        this.items = response.data.museums;
      } else if (Array.isArray(response.data)) {
        // Caso a resposta seja um array simples
        this.items = response.data;
      } else {
        console.error("Formato de dados inesperado:", response.data);
        this.items = [];
      }
    } else {
      console.error("Resposta da API está vazia ou inválida");
      this.items = [];
    }
  } catch (e) {
    console.error("Erro ao buscar dados:", e);
    this.items = [];
  }
},
randomImage() {
      const images = [
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-1.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-2.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-3.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-4.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-5.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-6.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-7.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-10.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-11.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-12.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-13.png',
        'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/museus/ai/museu-14.png'
      ];
      const randomIndex = Math.floor(Math.random() * images.length);
      return images[randomIndex];
    },
    goToMuseum(id) {
      this.$router.push(`/museum/${id}`); 
    }
  }
};
</script>

<style scoped>
.img {
  width: 100%;
  height: 150px;
  background-size: cover;
  background-position: center;
}
/* Estilo opcional para o botão Saber Mais */
.btn1 {
  text-transform: uppercase;
  font-size: 0.85rem;
  padding: 0.5rem 1rem;
  border-radius: 10px;
  background-color: #fff;
  font-family: 'poppins';
  color: var(--vt-c-orange);
  font-size: 16px;
  font-weight: 500;
  border: solid 2px var(--vt-c-orange);
}
.text-description{
  margin-bottom: 0rem;
  color: var(--vt-c-brown);
  font-weight: 500;
  font-size: 0.9rem;
  text-transform: capitalize;
}
.grid{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  flex-direction: row;
}
.box-1{
  width: 25%;
  display: flex;
  justify-content: center;
  max-width: 400px;
  min-width: 200px;
  margin-bottom: 50px;
}
.card-container {
  position: relative;
  width: 90%;
  margin-bottom: 100px;
}

.card-overlay {
  position: absolute;
  margin-top: -50px;
  width: 100%; /* Ajuste conforme necessário */
  background-color: rgba(255, 255, 255, 1);
  padding: 10px 15px;
  border-radius: 16px;
  z-index: 5;
  color: var(--vt-c-brown);
  font-family: 'poppins';
  box-shadow: 0px 0px 50px rgba(0, 0, 0, 0.4);
}

.card-title {
  margin-bottom: 0.4rem;
  font-size: 1.5rem;
  font-weight: 600;
}

.content {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.card-text {
  margin-bottom: 1rem;
}

.fa-heart {
  cursor: pointer;
}

.text-location{
  font-weight: 500;
  font-size: 14px;
  color: var(--vt-c-brown);
  opacity: 0.5;
}

.img {
  width: 100%;
  overflow: hidden;
  border-radius: 15px;
  object-fit: contain;
}

.btn {
  background-color: var(--vt-c-orange);
  font-family: 'poppins';
  font-weight: 600;
  color: #fff;
  font-size: 20px;
  border: none;
  height: 60px;
  max-width: 220px;
}
.btn:hover {
  background-color: var(--vt-c-brown);
  color: #fff;
  transform: scale(1.02);
  transition: 0.3s;
}
.card-img-top {
  width: 100%;
  height: 250px;
  display: block;
}
.icon {
  width: 40px;
  height: auto;
}


@media screen and (max-width: 768px) {
  .box-1 {
    width: 100%;
  }
  .card-overlay{
    width: 100%;
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    left: 0%;
  }
  .card-container{
    width: 100%;
  }
  .card-container{
    width: 100%;
  }
  .grid{
    justify-content: center;
  }

}
</style>
