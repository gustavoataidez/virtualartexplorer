<template>
  <HeaderNovo></HeaderNovo>
  <div class="museu-page">
    <div class="sec-resume">
      <div>
        <p class="form__label">Capa</p>
        <label class="custum-file-upload" for="file">
          <div class="icon">
            <svg xmlns="http://www.w3.org/2000/svg" fill="" viewBox="0 0 24 24"></svg>
          </div>
          <div class="text">
            <span>Upload da foto do museu</span>
          </div>
        </label>
        <input
          class="form-control"
          style="background:transparent;border: dashed 0px #cacaca;"
          type="file"
          id="file"
          :disabled="museumCreated"
          @change="onFileChange"
        />
      </div>
      <div class="resume-desc">
        <div class="form__group field">
          <label for="title" class="form__label">Título</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.title"
            :disabled="museumCreated"
            placeholder="Ex. Museu das Artes"
            required
          />

          <label for="description" class="form__label">Descrição</label>
          <textarea
            class="form__field texto"
            v-model="museum.description"
            :disabled="museumCreated"
            placeholder="Descreva sobre o museu"
            name="description"
            required
          ></textarea>

          <div class="category-section">
            <div class="category-field">
              <label for="category1" class="form__label">Categoria 1</label>
              <select class="form-select" v-model="museum.category1" :disabled="museumCreated">
                <option disabled value="">Selecione uma categoria</option>
                <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
              </select>
            </div>

            <div class="category-field">
              <label for="category2" class="form__label">Categoria 2</label>
              <select class="form-select" v-model="museum.category2" :disabled="museumCreated">
                <option disabled value="">Selecione uma categoria</option>
                <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
              </select>
            </div>
          </div>

          <label for="link" class="form__label">Link</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.link"
            :disabled="museumCreated"
            placeholder="Ex. http://museucultura.com"
          />

          <label for="address" class="form__label">Endereço</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.address"
            :disabled="museumCreated"
            placeholder="Ex. Avenida das Tradições, 303"
          />

          <label for="cep" class="form__label">CEP</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.cep"
            :disabled="museumCreated"
            placeholder="Ex. 88900-111"
          />

          <label for="state" class="form__label">Estado</label>
          <select
            class="form-select"
            v-model="museum.state"
            :disabled="museumCreated"
            name="state"
            @change="onStateChange"
            style="background-color:#e2e2e2;"
          >
            <option disabled value="">Selecione um estado</option>
            <option v-for="est in states" :key="est.uf" :value="est.uf">{{ est.nome }}</option>
          </select>

          <label for="city" class="form__label">Cidade</label>
          <select
            class="form-select"
            v-model="museum.city"
            :disabled="museumCreated"
            name="city"
          >
            <option disabled value="">Selecione uma cidade</option>
            <option v-for="c in cities" :key="c" :value="c">{{ c }}</option>
          </select>

          <label for="information" class="form__label">Mais Informações</label>
          <textarea
            class="form__field texto"
            v-model="museum.information"
            :disabled="museumCreated"
            placeholder="Ex. Inclui exposições sobre danças, culinária e festivais regionais."
          ></textarea>


        </div>
      </div>
    </div>
    
    <div class="save-museum-section" v-if="!museumCreated">
      <button class="btn btn-success" @click="createMuseum">Salvar Museu</button>
    </div>

    <div v-if="museumCreated" class="add-obras-section">
      <h3>Adicionar Obras</h3>
      <button class="btn btn-primary" @click="openObraModal">Adicionar Obra</button>

      <ul class="obra-list" v-if="works.length > 0">
        <li v-for="(obra, index) in works" :key="obra.id">
          <strong>{{ obra.name }}</strong> - {{ obra.description }} - {{ obra.author || 'Autor Desconhecido' }} - {{ obra.image }}
          <a href="#" @click.prevent="deleteObra(obra.id, index)" style="color:red; margin-left:10px;">Excluir</a>
        </li>
      </ul>

      <div class="finish-section">
        <button class="btn btn-success" @click="finishRegistration">Concluir</button>
      </div>
    </div>

    
    <div v-if="showObraModal" class="modal-overlay" @click="closeObraModal">
      <div class="modal-content" @click.stop>
        <h4>Adicionar Obra</h4>
        <label>Nome da Obra</label>
        <input type="text" v-model="newObra.name" placeholder="Ex. O Grito" required/>

        <label>Descrição</label>
        <textarea v-model="newObra.description" placeholder="Descrição da obra" required></textarea>

        <label>Autor</label>
        <input type="text" v-model="newObra.author" placeholder="Ex. Edvard Munch"/>
        <button class="btn btn-success my-2" @click="createObra">Salvar Obra</button>
        <button class="btn btn-danger" @click="closeObraModal">Fechar</button>
      </div>
    </div>
  </div>
  <FooterNovo></FooterNovo>
