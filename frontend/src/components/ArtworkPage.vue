<template>
  <div class="background">
    <div class="card">
      <!-- Conteúdo da sua card -->

      <!-- Modal de detalhes da obra -->
      <div v-if="showModal" class="modal">
        <div class="modal-content">
          <h5>{{ selectedArtwork.name }}</h5>
          <p>{{ selectedArtwork.author || 'Autor Desconhecido' }}</p>
          <p>{{ selectedArtwork.description }}</p>
          <button @click="closeModal" class="btn btn-danger">Fechar</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BoxArtwork from './BoxArtwork.vue';

export default {
  name: "NewArtwork",
  components: { BoxArtwork },
  data() {
    return {
      showModal: false,
      selectedArtwork: null,
    };
  },
  methods: {
    openModal(artwork) {
      this.selectedArtwork = artwork;
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
      this.selectedArtwork = null;
    },
  },
};
</script>


<style scoped>
.card {
  position: fixed;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 11;
  width: 40%;
  max-width: 500px;
  border-radius: 20px;
}
.background {
  position: absolute;
  background-color: rgba(0, 0, 0, 0.1);
  width: 100%;
  height: 100vh;
  z-index: 10;
}
.bg-image.hover-overlay {
  width: 100%;
  height: 100%;
}
.img-fluid {
  border-radius: 20px 20px 0px 0px;
}
.card-title {
  color: black;
  font-weight: bold;
  margin-bottom: 0px;
}
</style>
