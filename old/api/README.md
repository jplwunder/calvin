# Calvin API Documentation

This directory contains the API documentation for the Calvin API service.

## Overview

The Calvin API provides endpoints for managing customer information. It allows you to:

- Retrieve a list of all customers
- Get details for a specific customers
- Create new customers

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/healthcheck` | Check the health status of the API |
| GET | `/customers` | List all customers |
| GET | `/customers/:id` | Get a specific customer by ID |
| POST | `/customers` | Create a new customer |

## Data Models

### Customer

```json
{
  "id": "string",
  "name": "string",
  "phone": "number",
  "email": "string"
}
```

## Examples

Example requests and responses can be found in the [examples](./examples) directory.

## OpenAPI Specification

The complete API specification is available in [openapi.yaml](./openapi.yaml).

## Postman Collection

A Postman collection is available for testing the API: [calvin-postman.json](./calvin-postman.json).

## Interactive Documentation

The API documentation can be viewed in an interactive web interface. To start the documentation server:

```bash
# Navigate to the api directory
cd api

# Run the documentation server
./docs/serve.js
```

Then open your browser and go to [http://localhost:8000/](http://localhost:8000/)

The interactive documentation includes:
- OpenAPI specification with Swagger UI
- Example requests and responses
- Downloadable Postman collection
