<template>
  <div class="capa-obra" :style="{ backgroundImage: `url(${artwork.image || '../assets/obras/obra-1.jpg'})` }">
    <div class="column1">
      <h3>{{ artwork.name }}</h3>
      <p>{{ artwork.author || 'Autor Desconhecido' }}</p>
      <button class="btn btn-primary" @click.stop="openModal">Mais Detalhes</button>
    </div>

    <div v-if="showModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <h3>{{ artwork.name }}</h3>
        <p>Autor: {{ artwork.author || 'Autor Desconhecido' }}</p>
        <p>{{ artwork.description }}</p>
        <p>Museu ID: {{ artwork.museum_id }}</p>
        <img :src="artwork.image || '../assets/obras/obra-1.jpg'" alt="Obra" style="max-width:100%; margin-top:10px;"/>
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
  align-items: center;
  padding: 10px;
  padding-top: 180px;
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

/* Modal */
.modal-overlay {
  position: fixed;
  top: 0;
  left:0;
  right:0;
  bottom:0;
  background-color: rgba(0,0,0,0.5);
  display:flex;
  justify-content:center;
  align-items:center;
  z-index:9999;
}
.modal-content {
  background:#fff;
  padding:20px;
  border-radius:8px;
  max-width:400px;
  width:100%;
  text-align:center;
}
.modal-content h3 {
  color:#333;
  font-size:1.5rem;
  line-height:2rem;
}
.modal-content p {
  color:#333;
  font-size:1rem;
  line-height:1.2rem;
  margin:10px 0;
}
</style>
