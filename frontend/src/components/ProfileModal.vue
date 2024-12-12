<template>
    <div id="background">
      <div class="backdrop">
        <h2>Meu Perfil</h2>
        <p>Bem-vindo(a), {{ userEmail }}</p>
        <button class="btn-action" @click="$emit('navigateToCreate')">
          Criar Museu
        </button>
        <button class="btn-secondary" @click="$emit('closeProfile')">
          Fechar
        </button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "ProfileModal",
    props: {
      userEmail: {
        type: String,
        required: true,
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
    text-align: center;
  }
  
  h2 {
    color: var(--vt-c-black);
    margin-bottom: 20px;
    font-weight: bold;
  }
  
  p {
    color: var(--vt-c-black);
    margin-bottom: 20px;
    font-size: 0.9rem;
  }
  
  .btn-action {
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
    cursor: pointer;
  }
  
  .btn-action:hover {
    opacity: 0.6;
    transition: 0.3s;
  }
  
  .btn-secondary {
    background-color: var(--vt-c-black);
    border: 0px;
    border-radius: 50px;
    font-size: 0.9rem;
    padding: 12px;
    color: white;
    font-weight: bold;
    margin-top: 10px;
    width: 100%;
    text-align: center;
    cursor: pointer;
  }
  
  .btn-secondary:hover {
    opacity: 0.6;
    transition: 0.3s;
  }
  </style>
  