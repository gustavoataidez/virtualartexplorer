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

          <div class="category-section">
            <div class="category-field">
              <label for="category1" class="form__label">Categoria 1</label>
              <select class="form-select" v-model="museum.category1">
                <option disabled value="">Selecione uma categoria</option>
                <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
              </select>
            </div>

            <div class="category-field">
              <label for="category2" class="form__label">Categoria 2</label>
              <select class="form-select" v-model="museum.category2">
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

          <label for="state" class="form__label">Estado</label>
          <select
            class="form-select"
            v-model="museum.state"
            name="state"
            @change="onStateChange"
          >
            <option disabled value="">Selecione um estado</option>
            <option v-for="est in states" :key="est.uf" :value="est.uf">{{ est.nome }}</option>
          </select>

          <label for="city" class="form__label">Cidade</label>
          <select
            class="form-select"
            v-model="museum.city"
            name="city"
          >
            <option disabled value="">Selecione uma cidade</option>
            <option v-for="c in cities" :key="c" :value="c">{{ c }}</option>
          </select>

          <label for="information" class="form__label">Mais Informações</label>
          <textarea
            class="form__field texto"
            v-model="museum.information"
            placeholder="Ex. Inclui exposições sobre danças, culinária e festivais regionais."
          ></textarea>

          <label for="manager_id" class="form__label disabled-field">ID do Gerente</label>
          <input
            type="number"
            class="form__field disabled-field"
            v-model="museum.manager_id"
            placeholder="Ex. 6"
            disabled
          />

        </div>
      </div>
    </div>

    <!-- Componente para adicionar obras -->
<ArtworkCreate v-if="museumId" :museumId="museumId" />

<div class="final-actions" style="text-align:center; margin-top:20px;">
  <button class="btn btn-success" @click="updateMuseum">Salvar Museu</button>
  <button class="btn btn-danger" style="margin-left:10px;" @click="deactivateMuseum">Excluir</button>
</div>
</div>
<FooterNovo></FooterNovo>
</template>

<script>
import HeaderNovo from "../Header.vue";
import FooterNovo from "../Footer.vue";
import ArtworkCreate from "./ArtworkCreate.vue";
import { API_URL } from "@/config";

export default {
components: { HeaderNovo, FooterNovo, ArtworkCreate },
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
    image: "",
  },
  museumId: null,
  states: [],
  cities: [],
};
},
async created() {
this.museumId = parseInt(this.$route.params.id, 10);
await this.fetchMuseum();
await this.fetchStates();
},
methods: {
async fetchMuseum() {
  try {
    const response = await fetch(`${API_URL}/museums/${this.museumId}`);
    if (!response.ok) throw new Error("Erro ao buscar dados do museu");
    const data = await response.json();
    this.museum = { ...this.museum, ...data };
    if (this.museum.state) await this.onStateChange();
  } catch (error) {
    console.error("Erro ao carregar dados do museu:", error);
  }
},
async fetchStates() {
  try {
    const res = await fetch("https://servicodados.ibge.gov.br/api/v1/localidades/estados");
    const data = await res.json();
    this.states = data.map((state) => ({ uf: state.sigla, nome: state.nome }));
  } catch (error) {
    console.error("Erro ao carregar estados:", error);
  }
},
async onStateChange() {
  if (!this.museum.state) return;
  try {
    const res = await fetch(`https://servicodados.ibge.gov.br/api/v1/localidades/estados/${this.museum.state}/municipios`);
    const data = await res.json();
    this.cities = data.map((city) => city.nome);
    this.museum.city = "";
  } catch (error) {
    console.error("Erro ao carregar cidades:", error);
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
      headers: { "Content-Type": "application/json" },
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
async deactivateMuseum() {
  try {
    const response = await fetch(`${API_URL}/museums/${this.museumId}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ active: false }),
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
  padding: 20px;
  border-radius:20px;
  box-shadow: 0 0 100px rgba(0,0,0,0.7);
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
  margin-top: 20px;
}

.obra-list {
  list-style: none;
  padding: 0;
  margin: 20px 0;
}
.obra-list li {
  margin-bottom: 10px;
}

.modal-content {
  background:#fff;
  padding:20px;
  border-radius:20px;
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
