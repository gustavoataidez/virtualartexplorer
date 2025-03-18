<template>
    <HeaderPage></HeaderPage>
  
    <div class="page container-sm px-5 my-5">
      <div class="mt-5 container p-0 mb-0" style="width: 100%;">
        <span class="subtitle">Resultados da Busca:</span>
        <h2 class="title-h2 mb-0">{{ searchQuery }}</h2>
      </div>
  
      <!-- Passar os museus encontrados para o componente Visitados -->
      <Visitados :museus="museus" v-if="museus.length > 0"></Visitados>
  
      <!-- Mensagem caso nenhum museu seja encontrado -->
      <div v-else>
        <p>Nenhum museu encontrado para "{{ searchQuery }}".</p>
      </div>
    </div>
  
    <Footer></Footer>
  </template>
  
  <script>
  import Visitados from "./Visitados.vue";
  import Footer from "../Footer.vue";
  import HeaderPage from "../Header.vue";
  
  export default {
    components: { Visitados, Footer, HeaderPage },
    data() {
      return {
        searchQuery: "",
        museus: [],
      };
    },
    created() {
      this.searchQuery = this.$route.query.q || "";
      this.searchMuseus();
    },
    methods: {
      searchMuseus() {
        if (this.searchQuery.trim() !== "") {
          fetch(`http://localhost:3000/museums?q=${this.searchQuery}`)
            .then((response) => response.json())
            .then((data) => {
              this.museus = data;
            })
            .catch((error) => {
              console.error("Erro ao buscar museus:", error);
            });
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Adicione estilos específicos para a página de resultados de busca aqui */
  </style>