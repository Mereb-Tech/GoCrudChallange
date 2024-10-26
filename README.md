```markdown
# Go CRUD API

This project is a simple RESTful API for managing "Person" records, built using Go and designed with a layered architecture (Controller, UseCase, Repository). The API performs basic CRUD operations, including creating, reading, updating, and deleting person records. This project also includes end-to-end tests for each layer using `testify` and `httptest` packages.

## Features

- CRUD Operations on Person records
- Layered Architecture: Controller, UseCase, Repository
- Field validation for Person attributes
- Comprehensive testing for each layer
- CORS enabled for frontend integration

## Project Structure

```plaintext
GoCrudChallenge/
├── Controller      # Handles incoming HTTP requests and responses
├── Domain
│   ├── DTOs        # Data Transfer Objects for request and response formats
│   ├── Interfaces  # Interfaces for decoupling components
│   └── Models      # Data models for the application
├── Repository      # Manages data persistence
├── Routers         # Routes definitions for the API
├── UseCases        # Business logic for handling operations
├── tests           # Test suite for each layer
├── go.mod          # Go module definition
└── main.go         # Application entry point
```

## Setup

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or later
- [Gin Gonic](https://github.com/gin-gonic/gin) for routing
- [Testify](https://github.com/stretchr/testify) for testing

### Installation

1. Clone the repository

2. Install dependencies:

    ```bash
    go mod tidy
    ```

### Environment Variables

Configure your environment variables in a `.env` file as needed for additional settings. The following variables are available:
 ```plaintext
 BASE_URL="your base url"
 ```
## Usage

### Running the API

Start the server with:

```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

### API Endpoints

| Method | Endpoint       | Description                     |
|--------|----------------|---------------------------------|
| POST   | `/person`      | Create a new person            |
| GET    | `/person`      | Get all persons                |
| GET    | `/person/:id`  | Get a specific person by ID    |
| PUT    | `/person/:id`  | Update a specific person by ID |
| DELETE | `/person/:id`  | Delete a specific person by ID |

### Example `Person` Model

```json
{
  "id": "UUID",
  "name": "string",
  "age": "int",
  "hobbies": ["string", "string"]
}
```

### Data Validation

- **Name**: Required (string)
- **Age**: Required, must be >= 0
- **Hobbies**: Required, non-empty array

### Error Handling

- **400**: Bad Request – invalid data format
- **404**: Not Found – person not found
- **500**: Internal Server Error – server issues

## Testing

### Running Tests

End-to-end tests are located in the `tests` directory. These include tests for each layer: Repository, UseCase, and Controller. Tests are implemented using `testify/suite` and `httptest`.

Run the tests with:

```bash
go test -v ./tests
```

### Test Suite Overview

- **Repository Layer**: Verifies CRUD functionality for the `PersonRepository`.
- **UseCase Layer**: Tests business logic and validation for person creation, update, and deletion.
- **Controller Layer**: End-to-end tests for API endpoints, mocking dependencies using `httptest`.

## Project Highlights

- **Clean Architecture**: Decouples business logic, data handling, and HTTP handling.
- **Data Validation**: Ensures data integrity and prevents invalid data from entering the database.
- **Test Coverage**: Provides a robust test suite for thorough validation of each layer.
- **CORS**: Allows frontend applications on different domains to access the API.

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Testify for Testing](https://github.com/stretchr/testify)
- [uuid for Unique Identifiers](https://github.com/google/uuid)


## Postman collection

[Postman collection](https://documenter.getpostman.com/view/32287741/2sAY4sjQKN "Postman collection")

```
```
``` 

