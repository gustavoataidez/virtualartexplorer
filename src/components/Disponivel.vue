<template>
  <div class="container mt-5">
  <span class="subtitle" >Explore mais {{ items }}</span>
  <h2 class="title-h2">Algumas obras para você</h2>
  <div class="row">
    <div v-for="(image, index) in images" :key="index" class="col-md-4 mb-4">
      <div class="card-container">
        <div class="card-overlay ">
          <div class="row justify-content-md-center">
            <div class="col-md-3">
              <img
                :src="image"
                :alt="'Imagem ' + (index + 1)"
                class="card-img-top img"
              />
            </div>
            <div class="col-md-9">
              <div class="content">
                <h5 class="card-title">Charles V, bust length...</h5>
                <span>Maceió, Alagoaso</span>
                <strong>29/03/2024 - 01/06/2024</strong>
              </div>
            </div>
          </div>
          <div class="content">
            <FontAwesomeIcon :icon="faHouse" />
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>




<script>
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faHouse } from "@fortawesome/free-solid-svg-icons";

export default {
data() {
  return {
    images: [
      "../src/assets/geovane.png",
      "../src/assets/image2.png",
      "../src/assets/Group95.png",
    ],
    favorites: [], // Array para rastrear quais itens são favoritos,
      items: [],
      currentPage: 1,
      itemsPerPage: 8,
      totalPages: 0
  }},
  mounted(){
  /*
  fetch('https://jsonplaceholder.typicode.com/todos')
  .then(resp => resp.json())
  .then(data => this.todo = data)*/
  this.fetchData();;
},
methods: {
  toggleFavorite(index) {
    if (this.isFavorite(index)) {
      // Se já for favorito, remover da lista de favoritos
      this.favorites.splice(this.favorites.indexOf(index), 1);
    } else {
      // Se não for favorito, adicionar à lista de favoritos
      this.favorites.push(index);
    }
  },
  isFavorite(index) {
    // Verificar se o item está na lista de favoritos
    return this.favorites.includes(index);
  },
  async fetchData() {
      const response = await axios.get('https://jsonplaceholder.typicode.com/todos/'); // Substitua pela sua URL
      this.items = response.data.slice(0, this.itemsPerPage);
      this.totalPages = Math.ceil(response.data.length / this.itemsPerPage);
    },
    loadMore() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        const startIndex = (this.currentPage - 1) * this.itemsPerPage;
        const endIndex = startIndex + this.itemsPerPage;
        this.items  
 = this.items.concat(response.data.slice(startIndex, endIndex));
      }
    }
},
components: {
  FontAwesomeIcon,
},
};
</script>

<style scoped>

.card-overlay {
background-color: rgba(255, 255, 255, 1);
padding: 1rem;
border-radius: 17px;
z-index: 1;
}

.card-title {
margin-bottom: 0.5rem;
}

.content {
display: flex;
flex-direction: column;
}

.card-text {
margin-bottom: 1rem;
}

.img {
width: 100%;
overflow: hidden;
border-radius: 5px;
}
</style>
