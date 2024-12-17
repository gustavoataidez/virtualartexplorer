<template>
  <HeaderPage></HeaderPage>

  <div class="image-container">
    <h1 class="image-text">
      Tenha uma incrível <br />
      <span>experiência</span> aqui!
    </h1>
  </div>

  <div class="page container-sm px-5 my-5">
    <span class="subtitle">Mais procuradas</span>
    <h2 class="title-h2 mb-4">Categorias</h2>

    <div class="d-flex flex-row justify-content-between flex-wrap gap-1">
      <div id="categ-esportes" class="category" @click="goToCategory('Esportes')">
        <h5 class="categ-title">Esportes</h5>
      </div>
      <div id="categ-pessoas" class="category" @click="goToCategory('Pessoas')">
        <h5 class="categ-title">Pessoas</h5>
      </div>
      <div id="categ-escravidao" class="category" @click="goToCategory('Escravidão')">
        <h5 class="categ-title">Escravidão</h5>
      </div>
      <div id="categ-cultura" class="category" @click="goToCategory('Cultura')">
        <h5 class="categ-title">Cultura</h5>
      </div>
    </div>

    <div class="mt-5 container p-0" style="width: 100%;">
      <span class="subtitle">Conheça os museus mais procurados</span>
      <h2 class="title-h2 mb-0">Mais visitados</h2>
    </div>

    <Visitados urlAPI="museums"></Visitados>
  </div>

  <Footer></Footer>
</template>

<script>
import Visitados from "../Visitados.vue";
import Footer from "../Footer.vue";
import HeaderPage from "../HeaderPage.vue";

export default {
  components: { Visitados, Footer, HeaderPage },
  data() {
    return {
    valor: "",
    images: [
        "../../src/assets/categorias/categ-esportes.jpg",
        "../../src/assets/categorias/categ-pessoas.jpg",
        "../../src/assets/categorias/categ-escravidao.jpg",
        "../../src/assets/categorias/categ-cultura.jpeg",
      ],

    };
  },
  methods: {
    goToCategory(category) {
      this.$router.push(`/museums/${category}`);
    }
  },
  
};
</script>

<style>
.categ-title {
  font-family: "Poppins";
  color: white;
  font-size: 1.5rem;
  text-align: center;
}
.category {
  background-size: cover;
  background-position: center;
  height: 100px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 23%;
  min-width: 200px;
}
.category:hover {
  cursor: pointer;
  transform: scale(1.02);
  transition: 0.5s;
}
.image-container {
  padding-top: 0px;
  background-image: url("../../assets/categorias/capa-museu2.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 400px;
  gap: 20px;
}
#categ-esportes {
  background-image: url("../../assets/categorias/categ-esportes.jpg");
}
#categ-pessoas {
  background-image: url("../../assets/categorias/categ-pessoas.jpg");
}
#categ-escravidao {
  background-image: url("../../assets/categorias/categ-escravidao.jpg");
}
#categ-cultura {
  background-image: url("../../assets/categorias/categ-cultura2.jpg");
}

.card-tag {
  background-repeat: no-repeat;
  background-size: cover;
}
.button-search {
  border: none;
  background-color: rgba(0, 0, 0, 0);
}
.input-div {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  border: none;
  border-radius: 15px;
  background-color: #e9e9e9;
  max-width: 700px;
  width: 90%;
}
.input-overlay {
  position: relative;
  padding: 18px 10px;
  text-align: start;
  border: none;
  width: 100%;
  max-width: 700px;
  font-size: 1.1rem;
  font-family: "Poppins", sans-serif;
  background-color: rgba(0, 0, 0, 0);
  width: 90%;
  color: #000;
  font-weight: 500;
}

.image-text {
  position: relative;
  color: white;
  text-align: center;
  font-size: 3.5rem;
  font-weight: 500;
  font-family: "Poppins", sans-serif;
}

span {
  font-weight: 600;
  font-family: "Poppins", sans-serif;
  color: var(--vt-c-orange);
}

.search-icon {
  position: relative;
}

img {
  width: 100%;
}
textarea:focus,
input:focus,
select:focus {
  box-shadow: 0 0 0 0;
  border: 0 none;
  outline: 0;
}

@media screen and (max-width: 768px) {
  .image-text {
    font-size: 2.5rem;
  }
  .category {
    width: 100%;
    height: 70px;
    border-radius: 10px;
    margin-top: 5px;
  }
}
</style>
