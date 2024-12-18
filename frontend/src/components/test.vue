<template>
    <div class="login-container">
      <h2>Login</h2>
      <form @submit.prevent="realizarLogin">
        <label for="email">E-mail:</label>
        <input
          type="email"
          v-model="email"
          placeholder="Digite seu email"
          required
        />
  
        <label for="password">Senha:</label>
        <input
          type="password"
          v-model="password"
          placeholder="Digite sua senha"
          required
        />
  
        <button type="submit">Entrar</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        email: "", // Email do usuário
        password: "", // Senha do usuário
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
      alert("Usuário cadastrado");
    } else {
      alert("Erro ao realizar login. Verifique suas credenciais.");
    }
  } catch (error) {
    console.error("Erro ao realizar login:", error);
    alert("Erro ao acessar o servidor. Tente novamente mais tarde.");
  }
},
    },
  };
  </script>
  
  <style scoped>
  .login-container {
    max-width: 400px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  
  .login-container h2 {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .login-container label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  .login-container input {
    width: 100%;
    padding: 10px;
    margin-bottom: 15px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 1rem;
  }
  
  .login-container button {
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    cursor: pointer;
  }
  
  .login-container button:hover {
    background-color: #0056b3;
  }
  </style>
  