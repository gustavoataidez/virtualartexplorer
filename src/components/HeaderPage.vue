<template>
  <div
    class="fundo d-flex flex-row align-items-center justify-content-around align-items-baseline py-3"
  >
    <div class="menu-items">
      <ul class="d-flex flex-row p-0">
        <li><router-link to="/museum">Museus</router-link></li>
        <li><router-link class="px-2" to="/artwork">Obras</router-link></li>
        <li><router-link to="/museum/edit">Exposições</router-link></li>
      </ul>
    </div>
    <div class="">
      <router-link to="/">
        <img src="../assets/logo-virtual-marrom.png" alt="" class="logo" />
      </router-link>
    </div>
    <div
      v-if="sessionActive"
      class="my-profile d-flex flex-row p-0 align-items-center"
    >
      <a class="btn" v-on:click="openProfile"
        ><svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 448 512"
          width="20px"
          fill="#481810"
        >
          <path
            d="M304 128a80 80 0 1 0 -160 0 80 80 0 1 0 160 0zM96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM49.3 464l349.5 0c-8.9-63.3-63.3-112-129-112l-91.4 0c-65.7 0-120.1 48.7-129 112zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3z"
          />
        </svg>
        Meu Perfil</a
      >
    </div>
    <div v-if="!sessionActive" class="d-flex flex-row p-0">
      <a class="btn mx-3" v-on:click="openLogin">Entrar</a>
      <LoginModal v-if="loginActive" @closeLog="closeLogin" />
      <a class="btn mx-2 btn-outline-light register" v-on:click="openRegister"
        >Registrar</a
      >
      <RegisterModal v-if="registerActive" @closeReg="closeRegister" />
    </div>
  </div>
</template>

<script>
import LoginModal from "./LoginModal.vue";
import RegisterModal from "./RegisterModal.vue";

export default {
  components: { LoginModal, RegisterModal },
  data() {
    return {
      loginActive: false,
      registerActive: false,
      sessionActive: true,
      profileActive: false,
    };
  },
  methods: {
    scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: "smooth",
      });
    },
    openProfile() {
      this.profileActive = true;
    },
    closeProfile() {
      this.profileActive = false;
    },
    openLogin() {
      this.loginActive = true;
    },
    closeLogin() {
      this.loginActive = false;
    },
    openRegister() {
      this.registerActive = true;
    },
    closeRegister() {
      this.registerActive = false;
    },
  },
};
</script>

<style scoped>
.fundo {
  position: relative;
  background-color: var(--vt-c-header);
  width: 99%;
  z-index: 20;
}
.logo {
  width: 100%;
  max-width: 180px;
}
.arrow {
  max-width: 40px;
}
.code {
  max-width: 24px;
}
a {
  color: var(--vt-c-brown);
  font-family: "Poppins", sans-serif;
  cursor: pointer;
  font-size: 1rem;
}
ul {
  list-style: none;
}
.btn.mx-2.btn-outline-light.register {
  border: solid 1px var(--vt-c-brown);
}
</style>
