<template>
    <div>
      <h3>Adicionar Obras</h3>
      <button class="btn btn-primary" @click="openObraModal">Adicionar Obra</button>
      <ul class="obra-list" v-if="works.length > 0">
        <li v-for="(obra, index) in works" :key="obra.id">
          <strong>{{ obra.name }}</strong> - {{ obra.description }} - {{ obra.author || 'Autor Desconhecido' }}
          <a href="#" @click.prevent="deleteObra(obra.id, index)" style="color:red; margin-left:10px;">Excluir</a>
        </li>
      </ul>
      <div v-if="showObraModal" class="modal-overlay" @click="closeObraModal">
        <div class="modal-content" @click.stop>
          <h4>Adicionar Obra</h4>
          <label>Nome da Obra</label>
          <input type="text" v-model="newObra.name" placeholder="Ex. O Grito" required />
          <label>Descrição</label>
          <textarea v-model="newObra.description" placeholder="Descrição da obra" required></textarea>
          <label>Autor</label>
          <input type="text" v-model="newObra.author" placeholder="Ex. Edvard Munch" />
          <label>Link da Imagem</label>
          <input type="text" v-model="newObra.image" placeholder="URL da imagem" />
          <button class="btn btn-success my-2" @click="createObra">Salvar Obra</button>
          <button class="btn btn-danger" @click="closeObraModal">Fechar</button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  
import { API_URL } from "@/config";

  export default {
    props: {
      museumId: {
        type: Number,
        required: true, // `museumId` é obrigatório
      },
    },
    data() {
      return {
        works: [],
        newObra: {
          name: "",
          description: "",
          author: "",
          image: null,
          active: true,
        },
        showObraModal: false,
      };
    },
    methods: {
      openObraModal() {
        this.showObraModal = true;
      },
      closeObraModal() {
        this.showObraModal = false;
        this.newObra = { name: "", description: "", author: "", image: null, active: true };
      },
      async createObra() {
        if (!this.newObra.name || !this.newObra.description) {
          alert("Nome e Descrição são obrigatórios.");
          return;
        }
  
        const obraData = {
          ...this.newObra,
          museum_id: this.museumId,
        };
  
        try {
          const token = this.$store.state.token;
          const response = await fetch(`${API_URL}/artworks`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify(obraData),
          });
  
          if (response.ok) {
            const createdObra = await response.json();
            this.works.push(createdObra);
            alert("Obra cadastrada com sucesso!");
            this.closeObraModal();
          } else {
            const errorResponse = await response.json();
            console.error("Erro ao cadastrar obra:", errorResponse);
            alert(`Erro ao cadastrar obra: ${errorResponse.message || "Erro desconhecido"}`);
          }
        } catch (error) {
          console.error("Erro ao cadastrar obra:", error);
          alert("Erro ao cadastrar a obra.");
        }
      },
      async deleteObra(id, index) {
        try {
          const response = await fetch(`${API_URL}/artworks/${id}`, {
            method: "DELETE",
            headers: {
              Authorization: `Bearer ${this.$store.state.token}`,
            },
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
