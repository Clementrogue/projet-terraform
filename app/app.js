const express = require('express');
const app = express();
const port = 80;

app.get('/', (req, res) => {
  res.send('<h1>Bienvenue sur la Phase 4 ! 🐳 L\'application tourne dans un conteneur Docker !</h1>');
});

// Route indispensable pour le Health Check du Load Balancer
app.get('/health', (req, res) => res.status(200).send('OK'));

app.listen(port, () => {
  console.log(`Serveur démarré sur le port ${port}`);
});
