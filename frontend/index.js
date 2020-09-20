const express = require('express')
const PORT = 5000

app = express()

app.use(express.static(__dirname + '/src'));

app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'src', 'index.html'));
});

app.listen(PORT, () => {
  console.log(`app listening to port: ${PORT}`);
})