<template>
  <div class="mt-5 container p-0" style="width: 100%;">
    <span class="subtitle" > Tópicos para você {{ urlAPI }} </span>
    <h2 class="title-h2">Mais visitados </h2>
    <div class="grid">
      <div
      v-for="( item,index ,) in items" :key="item.id"
        class="box-1"
      >
        <div class="card-container position-relative">
          <div class="img">
            <img
              :src="item.strDrinkThumb"
              :alt="'Imagem ' + (index + 1)"
              class="card-img-top"
            />
          </div>
          <div class="card-overlay d-flex d-row align-items-center justify-content-center">
            <div class="content">
              <h5 class="card-title">{{ item.strDrink }}</h5>
              <span class="text-location">{{ item.idDrink }}</span><span>{{ item.idDrink }}</span>

            </div>
            <div class="">
              <img class="icon" src="@/assets/Icons.png" alt="Icons" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row justify-content-center mt-5">
      <button type="button" class="btn btn-warning w-25 mb-4">
        Mais Museus
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" fill="white" width="20px"><path d="M438.6 278.6c12.5-12.5 12.5-32.8 0-45.3l-160-160c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L338.8 224 32 224c-17.7 0-32 14.3-32 32s14.3 32 32 32l306.7 0L233.4 393.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0l160-160z"/></svg>

      </button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      items: null,
      drinkDetails: null,
      idDrinks: [],
      currentPage: 1,
      itemsPerPage: 8,
      totalPages: 0
  }},
  watch: {
    urlAPI: {
      immediate: true, // Para carregar o conteúdo na primeira vez
      handler(newUrl) {
        this.fetchData(newUrl);
      }
    }
  },
  mounted(){
  this.fetchData();
  //this.fetchDrinkDetails(this.items.idDrink);
},
props: {
urlAPI: String,
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
    //const url3 = `https://potterapi-fedeperin.vercel.app/pt/books?max=${this.itemsPerPage}`
    //const url2 = `https://virtualartexplorer.site/api/v1/museums`
    const urlBase = `https://www.thecocktaildb.com/api/json/v1/1/`
    const url = `${urlBase}${this.urlAPI}`

    const urlDrinkId = `${urlBase}lookup.php?i=${this.idDrink}`;
    try {
      const response = await axios.get(url); 
      this.items = response.data.drinks.slice(0, this.itemsPerPage);
      const teste = [];
      this.items.map(element => {
        teste.push(element.idDrink)
        this.idDrinks.push(element.idDrink)
      })
      console.log(teste)
      console.log(this.idDrinks)
      console.log("-------------1")











      /*
      this.idDrink = response.data.drinks.push(this.items.idDrink)
      console.log(urlDrinkId)*/
    } catch (e) {
      console.error('Error fetching data:', e);
    }
  },
    async fetchDrinkDetails(idDrink) {
      const urlBase = `https://www.thecocktaildb.com/api/json/v1/1/`;
      const urlDrinkId = `${urlBase}lookup.php?i=${idDrink}`;

      try {
        const response = await axios.get(urlDrinkId);
        console.log(urlDrinkId)
        this.drinkDetails = response.data.drinks[0];
        console.log("passou aqui")
      } catch (e) {
        console.error('Error fetching drink details:', e);
      }
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
}
}
</script>

<style scoped>
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
  padding: 10px;
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
    left: 0%;
  }
  .grid{
    justify-content: center;
  }

}
</style>
