<template>
    <div id="background" v-if="registerActive">
      <div class="backdrop px-10">
        <h2>
          <a v-on:click="closeRegister">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" width="30px">
              <path d="M512 256A256 256 0 1 0 0 256a256 256 0 1 0 512 0zM215 127c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-71 71L392 232c13.3 0 24 10.7 24 24s-10.7 24-24 24l-214.1 0 71 71c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0L103 273c-9.4-9.4-9.4-24.6 0-33.9L215 127z" />
            </svg>
          </a>
          Register
        </h2>
        <form class="forms" @submit.prevent="realizarRegistro">
          <div class="caixa-form">
            <label for="ger_firstname" class="required">Nome</label>
            <input
              class="input-text"
              v-model="formData.first_name"
              name="ger_firstname"
              type="text"
              placeholder="Digite seu nome"
              maxlength="50"
              required
            />
          </div>
          <div class="caixa-form">
            <label for="ger_lastname" class="required">Sobrenome</label>
            <input
              class="input-text"
              v-model="formData.last_name"
              name="ger_lastname"
              type="text"
              placeholder="Digite seu sobrenome"
              maxlength="50"
              required
            />
          </div>
          <div class="caixa-form">
            <label class="required" for="ger_email">Email</label>
            <input
              class="input-text"
              v-model="formData.email"
              name="ger_email"
              type="email"
              placeholder="Digite seu email"
              maxlength="80"
              required
            />
          </div>
          <div class="caixa-form">
            <label for="ger_password" class="required">Senha</label>
            <input
              class="input-text"
              v-model="formData.password"
              name="ger_password"
              type="password"
              placeholder="Digite sua senha"
              maxlength="50"
              required
            />
          </div>
          <input class="btn-login" type="submit" value="Registrar" />
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    name: "RegisterModal",
    data() {
      return {
        registerActive: true,
        formData: {
          first_name: "",
          last_name: "",
          email: "",
          password: ""
        }
      };
    },
    methods: {
      async realizarRegistro() {
        try {
          // Envia os dados para a API
          const response = await axios.post(
            "http://localhost:3000/api/v1/managers",
            this.formData
          );
          // Exibe uma mensagem de sucesso
          alert("Usuário registrado com sucesso!");
          console.log(response.data);
  
          // Fecha o modal após o registro
          this.closeRegister();
        } catch (error) {
          // Trata erros e exibe mensagens ao usuário
          console.error("Erro ao registrar usuário:", error);
          if (error.response && error.response.data) {
            alert(`Erro: ${error.response.data.message}`);
          } else {
            alert("Erro desconhecido ao registrar o usuário.");
          }
        }
      },
      openRegister() {
        this.registerActive = true;
      },
      closeRegister() {
        this.$emit("closeReg");
      }
    }
  };
  </script>  

<style scoped>
#background{
    height: 100%;
    width: 100%;
    background-color: rgba(0, 0, 0, .8);
    position: fixed;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    z-index: 10;
}
h2{
    color: var(--vt-c-black);
    margin-bottom: 20px;
    font-weight: bold;
}
.backdrop{
    position: absolute;
    width: 50%;
    max-width: 400px;
    min-width: 300px;
    padding: 30px 20px;
    background: white;
    border-radius: 20px;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    z-index: 11;
}
.forms{
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    text-align: left;
    gap: 0px;
}
.input-text{
    border: solid 1px rgb(179, 179, 179);
    border-radius: 10px;
    padding: 5px 10px;
    margin-bottom: 15px;
    font-size: 0.9rem;
    width: 100%;
    min-width: 100px;
}
.input-text:focus {
    outline: 0;
}
.caixa-form{
    display: flex;
    flex-direction: column;
    width: 49%;
    padding: 0% 1%;
}
.btn-login{
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
.btn-login:hover{
    opacity: 0.6;
    transition: 0.3s;
}
label{
    color: var(--vt-c-black);
    font-weight: 500;
    font-size: 0.9rem;
}

a{
    color: var(--color-orange);
}
a:hover{
cursor: pointer;
}
p{
    color: var(--vt-c-black);
    margin-top: 20px;
    font-size: 0.9rem;
    text-align: center;
}
.required:after {
  content:" *";
  color: red;
}
</style>
