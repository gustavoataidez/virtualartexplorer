<template>
  <div
    class="fundo d-flex flex-row align-items-center justify-content-around align-items-baseline py-3"
  >
    <div class="menu-items">
      <ul class="d-flex flex-row p-0">
        <li><router-link to="/museums">Museus</router-link></li>
        <li><router-link class="px-4" to="/artworks">Obras</router-link></li>
      </ul>
    </div>
    <div>
      <router-link to="/">
        <img src="../assets/logo-virtual-marrom.png" alt="" class="logo" />
      </router-link>
    </div>

    <!-- Sessão do usuário logado -->
    <div v-if="$store.getters.isLoggedIn" class="my-profile d-flex flex-row p-0 align-items-center">
  <a class="btn1" v-on:click="openProfile">
    Olá {{ $store.getters.userFirstName }} (#{{ $store.state.user?.id }})
  </a>
  <a class="btn mx-3 btn-danger" v-on:click="logout">Sair</a>
</div>

    <!-- Sessão de login e registro -->
    <div v-else class="d-flex flex-row p-0">
      <a class="btn" v-on:click="openLogin">Entrar</a>
      <LoginModal
        v-if="loginActive"
        @closeLog="closeLogin"
        @userLoggedIn="setUserSession"
      />
      <a class="btn mx-2 btn-outline-light register" v-on:click="openRegister">
        Registrar
      </a>
      <RegisterModal v-if="registerActive" @closeReg="closeRegister" />
    </div>

    <!-- Modal do Perfil -->
    <ProfileModal
      v-if="profileActive"
      :userId="$store.state.user?.id"
      :userEmail="$store.state.user?.email"
      @closeProfile="closeProfile"
      @navigateToCreate="navigateToCreate"
    />
  </div>
</template>

<script>
import LoginModal from "./LoginModal.vue";
import RegisterModal from "./RegisterModal.vue";
import ProfileModal from "./ProfileModal.vue";

export default {
  components: { LoginModal, RegisterModal, ProfileModal },
  data() {
    return {
      loginActive: false,
      registerActive: false,
      profileActive: false,
    };
  },
  mounted() {
  console.log("Vuex State:", this.$store.state);
  console.log("Is Logged In:", this.$store.getters.isLoggedIn);
},
  methods: {
    openProfile() {
      this.profileActive = true;
    },
    closeProfile() {
      this.profileActive = false;
    },
    navigateToCreate() {
      this.$router.push("/museum/create");
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
    setUserSession(user) {
      this.$store.dispatch("login", user); // Envia os dados para o Vuex
      this.loginActive = false; // Fecha o modal de login
    },
    logout() {
      this.$store.dispatch("logout"); // Executa o logout via Vuex
      this.$router.push("/"); // Redireciona para a página inicial
    },
  },
};
</script>

<style scoped>
.fundo {
  position: relative;
  background-color: var(--vt-c-header);
  width: 100%;
  z-index: 20;
}
.logo {
  width: 100%;
  max-width: 180px;
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