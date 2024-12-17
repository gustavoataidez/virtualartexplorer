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
            placeholder="Ex. Museu das Artes"
            required
          />
          <label for="description" class="form__label">Descrição</label>
          <textarea
            class="form__field texto"
            v-model="museum.description"
            placeholder="Descreva sobre o museu"
            name="description"
            required
          ></textarea>
          <label for="category1" class="form__label">Categoria 1</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.category1"
            placeholder="Ex. Cultura"
          />
          <label for="category2" class="form__label">Categoria 2</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.category2"
            placeholder="Ex. Tradições"
          />
          <label for="link" class="form__label">Link</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.link"
            placeholder="Ex. http://museucultura.com"
          />
          <label for="address" class="form__label">Endereço</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.address"
            placeholder="Ex. Avenida das Tradições, 303"
          />
          <label for="cep" class="form__label">CEP</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.cep"
            placeholder="Ex. 88900-111"
          />
          <label for="city" class="form__label">Cidade</label>
          <input
            type="input"
            class="form__field"
            v-model="museum.city"
            placeholder="Ex. Recife"
            required
          />
          <label for="state" class="form__label">Estado</label>
          <select
            class="form-select"
            style="background-color: #e2e2e2;"
            v-model="museum.state"
            name="state"
          >
            <option value="AC">Acre</option>
            <option value="AL">Alagoas</option>
            <option value="AP">Amapá</option>
            <option value="AM">Amazonas</option>
            <option value="BA">Bahia</option>
            <option value="CE">Ceará</option>
            <option value="DF">Distrito Federal</option>
            <option value="ES">Espírito Santo</option>
            <option value="GO">Goiás</option>
            <option value="MA">Maranhão</option>
            <option value="MT">Mato Grosso</option>
            <option value="MS">Mato Grosso do Sul</option>
            <option value="MG">Minas Gerais</option>
            <option value="PA">Pará</option>
            <option value="PB">Paraíba</option>
            <option value="PR">Paraná</option>
            <option value="PE">Pernambuco</option>
            <option value="PI">Piauí</option>
            <option value="RJ">Rio de Janeiro</option>
            <option value="RN">Rio Grande do Norte</option>
            <option value="RS">Rio Grande do Sul</option>
            <option value="RO">Rondônia</option>
            <option value="RR">Roraima</option>
            <option value="SC">Santa Catarina</option>
            <option value="SP">São Paulo</option>
            <option value="SE">Sergipe</option>
            <option value="TO">Tocantins</option>
          </select>
          <label for="information" class="form__label">Mais Informações</label>
          <textarea
            class="form__field texto"
            v-model="museum.information"
            placeholder="Ex. Inclui exposições sobre danças, culinária e festivais regionais."
          ></textarea>

          <label for="manager_id" class="form__label">ID do Gerente</label>
          <input
            type="number"
            class="form__field"
            v-model="museum.manager_id"
            placeholder="Ex. 6"
            disabled
          />
          
          <button class="btn btn-success" @click="updateMuseum">Salvar Museu</button>
        </div>
        <button class="btn btn-danger mt-2" @click="deactivateMuseum">Excluir</button>
      </div>
    </div>

    <!-- Seção de Adicionar Obras (já que o museu já está criado) -->
    <div class="add-obras-section" v-if="museumId">
      <h3>Adicionar Obras</h3>
      <button class="btn btn-primary" @click="openObraModal">Adicionar Obra</button>

      <!-- Listagem das obras adicionadas -->
      <ul class="obra-list" v-if="works.length > 0">
        <li v-for="(obra, index) in works" :key="obra.id">
          <strong>{{ obra.name }}</strong> - {{ obra.description }} - {{ obra.author || 'Autor Desconhecido' }} - {{ obra.image }}
          <a href="#" @click.prevent="deleteObra(obra.id, index)" style="color:red; margin-left:10px;">Excluir</a>
        </li>
      </ul>

      <!-- Botão para finalizar cadastro (direciona para a página home com alerta) -->
      <div class="finish-section">
        <button class="btn btn-success" @click="finishRegistration">Concluir</button>
      </div>
    </div>

    <!-- Modal para adicionar obra -->
    <div v-if="showObraModal" class="modal-overlay" @click="closeObraModal">
      <div class="modal-content" @click.stop>
        <h4>Adicionar Obra</h4>
        <label>Nome da Obra</label>
        <input type="text" v-model="newObra.name" placeholder="Ex. O Grito" required/>

        <label>Descrição</label>
        <textarea v-model="newObra.description" placeholder="Descrição da obra" required></textarea>

        <label>Autor</label>
        <input type="text" v-model="newObra.author" placeholder="Ex. Edvard Munch"/>

        <label>Link da Imagem</label>
        <input type="text" v-model="newObra.image" placeholder="URL da imagem"/>

        <button class="btn btn-success" @click="createObra">Salvar Obra</button>
        <button class="btn btn-danger" @click="closeObraModal">Fechar</button>
      </div>
    </div>

  </div>
  <FooterNovo></FooterNovo>
