<template>
  <div class="background">
    <div class="card">
      <div
        class="bg-image hover-overlay"
        data-mdb-ripple-init
        data-mdb-ripple-color="light"
      >
        <img src="../assets/obras/obra-1.jpg" class="img-fluid" />
        <a href="#!">
          <!--<div class="mask" style="background-color: rgba(0, 251, 251, 0.95);"></div>-->
        </a>
      </div>
      <div class="card-body">
        <h5 class="card-title">Nome Obra</h5>
        <p class="">Autor Desconhecido</p>
        <p class="">
          Some quick example text to build on the card title and make up the
          bulk of the card's content.
        </p>
        <router-link to="/" class="btn btn-danger">Fechar </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "NewArtwork",
  data() {
    return {
      boxActive: true,
      showPopup: false,
    };
  },
  methods: {
    registerBoxArtwork() {
      // Lógica para realizar o login
    },
    submitForm() {
      // Simular envio do formulário
      // Aqui você poderia enviar os dados para um servidor usando axios ou fetch

      // Mostrar o pop-up
      this.showPopup = true;

      // Esconder o pop-up após 2 segundos
      setTimeout(() => {
        this.showPopup = false;
      }, 2000);
    },
    openNewBox: function () {
      this.boxActive = true;
    },
    closeNewBox: function () {
      this.boxActive = false;
      this.$emit("closeBox");
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
