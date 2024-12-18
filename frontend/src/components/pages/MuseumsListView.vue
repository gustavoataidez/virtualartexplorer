<template>
    <div>
      <HeaderPage></HeaderPage>
  
      <div class="page container-sm px-5 my-5">
        <h2 class="title-h2 mb-4">Museus</h2>
  
        <Visitados :urlAPI="'museums'"></Visitados>
      </div>
  
      <Footer></Footer>
    </div>
  </template>
  
  <script>
  import HeaderPage from "../Header.vue";
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
  