</template>

<script>
import HeaderNovo from "../HeaderPage.vue";
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
        state: "PE",
        information: "",
        manager_id: null,
        capa: null,
        image: ""
      },
      museumId: null,
      showObraModal: false,
      newObra: {
        name: "",
        description: "",
        author: "",
        image: ""
      },
      works: []
    };
  },
  async created() {
    this.museumId = this.$route.params.id;
    await this.fetchMuseum();
    await this.fetchArtworks();
  },
  methods: {
    onFileChange(event) {
      const file = event.target.files[0];
      this.museum.capa = file;
    },
    async deactivateMuseum() {
      try {
        const response = await fetch(`${API_URL}/museums/${this.museumId}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            title: "",
            description: "",
            image: "",
            category1: "",
            category2: "",
            link: "",
            address: "",
            cep: "",
            city: "",
            state: "",
            information: "",
            manager_id: null,
            active: false
          }),
        });

        if (response.ok) {
          alert("Museu desativado com sucesso!");
          this.$router.push("/museums");
        } else {
          alert("Erro ao desativar o museu.");
        }
      } catch (error) {
        console.error("Erro ao desativar:", error);
        alert("Erro ao desativar o museu.");
      }
    },
    async fetchMuseum() {
      try {
        const response = await fetch(`${API_URL}/museums/${this.museumId}`);
        if (!response.ok) {
          throw new Error("Erro ao buscar dados do museu");
        }
        const data = await response.json();
        this.museum = {
          ...this.museum,
          ...data
        };
      } catch (error) {
        console.error("Erro ao carregar dados do museu:", error);
      }
    },
    async fetchArtworks() {
      try {
        const response = await fetch(`${API_URL}/artworks?museum_id=${this.museumId}`);
        if (!response.ok) {
          throw new Error("Erro ao buscar obras de arte");
        }
        this.works = await response.json();
      } catch (error) {
        console.error("Erro ao carregar obras:", error);
      }
    },
    async updateMuseum() {
      if (!this.museum.title || !this.museum.description) {
        alert("Título e Descrição são obrigatórios.");
        return;
      }

      const museumData = {
        ...this.museum,
        image: this.museum.capa ? URL.createObjectURL(this.museum.capa) : this.museum.image,
      };

      try {
        const response = await fetch(`${API_URL}/museums/${this.museumId}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(museumData),
        });

        if (response.ok) {
          alert("Museu atualizado com sucesso!");
        } else {
          alert("Erro ao atualizar o museu.");
        }
      } catch (error) {
        console.error("Erro ao atualizar:", error);
        alert("Erro ao atualizar o museu.");
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
      if (!this.newObra.name || !this.newObra.description) {
        alert("Nome da Obra e Descrição são obrigatórios.");
        return;
      }

      const obraData = {
        museum_id: this.museumId,
        name: this.newObra.name,
        description: this.newObra.description,
        author: this.newObra.author,
        image: this.newObra.image || "",
        active: true
      };

      try {
        const response = await fetch(`${API_URL}/artworks`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(obraData),
        });

        if (response.ok) {
          const createdObra = await response.json();
          this.works.push(createdObra);
          this.closeObraModal();
        } else {
          alert("Erro ao cadastrar a obra.");
        }
      } catch (error) {
        console.error("Erro ao cadastrar obra:", error);
        alert("Erro ao cadastrar a obra.");
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
    finishRegistration() {
      alert("Museu atualizado com sucesso!");
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
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
  margin-top: 20px;
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
  background-color: rgba(0,0,0,0.5);
  display:flex;
  justify-content:center;
  align-items:center;
  z-index:9999;
}
.modal-content {
  background:#fff;
  padding:20px;
  border-radius:8px;
  max-width:400px;
  width:100%;
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
/* Estilização do modal */
.modal-overlay {
    height: 100%;
    width: 100%;
    background-color: rgba(0, 0, 0, .5);
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    z-index: 1;
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
  padding: 7px 10px;
  background-color: #e2e2e2;
  transition: border-color 0.2s;
  margin-bottom: 10px;
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
