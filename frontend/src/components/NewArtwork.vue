<template>
    <div  id="new-artwork" >
      
      <div  class="background" v-on:click="closeNewBox"></div>
        <form class="form" v-if="boxActive">
  <span class="title">New Artwork!</span>
  <!--<span class="sub mb">Register to get full access now :)</span>-->
 <input id="file" type="file">
  <label class="avatar" for="file"><span> <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><g stroke-width="0" id="SVGRepo_bgCarrier"></g><g stroke-linejoin="round" stroke-linecap="round" id="SVGRepo_tracerCarrier"></g><g id="SVGRepo_iconCarrier"> <path fill="#F79E1C" d="M17.1813 16.3254L15.3771 14.5213C16.5036 13.5082 17.379 12.9869 18.2001 12.8846C19.0101 12.7837 19.8249 13.0848 20.8482 13.8687C20.8935 13.9034 20.947 13.9202 21 13.9202V15.024C21 19.9452 19.9452 21 15.024 21H8.976C4.05476 21 3 19.9452 3 15.024V13.7522C3.06398 13.7522 3.12796 13.7278 3.17678 13.679L4.45336 12.4024C5.31928 11.5365 6.04969 10.8993 6.71002 10.4791C7.3679 10.0605 7.94297 9.86572 8.50225 9.86572C9.06154 9.86572 9.6366 10.0605 10.2945 10.4791C10.9548 10.8993 11.6852 11.5365 12.5511 12.4024L16.8277 16.679C16.9254 16.7766 17.0836 16.7766 17.1813 16.679C17.2789 16.5813 17.2789 16.423 17.1813 16.3254Z" opacity="0.1"></path> <path stroke-width="2" stroke="#F79E1C" d="M3 8.976C3 4.05476 4.05476 3 8.976 3H15.024C19.9452 3 21 4.05476 21 8.976V15.024C21 19.9452 19.9452 21 15.024 21H8.976C4.05476 21 3 19.9452 3 15.024V8.976Z"></path> <path stroke-linecap="round" stroke-width="2" stroke="#F79E1C" d="M17.0045 16.5022L12.7279 12.2256C9.24808 8.74578 7.75642 8.74578 4.27658 12.2256L3 13.5022"></path> <path stroke-linecap="round" stroke-width="2" stroke="#F79E1C" d="M21.0002 13.6702C18.907 12.0667 17.478 12.2919 15.1982 14.3459"></path> <path stroke-width="2" stroke="#F79E1C" d="M17 8C17 8.55228 16.5523 9 16 9C15.4477 9 15 8.55228 15 8C15 7.44772 15.4477 7 16 7C16.5523 7 17 7.44772 17 8Z"></path> </g></svg></span></label>
    <input type="name" class="input" placeholder="Title">
    <input type="text" class="input" placeholder="Author">
    <textarea type="text" class="input" placeholder="Description" rows="3"></textarea>
    <button >Salve</button>
</form>
<div v-if="showPopup" class="alert alert-success d-flex align-items-center" role="alert">
  <svg class="bi flex-shrink-0 me-2" role="img" aria-label="Success:"><use xlink:href="#check-circle-fill"/></svg>
  <div>
    An example success alert with an icon
  </div>
</div>
    </div>
</template>

<script>

export default{
  name: 'NewArtwork',
  data(){
    return{
        boxActive: true,
        showPopup: false,
    }
  },
    methods:{
        registerBoxArtwork() {
        // Lógica para realizar o login
        },
        submitForm() {
        // Simular envio do formulário
        // Aqui você poderia enviar os dados para um servidor usando axios ou fetch

        // Mostrar o pop-up
        this.showPopup = true;

        // Esconder o pop-up após 2 segundos
        setTimeout(() => {
          this.showPopup = false;
        }, 2000);
      },
        openNewBox: function(){
            this.boxActive = true;
        },
        closeNewBox: function(){
            this.boxActive = false;
            this.$emit('closeBox')
        }
    }
}
</script>


<style scoped>
.bi.flex-shrink-0.me-2{
  position: fixed;
  display: none;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 11;
}
.background{
  position:absolute;
  background-color:  rgba(0, 0, 0, 0.1);
  width: 100%;
  height: 100vh;
  z-index: 10;
}
.form {
  position: absolute;
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 300px;
  background-color: #fff;
  border-radius: 20px;
  padding: 30px 20px;
  box-shadow: 100px 100px 80px rgba(0, 0, 0, 0.03);
  z-index: 12;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.title {
  color: black;
  font-weight: bold;
  text-align: center;
  font-size: 20px;
  margin-bottom: 0px;
}

.avatar {
  height: 20%;
  width: 30%;
  background-color: #F9EBBF;
  margin-bottom: 4px;
  border-radius: 10%;
  align-self: center;
  padding: 6px;
  cursor: pointer;
  box-shadow: 12.5px 12.5px 10px rgba(0, 0, 0, 0.015),100px 100px 80px rgba(0, 0, 0, 0.03);
}

.form button {
  align-self: flex-end;
}

.input, button {
  border: none;
  outline: none;
  width: 100%;
  padding: 10px;
  background-color: rgb(247, 243, 243);
  border-radius: 10px;
  box-shadow: 12.5px 12.5px 10px rgba(0, 0, 0, 0.015),100px 100px 80px rgba(0, 0, 0, 0.03);
  font-size: 0.9rem;
}
button {
  margin-top: 12px;
  background-color: var(--vt-c-orange);
  color: #fff;
  text-transform: uppercase;
  font-weight: bold;
}

.input:focus {
  border: 1px solid var(--vt-c-orange);
}

#file {
  display: none;
}
</style>