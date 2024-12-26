# Virtual Art Explore
<img src="https://raw.githubusercontent.com/gustavoataidez/visualartexplorer/0b6b0fbffa72512b2b12d732168a04b6737f1b53/frontend/src/assets/image.png" alt="Print Home" style="width: 100%">

### Acesse a partir do link: [visualartexplorer.vercel.app](https://visualartexplorer.vercel.app/) 

## 📚 Descrição

**Virtual Art Explore** é uma plataforma online dedicada à visualização e gerenciamento de museus e suas obras de arte. O projeto visa facilitar o acesso a informações culturais, permitindo que visitantes encontrem facilmente museus, explorem suas coleções e se informem sobre eventos e exposições. Além disso, oferece uma área administrativa para gerentes de museus atualizarem e gerenciarem suas informações de forma eficiente.

## 🌐 Visualização

O projeto está disponível para visualização através do seguinte link:

Frontend: [https://visualartexplorer.vercel.app/](https://visualartexplorer.vercel.app/)

Backend: [https://visualartexplorer.onrender.com](https://visualartexplorer.onrender.com)

## 👥 Integrantes

- **Luis Gustavo Souza Ataide**
- **André Luiz Cardoso da Silva**
  
## 🛠️ Tecnologias Utilizadas

### **Backend**
- **Golang:** Linguagem de programação responsável pela API, oferecendo um ambiente limpo, rápido e escalável.
- **GoGin:** Framework web utilizado para facilitar a implementação do servidor, das rotas e dos controladores.
- **Viper:** Utilizado para carregar as configurações do projeto.
- **Postgres:** Banco de dados relacional escolhido para armazenar as informações.
- **Testify:** Framework de testes para Golang.

### **Frontend**
- **HTML/CSS:** Estruturação e estilização das páginas.
- **JavaScript:** Tornar as páginas interativas e dinâmicas.
- **Vue.js:** Framework JavaScript progressivo para construir interfaces reativas e gerenciar o estado da aplicação.
- **Bootstrap:** Framework CSS para estilização rápida e responsiva, oferecendo componentes prontos e um sistema de grid para layout consistente.
- **Vite:** Ferramenta de bundling e build, proporcionando um ambiente de desenvolvimento ágil com recarregamento instantâneo e otimizações para produção.

### **Ferramentas de Desenvolvimento**
- **VSCode:** Ambiente de desenvolvimento integrado (IDE) utilizado para escrever e editar o código do projeto, com suporte a diversas extensões.
- **Goland:** IDE específica para desenvolvimento em Golang, oferecendo recursos avançados para programação, depuração e testes.
- **Postman:** Ferramenta para testar e documentar a API, permitindo o envio de requisições HTTP e verificação de respostas.
- **PGAdmin:** Interface gráfica para gerenciar o banco de dados PostgreSQL, utilizada para criar, alterar e consultar tabelas, além de administrar usuários e permissões.

## 🚀 Como Executar o Projeto

### Pré-requisitos

- **Git** para clonar o repositório.

### Passo a Passo

1. **Clone o repositório:**

```
git clone https://github.com/seu-usuario/virtual-art-explore.git
cd virtual-art-explore
```

## Principais Endpoints

### Gerenciamento de Gerentes:
POST /api/v1/managers: Cadastro de novo gerente.

PUT /api/v1/managers/:id: Atualizar informações do gerente.

DELETE /api/v1/managers/:id/disable: Desativar gerente.

### Autenticação:

POST /api/v1/login: Login e obtenção de JWT.

### Museus:

GET /api/v1/museums: Listar todos os museus.

GET /api/v1/museums/city/:city: Filtrar museus por cidade.

GET /api/v1/museums/state/:state: Filtrar museus por estado.

GET /api/v1/museums/name/:name: Filtrar museus por nome.

POST /api/v1/museums: Cadastrar novo museu (Autenticado).

PUT /api/v1/museums/:id: Atualizar museu (Autenticado).

DELETE /api/v1/museums/:id/disable: Desativar museu (Autenticado).

### Obras de Arte:

GET /api/v1/artworks: Listar todas as obras.

GET /api/v1/artworks/museum/:name: Filtrar obras por nome do museu.

GET /api/v1/artworks/author/:author: Filtrar obras por autor.

GET /api/v1/artworks/name/:name: Filtrar obras por nome.

GET /api/v1/artworks/year/:year: Filtrar obras por ano.

POST /api/v1/artworks: Cadastrar nova obra (Autenticado).

PUT /api/v1/artworks/:id: Atualizar obra (Autenticado).

DELETE /api/v1/artworks/:id/disable: Desativar obra (Autenticado).


