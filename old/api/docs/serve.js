#!/usr/bin/env node

const http = require('http');
const fs = require('fs');
const path = require('path');

const PORT = 8000;
const API_DIR = path.join(__dirname, '..');

// MIME types for different file extensions
const MIME_TYPES = {
  '.html': 'text/html',
  '.css': 'text/css',
  '.js': 'text/javascript',
  '.json': 'application/json',
  '.yaml': 'text/yaml',
  '.yml': 'text/yaml',
  '.png': 'image/png',
  '.jpg': 'image/jpeg',
  '.jpeg': 'image/jpeg',
  '.gif': 'image/gif',
  '.svg': 'image/svg+xml',
  '.ico': 'image/x-icon',
};

// Create a simple HTTP server
const server = http.createServer((req, res) => {
  console.log(`${req.method} ${req.url}`);

  // Handle root URL
  let url = req.url;
  if (url === '/') {
    url = '/docs/index.html';
  }

  // Construct the file path
  const filePath = path.join(API_DIR, url);
  const extname = path.extname(filePath);
  const contentType = MIME_TYPES[extname] || 'text/plain';

  // Read the file
  fs.readFile(filePath, (err, content) => {
    if (err) {
      if (err.code === 'ENOENT') {
        // File not found
        res.writeHead(404);
        res.end('404 Not Found');
      } else {
        // Server error
        res.writeHead(500);
        res.end(`Server Error: ${err.code}`);
      }
    } else {
      // Success
      res.writeHead(200, { 'Content-Type': contentType });
      res.end(content, 'utf-8');
    }
  });
});

// Start the server
server.listen(PORT, () => {
  console.log(`Documentation server running at http://localhost:${PORT}/`);
  console.log(`Press Ctrl+C to stop the server`);
});
