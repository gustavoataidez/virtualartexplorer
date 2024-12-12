<template>
  <div class="login-container">
    <div class="login-card">
      <h2>Login</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="email">Email</label>
          <input
            id="email"
            type="email"
            v-model="email"
            placeholder="Digite seu email"
            required
          />
        </div>
        <div class="form-group">
          <label for="password">Senha</label>
          <input
            id="password"
            type="password"
            v-model="password"
            placeholder="Digite sua senha"
            required
          />
        </div>
        <button type="submit" class="btn btn-primary" :disabled="loading">
          {{ loading ? "Entrando..." : "Entrar" }}
        </button>
        <div v-if="error" class="error-message">{{ error }}</div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Login",
  data() {
    return {
      email: "user@example.com", // Preenchido para teste
      password: "123456", // Preenchido para teste
      loading: false,
      error: null,
    };
  },
  methods: {
    async handleLogin() {
      const apiUrl = "http://localhost:3000/auth/login"; // Substitua pelo seu endpoint de autenticação
      this.loading = true;
      this.error = null;

      try {
        const response = await axios.post(apiUrl, {
          email: this.email,
          password: this.password,
        });

        if (response.data && response.data.token) {
          // Salvar o token no localStorage
          localStorage.setItem("authToken", response.data.token);

          // Redirecionar para a página principal ou painel do usuário
          this.$router.push("/dashboard"); // Substitua pela rota desejada
        } else {
          throw new Error("Credenciais inválidas");
        }
      } catch (err) {
        this.error = err.response?.data?.message || "Erro ao realizar login";
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}

.login-card {
  width: 300px;
  padding: 20px;
  border-radius: 8px;
  background: white;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.form-group {
  margin-bottom: 15px;
  text-align: left;
}

label {
  font-size: 0.9rem;
  color: #333;
}

input {
  width: 100%;
  padding: 8px;
  margin-top: 5px;
  font-size: 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  font-size: 1rem;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

button:disabled {
  background-color: #d6d6d6;
  cursor: not-allowed;
}

.error-message {
  color: red;
  font-size: 0.9rem;
  margin-top: 10px;
}
</style>
