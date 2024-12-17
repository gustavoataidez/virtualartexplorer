<template>
  <HeaderNovo></HeaderNovo>
  <div class="museu-page">
    <div class="sec-resume">
      <div>
        <p class="form__label">Capa</p>
        <label class="custum-file-upload" for="file">
          <div class="icon">
            <svg xmlns="http://www.w3.org/2000/svg" fill="" viewBox="0 0 24 24">
              <path
                fill=""
                d="M10 1C9.73478 1 9.48043 1.10536 9.29289 1.29289L3.29289 7.29289C3.10536 7.48043 3 7.73478 3 8V20C3 21.6569 4.34315 23 6 23H7C7.55228 23 8 22.5523 8 22C8 21.4477 7.55228 21 7 21H6C5.44772 21 5 20.5523 5 20V9H10C10.5523 9 11 8.55228 11 8V3H18C18.5523 3 19 3.44772 19 4V9C19 9.55228 19.4477 10 20 10C20.5523 10 21 9.55228 21 9V4C21 2.34315 19.6569 1 18 1H10ZM9 7H6.41421L9 4.41421V7ZM14 15.5C14 14.1193 15.1193 13 16.5 13C17.8807 13 19 14.1193 19 15.5V16V17H20C21.1046 17 22 17.8954 22 19C22 20.1046 21.1046 21 20 21H13C11.8954 21 11 20.1046 11 19C11 17.8954 11.8954 17 13 17H14V16V15.5ZM16.5 11C14.142 11 12.2076 12.8136 12.0156 15.122C10.2825 15.5606 9 17.1305 9 19C9 21.2091 10.7909 23 13 23H20C22.2091 23 24 21.2091 24 19C24 17.1305 22.7175 15.5606 20.9844 15.122C20.7924 12.8136 18.858 11 16.5 11Z"
                clip-rule="evenodd"
                fill-rule="evenodd"
              ></path>
            </svg>
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
          />
          <button class="btn btn-success" @click="createMuseum">Salvar</button>
        </div>
      </div>
    </div>
    <FooterNovo></FooterNovo>
  </div>
</template>

<script>
import HeaderNovo from "../HeaderPage.vue";
import FooterNovo from "../Footer.vue";
import { API_URL } from '@/config'; // Importe a constante da configuração

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
      },
    };
  },
  methods: {
    onFileChange(event) {
      const file = event.target.files[0];
      this.museum.capa = file;
    },
    async createMuseum() {
      const museumData = {
        ...this.museum,
        image: this.museum.capa ? URL.createObjectURL(this.museum.capa) : "",
      };

      try {
        const response = await fetch(`${API_URL}/museums`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(museumData),
        });

        if (response.ok) {
          alert("Museu cadastrado com sucesso!");
          this.resetForm();
        } else {
          alert("Erro ao cadastrar o museu.");
        }
      } catch (error) {
        console.error("Erro ao cadastrar:", error);
        alert("Erro ao cadastrar o museu.");
      }
    },
    resetForm() {
      this.museum = {
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
      };
    },
  },
};
</script>



<style scoped>
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
