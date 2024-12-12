<template>
  <div
    class="fundo d-flex flex-row align-items-center justify-content-around align-items-baseline pt-4"
  >
    <div class="menu-items">
      <ul class="d-flex flex-row p-0">
        <li><router-link to="/museum">Museus</router-link></li>
        <li><router-link class="px-2" to="/artwork">Obras</router-link></li>
        <li><router-link to="/museum">Exposições</router-link></li>
      </ul>
    </div>
    <div class="">
      <router-link to="/">
        <img src="../assets/logo-virtual-branca.png" alt="" class="logo" />
      </router-link>
    </div>
    <div
      v-if="sessionActive"
      class="my-profile d-flex flex-row p-0 align-items-center"
    >
      <a class="btn1" v-on:click="openProfile">
        Hi, {{ userEmail }}
      </a>
      <a class="btn mx-3 btn-danger" v-on:click="logout">Sair</a>
    </div>
    <div class="d-flex flex-row p-0" v-if="!sessionActive">
      <a class="btn mx-3" v-on:click="openLogin">Entrar</a>
      <LoginModal
        v-if="loginActive"
        @closeLog="closeLogin"
        @userLoggedIn="setUserSession"
      />
      <a class="btn mx-2 btn-outline-light" v-on:click="openRegister">
        Registrar
      </a>
      <RegisterModal v-if="registerActive" @closeReg="closeRegister" />
    </div>

    <!-- Modal do Perfil -->
    <ProfileModal
      v-if="profileActive"
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
      sessionActive: false,
      profileActive: false,
      userEmail: "", // Armazena o email do usuário logado
    };
  },
  mounted() {
    // Verifica se há uma sessão ativa no localStorage
    const token = localStorage.getItem("authToken");
    const email = localStorage.getItem("userEmail");

    if (token && email) {
      this.userEmail = email;
      this.sessionActive = true;
    }
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
    setUserSession(email) {
      this.userEmail = email;
      this.sessionActive = true;

      // Salvar os dados da sessão no localStorage
      localStorage.setItem("userEmail", email);
    },
    logout() {
      // Limpa os dados de autenticação no localStorage
      localStorage.removeItem("authToken");
      localStorage.removeItem("userEmail");

      // Reseta o estado de sessão
      this.userEmail = "";
      this.sessionActive = false;

      // Redireciona para a página inicial
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.fundo {
  position: absolute;
  background-color: rgb(0, 0, 0, 0);
  width: 100%;
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
  color: white;
  font-family: "Poppins", sans-serif;
  cursor: pointer;
  font-size: 1rem;
}
ul {
  list-style: none;
}
</style>
