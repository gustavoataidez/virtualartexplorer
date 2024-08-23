<template>
    <div id="background" v-if="loginActive" >
    <div class="backdrop">
        <h2><a v-on:click="closeLogin">⭠</a>Login</h2>
        <form @submit.prevent="realizarLogin">
                <label for="email">E-mail</label>
                <input class="input-text" type="email" placeholder="Digite seu email">
                <label for="password">Senha</label>
                <input class="input-text" type="password" placeholder="Digite sua senha">
                <input class="btn-login" type="submit" value="Login">
            </form>
            <p>Não tem uma conta? <a href="./Register.vue">Resgistrar-se</a></p>
    </div>
    </div>
</template>

<script>
import RegisterModal from "./RegisterModal.vue";

export default {
  components: { RegisterModal },
  name: 'LoginModal',
  data(){
    return{
        loginActive: true,
        registerActive: false,
    }
  },
    methods:{
        realizarLogin() {
        // Lógica para realizar o login
        },
        openLogin: function(){
            this.loginActive = true;
        },
        closeLogin: function(){
            this.loginActive = false;
            this.$emit('closeLog')
        },
        openRegister() {
            this.registerActive = true;
        }
    }
  }
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
form{
    display: flex;
    flex-direction: column;
    text-align: left;
    gap: 2px;
}
.input-text{
    border: solid 1px rgb(179, 179, 179);
    border-radius: 10px;
    padding: 5px 10px;
    margin-bottom: 15px;
    font-size: 0.9rem;
}
.input-text:focus {
    outline: 0;
}
.btn-login{
    background-color: var(--color-orange);
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

</style>