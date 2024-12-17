<template>
    <div>
      <!-- Cabeçalho -->
      <HeaderPage></HeaderPage>
  
      <!-- Conteúdo principal -->
      <div class="page container-sm px-5 my-5">
        <h2 class="title-h2 mb-4">Museus</h2>
  
        <!-- Visitados com o endpoint 'museums' para listar todos -->
        <Visitados :urlAPI="'museums'"></Visitados>
      </div>
  
      <!-- Rodapé -->
      <Footer></Footer>
    </div>
  </template>
  
  <script>
  import HeaderPage from "../HeaderPage.vue";
  import Footer from "../Footer.vue";
  import Visitados from "../Visitados.vue";
  
  export default {
    components: {
      HeaderPage,
      Footer,
      Visitados
    }
  };
  </script>
  
  <style scoped>
  /* Ajuste de estilos caso necessário */
  </style>
  