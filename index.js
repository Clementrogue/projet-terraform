const express = require('express');
const app = express();

const port = process.env.APP_PORT || 3000;

app.get('/', (req, res) => {
  res.send('Student App is running');
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
