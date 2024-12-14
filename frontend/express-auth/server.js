const express = require("express");
const cors = require("cors");
const jwt = require("jsonwebtoken");
const bodyParser = require("body-parser");

const app = express();

// Configurando o middleware CORS
app.use(cors());
app.use(bodyParser.json());

const users = [
  {
    id: 1,
    email: "user@example.com",
    password: "123456",
    active: true,
    first_name: "Teste",
    last_name: "User",
  },
  {
    id: 2,
    email: "john.doe@example.com",
    password: "12345",
    active: true,
    first_name: "John",
    last_name: "Doe",
  },
  {
    id: 3,
    email: "jane.smith@example.com",
    password: "12345",
    active: false,
    first_name: "Jane",
    last_name: "Smith",
  },
  {
    id: 4,
    email: "admin.user@example.com",
    password: "12345",
    active: true,
    first_name: "Admin",
    last_name: "User",
  },
];

app.post("/auth/login", (req, res) => {
  const { email, password } = req.body;

  const user = users.find((u) => u.email === email && u.password === password);

  if (user) {
    if (!user.active) {
      return res.status(403).json({ message: "Conta inativa. Entre em contato com o suporte." });
    }
    const token = jwt.sign(
      {
        id: user.id,
        email: user.email,
        first_name: user.first_name,
        last_name: user.last_name,
      },
      "secretKey",
      { expiresIn: "1h" }
    );
    res.json({ token, first_name: user.first_name, last_name: user.last_name, id: user.id });
  } else {
    res.status(401).json({ message: "Credenciais inválidas" });
  }
});

app.listen(5500, () => {
  console.log("Server running on http://localhost:5500");
});