</template>
<script>
import HeaderNovo from "../Header.vue";
import FooterNovo from "../Footer.vue";
import { API_URL } from '@/config';

export default {
  components: { HeaderNovo, FooterNovo },
  data() {
    return {
      museum: {
        title: "",
        description: "",
        category1: "",
        category2: "",
        link: "",
        address: "",
        cep: "",
        city: "",
        state: "",
        information: "",
        manager_id: null,
        capa: null,
      },
      museumCreated: false,
      museumId: null,
      showObraModal: false,
      newObra: {
        name: "",
        description: "",
        author: "",
        image: null, // Define como null por padrão
      },
      works: [],
      categories: ["esportes", "pessoas", "escravidão", "cultura"],
      states: [],
      cities: []
    };
  },
  async created() {
    await this.fetchStates();
  },
  methods: {
    async fetchStates() {
      try {
        const res = await fetch('https://servicodados.ibge.gov.br/api/v1/localidades/estados');
        const data = await res.json();
        this.states = data.map(state => ({
          uf: state.sigla,
          nome: state.nome
        }));
      } catch (error) {
        console.error("Erro ao carregar estados:", error);
      }
    },
    async onStateChange() {
      if (!this.museum.state) return;
      try {
        const res = await fetch(`https://servicodados.ibge.gov.br/api/v1/localidades/estados/${this.museum.state}/municipios`);
        const data = await res.json();
        this.cities = data.map(city => city.nome);
        this.museum.city = "";
      } catch (error) {
        console.error("Erro ao carregar cidades:", error);
      }
    },
    async deleteObra(id, index) {
      try {
        const response = await fetch(`${API_URL}/artworks/${id}`, {
          method: 'DELETE',
        });
        if (response.ok) {
          this.works.splice(index, 1);
        } else {
          alert("Erro ao excluir a obra.");
        }
      } catch (error) {
        console.error("Erro ao excluir a obra:", error);
        alert("Erro ao excluir a obra.");
      }
    },
    onFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        this.museum.capa = file;
      }
    },
    
    async createMuseum() {
  if (!this.museum.title || !this.museum.description) {
    alert("Título e Descrição são obrigatórios.");
    return;
  }

  // FormData para enviar os dados do museu
  const formData = new FormData();
  formData.append("title", this.museum.title);
  formData.append("description", this.museum.description);
  formData.append("category1", this.museum.category1);
  formData.append("category2", this.museum.category2);
  formData.append("link", this.museum.link);
  formData.append("address", this.museum.address);
  formData.append("cep", this.museum.cep);
  formData.append("city", this.museum.city);
  formData.append("state", this.museum.state);
  formData.append("information", this.museum.information);
  
  // Adicionar imagem ao FormData (se existir)
  if (this.museum.capa) {
    formData.append("image", this.museum.capa);
  }

  try {
    // Obtendo o token do Vuex
    const token = this.$store.state.token;

    const response = await fetch(`${API_URL}/museums`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`, // Adiciona o token no cabeçalho
      },
      body: formData, // Envia o FormData
    });

    if (response.ok) {
      const createdMuseum = await response.json();
      this.museumId = createdMuseum.id; // Salva o ID do museu criado
      this.museumCreated = true; // Atualiza o estado para indicar que o museu foi criado
      alert("Museu criado com sucesso!");
    } else {
      const errorResponse = await response.json();
      alert(`Erro ao cadastrar o museu: ${errorResponse.error}`);
    }
  } catch (error) {
    console.error("Erro ao cadastrar:", error);
    alert("Erro ao cadastrar o museu.");
  }
},
    openObraModal() {
      this.showObraModal = true;
    },
    closeObraModal() {
      this.showObraModal = false;
      this.newObra = {
        name: "",
        description: "",
        author: "",
        image: ""
      };
    },
    async createObra() {
  // Validação dos campos obrigatórios
  if (!this.newObra.name || !this.newObra.description) {
    alert("Nome da Obra e Descrição são obrigatórios.");
    return;
  }

  // Criação dos dados da obra
  const obraData = {
    museum_id: this.museumId,
    name: this.newObra.name.trim(), // Garante que não haja espaços em branco
    description: this.newObra.description.trim(),
    author: this.newObra.author?.trim() || "Autor Desconhecido", // Valor padrão para autor
    active: true, // Valor padrão para ativo
  };

  // Adicionar imagem se estiver presente
  if (this.newObra.image) {
    obraData.image = this.newObra.image;
  }

  try {
    // Recuperar o token do Vuex
    const token = this.$store.state.token;

    // Requisição para a API
    const response = await fetch(`${API_URL}/artworks`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`, // Adicionar o token no cabeçalho
      },
      body: JSON.stringify(obraData), // Converte os dados para JSON
    });

    if (response.ok) {
      const createdObra = await response.json();
      this.works.push(createdObra); // Atualiza a lista de obras
      this.closeObraModal(); // Fecha o modal
      alert("Obra cadastrada com sucesso!");
    } else {
      const errorResponse = await response.json();
      alert(`Erro ao cadastrar a obra: ${errorResponse.error}`);
    }
  } catch (error) {
    console.error("Erro ao cadastrar obra:", error);
    alert("Erro ao cadastrar a obra.");
  }
},
    finishRegistration() {
      alert("Museu criado com sucesso!");
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.category-section {
  display: flex;
  gap: 20px;
  margin-bottom: 15px;
}
.disabled-field {
  opacity: 0.5;
}
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width:100%;
  height:100%;
  display:flex;
  justify-content:center;
  align-items:center;
}
.modal-content {
  position: fixed;
  background: #fff;
  border-radius:20px;
  box-shadow: 0 0 100px rgba(0,0,0,0.6);
  padding:20px;
  max-width:400px;
  width:100%;
}
.category-section {
  display: flex;
  gap: 20px;
  width: 100%;
}
.disabled-field {
  opacity: 0.4;
}
.category-field{
  width: 50%;
}
.museu-page {
  padding: 20px;
}
.form__label {
  margin-top: 10px;
  font-size: 0.9rem;
}
.form__field {
  margin-bottom: 10px;
}
.texto {
  height: 100px;
}
.save-museum-section, .finish-section {
  text-align: center;
  margin-bottom: 30px;
}

