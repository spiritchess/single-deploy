const express = require('express');
const app = express();
const port = process.env.PORT ?? 3002;
const path = require('path');

app.use(express.static(path.resolve(__dirname, '../dist')));

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
