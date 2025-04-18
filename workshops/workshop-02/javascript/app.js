// Node.js health check example using Express
const express = require('express');
const fs = require('fs');
const app = express();
const port = 80;

app.get('/', (req, res) => {
  res.send('Hello, this is code from Node.js!');
});

function isFileExists(filePath) {
  return fs.existsSync(filePath);
}

app.get('/health/readiness', (req, res) => {
  if (isFileExists('/tmp/ready')) {
    res.writeHead(200, { 
      "Content-Type": "application/json",
      "Description": "200 OK, it's ready!",
     });
    res.end("200 OK, it's ready!");
  } 
  else {
    res.writeHead(503, { 
      "Content-Type": "application/json",
      "Description": "503 Service unavailable, it's not ready!",
     });
    res.end("503 Service unavailable, it's not ready!");
  }
});

app.get('/health/liveness', (req, res) => {
  if (isFileExists('/tmp/ready')) {
    res.writeHead(200, { 
      "Content-Type": "application/json",
      "Description": "200 OK, it lives!",
     });
    res.end("200 OK, it lives!");
  } 
  else {
    res.writeHead(503, { 
      "Content-Type": "application/json",
      "Desscription": "503 Service unavailable, it doesn't live!",
    });
    res.end("503 Service unavailable, it doesn't live!");
  }
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