.add-obras-section {
  margin-top: 30px;
}

.obra-list {
  list-style: none;
  padding: 0;
  margin: 20px 0;
}
.obra-list li {
  margin-bottom: 10px;
}

/* Modal styles */
.modal-overlay {
  position: fixed;
  top:0;left:0;right:0;bottom:0;
  display:flex;
  justify-content:center;
  align-items:center;
  z-index:9999;
}
.modal-content h4 {
  margin-bottom:15px;
}
.modal-content label {
  font-size: 0.9rem;
  margin: 5px 0;
}
.modal-content input, .modal-content textarea {
  width:100%;
  margin-bottom:10px;
  padding:5px;
  border:1px solid #ccc;
  border-radius:4px;
}

.modal-form {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;    
  position: absolute;
    max-width: 400px;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    z-index: 11;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  background: none;
  border: none;
  font-size: 26px;
  cursor: pointer;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}



.form__group {
  position: relative;
  padding: 20px 0;
  width: 100%;
  max-width: 100%;
}


.form__field {
  font-family: inherit;
  width: 100%;
  border: none;
  border-bottom: 2px solid #9b9b9b;
  outline: 0;
  font-size: 1rem;
  color: #000;
  padding: 4px 8px;
  background-color: #e2e2e2;
  transition: border-color 0.2s;
  margin-bottom: 5px;
}
.form__field.texto{
  font-size: 1rem;
  height: 10  0px;
}

.form__field::placeholder {
  color: #cbcbcb;
}

.form__field:placeholder-shown ~ .form__label {
  font-size: 17px;
  cursor: text;
}

.form__label {
  position: relative;
  top: 0;
  display: block;
  transition: 0.2s;
  font-size: 17px;
  color: rgba(75, 85, 99, 1);
  pointer-events: none;
}


/* reset input */
.form__field:required, .form__field:invalid {
  box-shadow: none;
}




.custum-file-upload {
width: 100%;
height: 400px;
border-radius: 50px;
  display: flex;
  flex-direction: column;
  align-items: space-between;
  gap: 20px;
  cursor: pointer;
  align-items: center;
  justify-content: center;
  border: 2px dashed #cacaca;
  background-color: #e2e2e2;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0px 48px 35px -48px rgba(0,0,0,0.1);
}

.custum-file-upload .icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.custum-file-upload .icon svg {
  height: 80px;
  fill: rgba(75, 85, 99, 1);
}

.custum-file-upload .text {
  display: flex;
  align-items: center;
  justify-content: center;
}

.custum-file-upload .text span {
  font-weight: 400;
  color: rgba(75, 85, 99, 1);
}

.custum-file-upload input {
  display: none;
}


.sec-resume{
display: flex;
flex-direction: row;
justify-content: center;
align-items: center;
gap: 50px;
}
.resume-desc{
width: 50%;
}
.resume-desc p{
font-size: 1.1rem;
}
.capa-museu{
    width: 50%;
    line-height: 1rem;
}
.grade-obras{
display: flex;
flex-direction: row;
flex-wrap: wrap;
justify-content: center;
gap: 30px;
margin: 40px 0;
}

.resume-desc select{
  color: rgba(75, 85, 99, 1);
  background-color: #e2e2e2;
  border-radius: 0;
  border-bottom: 2px solid #9b9b9b;
  outline: 0;
  margin-bottom: 10px;
}
</style>
