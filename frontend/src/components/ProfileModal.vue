<template>
  <div id="background">
    <div class="backdrop">
      <h2>Meu Perfil</h2>
      <p>Bem-vindo(a), {{ $store.getters.userFirstName }}</p>
      <button class="btn-action" @click="$emit('navigateToCreate')">
        Criar Museu
      </button>

      <div v-if="museums.length > 0" class="museums-list">
        <h3>Seus Museus</h3>
        <div v-for="museum in museums" :key="museum.id">
          <button class="btn-museum" @click="openMuseum(museum.id)">
            {{ museum.title }}
          </button>
        </div>
      </div>

      <button class="btn-secondary" @click="$emit('closeProfile')">
        Fechar
      </button>
    </div>
  </div>
</template>

<script>
import { API_URL } from "@/config";

export default {
  name: "ProfileModal",
  data() {
    return {
      museums: [], // Lista de museus do usuário
    };
  },
  computed: {
    // Obtém o token diretamente do Vuex para autenticação
    token() {
      return this.$store.state.token;
    },
  },
  mounted() {
    this.fetchMuseums();
  },
  methods: {
    async fetchMuseums() {
      try {
        // Faz a requisição para a API usando o token de autenticação
        const response = await fetch(`${API_URL}/museums/my`, {
          method: "GET",
          headers: {
            Authorization: `Bearer ${this.token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          this.museums = data; // Atualiza a lista de museus
        } else {
          console.error("Erro ao buscar museus: ", await response.text());
        }
      } catch (error) {
        console.error("Erro ao buscar os museus:", error);
      }
    },
    openMuseum(museumId) {
      // Redireciona para a página de edição do museu
      this.$router.push(`/museum/edit/${museumId}`);
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

h3 {
  color: var(--vt-c-black);
  margin-bottom: 10px;
  font-weight: 600;
}

p {
  color: var(--vt-c-black);
  margin-bottom: 20px;
  font-size: 0.9rem;
}

.museums-list {
  margin-top: 20px;
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

.btn-museum {
  background-color: var(--vt-c-orange);
  border: 0px;
  border-radius: 20px;
  font-size: 0.9rem;
  padding: 8px;
  color: white;
  font-weight: bold;
  margin: 10px 0;
  width: 100%;
  text-align: center;
  cursor: pointer;
}

.btn-museum:hover {
  opacity: 0.7;
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