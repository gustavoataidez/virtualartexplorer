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
    <div
      v-if="sessionActive"
      class="my-profile d-flex flex-row p-0 align-items-center"
    >
      <a class="btn1" v-on:click="openProfile">
        Hi, {{ firstName }} ({{ userId }})
      </a>
      <a class="btn mx-3 btn-danger" v-on:click="logout">Sair</a>
    </div>
    <div v-if="!sessionActive" class="d-flex flex-row p-0">
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
      :userId="userId"
      :userEmail="userEmail"
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
      firstName: "", // Nome do usuário logado
      userId: null, // ID do usuário logado
      userEmail: "", // Email do usuário logado
    };
  },
  mounted() {
  // Verifica se há uma sessão ativa no localStorage
  const token = localStorage.getItem("authToken");
  const firstName = localStorage.getItem("firstName");
  const userId = localStorage.getItem("userId");
  const userEmail = localStorage.getItem("userEmail");

  if (token && firstName && userId && userEmail) {
    this.firstName = firstName;
    this.userId = userId;
    this.userEmail = userEmail;
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
    setUserSession(user) {
      this.firstName = user.first_name;
      this.userId = user.id;
      this.userEmail = user.email;
      this.sessionActive = true;

      // Salvar os dados da sessão no localStorage
      localStorage.setItem("authToken", user.token);
      localStorage.setItem("firstName", user.first_name);
      localStorage.setItem("userId", user.id);
      localStorage.setItem("userEmail", user.email);
    },
    logout() {
      // Limpa os dados de autenticação no localStorage
      localStorage.removeItem("authToken");
      localStorage.removeItem("firstName");
      localStorage.removeItem("userId");
      localStorage.removeItem("userEmail");

      // Reseta o estado de sessão
      this.firstName = "";
      this.userId = null;
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
