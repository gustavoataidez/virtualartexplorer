<template>
  <div id="background" v-if="loginActive">
    <div class="backdrop">
      <h2>
        <a v-on:click="closeLogin">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" width="30px">
            <path
              d="M512 256A256 256 0 1 0 0 256a256 256 0 1 0 512 0zM215 127c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-71 71L392 232c13.3 0 24 10.7 24 24s-10.7 24-24 24l-214.1 0 71 71c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0L103 273c-9.4-9.4-9.4-24.6 0-33.9L215 127z"
            />
          </svg>
        </a>
        Login
      </h2>
      <form @submit.prevent="realizarLogin">
        <label for="email">E-mail</label>
        <input
          class="input-text"
          type="email"
          v-model="email"
          placeholder="Digite seu email"
          required
        />
        <label for="password">Senha</label>
        <input
          class="input-text"
          type="password"
          v-model="password"
          placeholder="Digite sua senha"
          required
        />
        <input class="btn-login" type="submit" value="Login" />
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "LoginModal",
  data() {
    return {
      loginActive: true,
      email: "", // Entrada do email do usuário
      password: "", // Entrada da senha do usuário
    };
  },
  methods: {
    async realizarLogin() {
      try {
        const response = await axios.post("http://localhost:3000/api/v1/login", {
          email: this.email,
          password: this.password,
        });

        if (response.data && response.data.token) {
          // Emita o evento passando os dados do usuário logado
          this.$emit("userLoggedIn", response.data);
          alert("Login realizado com sucesso!");
          this.closeLogin();
        } else {
          alert("Login falhou. Tente novamente.");
        }
      } catch (error) {
        console.error("Erro ao realizar login:", error);
        alert("Credenciais inválidas ou erro no servidor.");
      }
    },
    closeLogin() {
      this.loginActive = false;
      this.$emit("closeLog");
    },
  },
};
</script>

<style scoped>
#background {
  height: 100%;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 10;
}

h2 {
  color: var(--vt-c-black);
  margin-bottom: 20px;
  font-weight: bold;
}

.backdrop {
  position: absolute;
  max-width: 400px;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 11;
  width: 50%;
  min-width: 300px;
  padding: 30px 20px;
  background: white;
  border-radius: 20px;
}

form {
  display: flex;
  flex-direction: column;
  text-align: left;
  gap: 2px;
}

.input-text {
  border: solid 1px rgb(179, 179, 179);
  border-radius: 10px;
  padding: 5px 10px;
  margin-bottom: 15px;
  font-size: 0.9rem;
}

.input-text:focus {
  outline: 0;
}

.btn-login {
  background-color: var(--vt-c-orange);
  border: 0px;
  border-radius: 50px;
  font-size: 0.9rem;
  padding: 12px;
  color: white;
  font-weight: bold;
  margin-top: 10px;
  width: 100%;
  text-align: center;
}

.btn-login:hover {
  opacity: 0.6;
  transition: 0.3s;
}

label {
  color: var(--vt-c-black);
  font-weight: 500;
  font-size: 0.9rem;
}

a {
  color: var(--color-orange);
}

a:hover {
  cursor: pointer;
}

p {
  color: var(--vt-c-black);
  margin-top: 20px;
  font-size: 0.9rem;
  text-align: center;
}
</style>
