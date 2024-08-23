<template>
    <div id="background" v-if="registerActive" >
    <div class="backdrop px-10">
        <h2><a v-on:click="closeRegister">⭠</a>Register</h2>
        <form class="forms" @submit.prevent="realizarRegister">
            <div class="caixa-form">
                <label for="ger_firstname" class="required">Nome</label>
                <input class="input-text" name="ger_firstname" type="text" placeholder="Digite seu nome" maxlength="50">
            </div>
            <div class="caixa-form">
                <label for="ger_lastname" class="required">Nome</label>
                <input class="input-text" name="ger_lastname" type="text" placeholder="Digite seu nome" maxlength="50">
            </div>
            <div class="caixa-form">
                <label class="required" for="ger_email">Email</label>
                <input class="input-text" name="ger_email" type="email" placeholder="Digite seu email" maxlength="80" required >
            </div>
            <div class="caixa-form">
                <label for="ger_password" class="required">Senha</label>
                <input class="input-text" name="ger_password" type="password" placeholder="Digite sua senha" maxlength="50" required>
            </div>
                <input class="btn-login" type="button" value="Registrar">
            </form>
            <p>Já tem uma conta?<a href="">Faça Login</a></p>
    </div>
    </div>
</template>

<script>
export default {
  name: 'RegisterModal',
  data(){
    return{
        registerActive: true
    }
  },
    methods:{
        realizarRegistro() {
        // Lógica para realizar registro
        },
        openRegister: function(){
            this.registerActive = true;
        },
        closeRegister: function(){
            this.$emit('closeReg');
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
.required:after {
  content:" *";
  color: red;
}
</style>
