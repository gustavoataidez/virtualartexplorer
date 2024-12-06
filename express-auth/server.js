const express = require("express");
const cors = require("cors");
const jwt = require("jsonwebtoken");
const bodyParser = require("body-parser");

const app = express();

// Configurando o middleware CORS
app.use(cors());
app.use(bodyParser.json());

const users = [
  { id: 1, email: "user@example.com", password: "123456" },
];

app.post("/auth/login", (req, res) => {
  const { email, password } = req.body;

  const user = users.find((u) => u.email === email && u.password === password);

  if (user) {
    const token = jwt.sign({ id: user.id, email: user.email }, "secretKey", {
      expiresIn: "1h",
    });
    res.json({ token });
  } else {
    res.status(401).json({ message: "Credenciais inválidas" });
  }
});

app.listen(5500, () => {
  console.log("Server running on http://localhost:5500");
});
