<template>
  <div class="mt-5 container p-0" style="width: 100%;">
    <h2 class="title-h2">Mais visitados</h2>
    <div class="grid">
      <div v-for="(museum, index) in items" :key="museum.id" class="box-1">
        <div class="card-container position-relative">
          <div class="img">
            <img
              :src="museum.image"
              :alt="'Imagem do ' + museum.title"
              class="card-img-top"
            />
          </div>
          <div class="card-overlay d-flex flex-column align-items-center justify-content-center">
            <div class="content">
              <h5 class="card-title">{{ truncateTitle(museum.title) }}</h5>
              <span class="text-location">{{ museum.city }}, {{ museum.state }}</span>
              <p class="text-description">{{ museum.category1	}}, {{ museum.category2	}}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      items: [], // Lista de museus
      currentPage: 1, // Página atual para paginação
      itemsPerPage: 8 // Número de itens por página
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    truncateTitle(title) {
      if (title.length > 20) {
        return title.substring(0, 20) + "...";
      }
      return title;
    },
    async fetchData() {
      const url = `http://localhost:3000/museums`;

      try {
        const response = await axios.get(url);
        const startIndex = (this.currentPage - 1) * this.itemsPerPage;
        const endIndex = startIndex + this.itemsPerPage;
        this.items = response.data.slice(startIndex, endIndex);
      } catch (e) {
        console.error("Error fetching data:", e);
      }
    }
  }
};
</script>


<style scoped>
.text-description{
  margin-bottom: 0rem;
  color: var(--vt-c-orange);
  font-weight: 500;
  font-size: 0.9rem;
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
  font-size: 1rem;
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
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .card-container{
    width: 100%;
  }
  .grid{
    justify-content: center;
  }

}
</style>
