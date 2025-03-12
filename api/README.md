# Contacts Manager API Documentation

This directory contains the API documentation for the Contacts Manager service.

## Overview

The Contacts Manager API provides endpoints for managing contact information. It allows you to:

- Retrieve a list of all contacts
- Get details for a specific contact
- Create new contacts

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/healthcheck` | Check the health status of the API |
| GET | `/contacts` | List all contacts |
| GET | `/contacts/:id` | Get a specific contact by ID |
| POST | `/contacts` | Create a new contact |

## Data Models

### Contact

```json
{
  "id": "string",
  "name": "string",
  "phone": number,
  "email": "string"
}
```

## Examples

Example requests and responses can be found in the [examples](./examples) directory.

## OpenAPI Specification

The complete API specification is available in [openapi.yaml](./openapi.yaml).

## Postman Collection

A Postman collection is available for testing the API: [contacts-manager-postman.json](./contacts-manager-postman.json).

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
