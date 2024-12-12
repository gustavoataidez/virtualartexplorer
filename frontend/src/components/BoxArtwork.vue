<template>
    <div>
      <div class="capa-obra" :style="{ backgroundImage: `url(${artwork.image || '../assets/obras/obra-1.jpg'})` }">
        <div class="column1">
          <h3>{{ artwork.name }}</h3>
          <p>{{ artwork.author || 'Autor Desconhecido' }}</p>
        </div>
        
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "BoxArtwork",
    props: {
      artwork: {
        type: Object,
        required: true,
      },
    },
  };
  </script>
  
  <style scoped>
  .capa-obra {
    font-family: "Poppins";
    background-size: cover;
    background-color: rgb(0, 0, 0, 0.5);
    display: flex;
    flex-direction: row;
    width: 250px;
    height: 250px;
    align-items: center;
    padding: 10px;
    padding-top: 180px;
    justify-content: space-around;
    border-radius: 20px;
    position: relative;
  }
  .detalhes {
    width: min-content;
    height: min-content;
  }
  h3 {
    color: white;
    font-size: 1.3rem;
    line-height: 1.2rem;
  }
  p {
    color: white;
    font-size: 0.8rem;
    line-height: 1rem;
  }
  .detalhes {
    background-color: var(--vt-c-brow);
    border: 0px;
    border-radius: 100%;
    color: white;
    font-size: 1rem;
    padding: 0px;
    width: 40px;
    height: 40px;
  }
  .detalhes:hover {
    background-color: var(--vt-c-indigo);
    transition: 0.3s;
  }
  </style>
  