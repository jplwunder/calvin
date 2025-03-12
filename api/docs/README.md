# API Documentation UI

This directory contains files for the interactive API documentation UI.

## Files

- `index.html` - Main documentation page with Swagger UI for the OpenAPI specification
- `examples.html` - Page displaying example requests and responses
- `serve.js` - Simple Node.js server to serve the documentation locally

## Running the Documentation Server

To start the documentation server:

```bash
# Make sure you're in the api directory
cd api

# Run the server
./docs/serve.js
```

Then open your browser and go to [http://localhost:8000/](http://localhost:8000/)

## Features

The interactive documentation includes:

1. **OpenAPI Specification Viewer**
   - Interactive documentation of all API endpoints
   - Request and response schemas
   - Try it out functionality (when connected to a running API)

2. **Examples Page**
   - Example requests and responses for each endpoint
   - Syntax highlighted JSON
   - Both success and error examples

3. **Postman Collection**
   - Downloadable Postman collection for testing the API
