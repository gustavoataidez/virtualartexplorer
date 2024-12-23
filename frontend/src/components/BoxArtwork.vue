<template>
  <div class="capa-obra" :style="{ backgroundImage: `url(${artwork.image || 'https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/refs/heads/main/frontend/src/assets/obras/ai/img-obra.png'})` }">
    <div class="column1">
      <h3>{{ artwork.name }}</h3>
      <p>{{ artwork.author || 'Autor Desconhecido' }}</p>
      <button class="btn btn-warning" @click.stop="openModal">Mais Detalhes</button>
    </div>

    <div v-if="showModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <h3>{{ artwork.name }}</h3>
        <p>Autor: {{ artwork.author || 'Autor Desconhecido' }}</p>
        <p>{{ artwork.description }}</p>
        <p>Museu ID: {{ artwork.museum_id }}</p>
        <!-- <img :src="artwork.image || '../assets/obras/obra-1.jpg'" alt="Obra" style="max-width:100%; margin-top:10px;"/> -->
        <button class="btn btn-secondary" @click="closeModal" style="margin-top: 10px;">Fechar</button>
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
  data() {
    return {
      showModal: false,
    };
  },
  methods: {
    openModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
  }
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
  align-items: end;
  padding: 15px;
  border-radius: 20px;
  position: relative;
}

h3 {
  color: white;
  font-size: 1.3rem;
  line-height: 1.2rem;
  font-weight: 600;
}
p {
  color: white;
  font-size: 0.8rem;
  line-height: 1rem;
}

.modal-overlay {
  position: fixed;
  display:flex;
  justify-content:center;
  align-items:center;
  z-index:11;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  /* Removido o background-color para não criar overlay */
  /* background-color: rgba(0, 0, 0, 0.5); */
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 20px;
  max-width: 400px;
  width: 100%;
  text-align: center;
  /* Adicionar sombra escura para destacar o modal */
  box-shadow: 0 0 100px rgba(0,0,0,0.9);
}

.modal-content h3 {
  color: #333;
  font-size: 1.5rem;
  line-height: 2rem;
}
.modal-content p {
  color: #333;
  font-size: 1rem;
  line-height: 1.2rem;
  margin: 10px 0;
}
</style>
