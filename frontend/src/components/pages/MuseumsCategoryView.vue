<template>
  <div>
    <!-- Cabeçalho -->
    <HeaderPage></HeaderPage>

    <!-- Conteúdo principal -->
    <div class="page container-sm px-5 my-5">
      <h2 class="title-h2 mb-4" style="text-transform: capitalize;">Categoria: {{ category }}</h2>

      <!-- Visitados com o endpoint correto -->
      <Visitados :urlAPI="`museums?category_like=${category}`"></Visitados>
    </div>

    <!-- Rodapé -->
    <Footer></Footer>
  </div>
</template>

<script>
import HeaderPage from "../Header.vue";
import Footer from "../Footer.vue";
import Visitados from "./Visitados.vue";

export default {
  components: {
    HeaderPage,
    Footer,
    Visitados,
  },
  data() {
    return {
      category: this.$route.params.category,
    };
  },
  watch: {
    "$route.params.category"(newVal) {
      this.category = newVal;
    },
  },
};
</script>

<style scoped>
/* Ajuste de estilos caso necessário */
</style>
