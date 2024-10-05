# Project Structure Overview

```
.
├── api
│   ├── controller
│   │   ├── base.go          // Common controller methods for handling HTTP responses.
│   │   ├── dtos.go          // Data Transfer Objects (DTOs) for request and response formats.
│   │   ├── icontroller.go    // Interface for controller methods to register routes.
│   │   └── person.go        // Controller handling person-related HTTP requests.
│   ├── error
│   │   └── api_error.go     // Structured error handling for the API layer.
│   └── router
│       └── router.go        // HTTP server and routing setup, managing different access levels.
├── app
│   ├── common
│   │   ├── cqrs
│   │   │   └── cqrs.go      // CQRS pattern implementation; defines command/query handler interfaces.
│   │   └── interface
│   │       └── repo
│   │           └── person.go // Repository interface for person data operations.
│   └── person
│       ├── commands
│       │   ├── create.go    // Command handling for creating a person.
│       │   ├── delete.go    // Command handling for deleting a person.
│       │   └── update.go    // Command handling for updating a person.
│       └── query
│           ├── get_all.go   // Query handling for retrieving all persons.
│           └── get.go       // Query handling for retrieving a person by ID.
├── cmd
│   └── main.go              // Main entry point for the application.
├── config
│   └── env.go              // Configuration management, loading environment variables.
│   └── API_DEFINITION.md    // API definitions and documentation.
├── domain
│   ├── common
│   │   └── error.go         // Common error types and handling in the domain layer.
│   ├── error
│   │   └── domain_error.go   // Domain-specific error definitions.
│   └── models
│       └── person.go        // Domain model representing a person.
├── infrastructure
│   └── repo
│       └── person.go        // Repository implementation for person data operations.
```

#### `api` Package

The `api` package handles the HTTP interface for the application.

- **controller**

  - **base.go**: Contains common methods for handling HTTP responses, including standardizing error responses and formatting successful results.
  - **dtos.go**: Defines Data Transfer Objects (DTOs) that serve as structured representations for request and response data.
  - **icontroller.go**: Interface defining methods for registering routes at different access levels (public, protected, privileged).
  - **person.go**: Implements the `PersonController`, which handles person-related HTTP requests such as creating, updating, deleting, and retrieving persons.

- **error**

  - **api_error.go**: Implements structured error handling specific to the API layer, providing different types of errors (bad request, not found, etc.) with associated HTTP status codes.

- **router**
  - **router.go**: Responsible for setting up and running the HTTP server. It manages routes, middleware, and different access levels for public and protected routes.

#### `app` Package

The `app` package contains the core business logic and application structure.

- **common**

  - **cqrs**
    - **cqrs.go**: Implements the Command Query Responsibility Segregation (CQRS) pattern. It defines interfaces for command and query handlers, enabling separation of responsibilities.
  - **interface/repo**
    - **person.go**: Defines the repository interface for person data operations, ensuring consistent access patterns.

- **person**
  - **commands**
    - **create.go**: Implements the command handling logic for creating a person.
    - **delete.go**: Implements the command handling logic for deleting a person.
    - **update.go**: Implements the command handling logic for updating a person's details.
  - **query**
    - **get_all.go**: Implements the query handling logic for retrieving all persons.
    - **get.go**: Implements the query handling logic for retrieving a person by their unique ID.

#### `config` Package

- **env.go**: Manages application configuration by loading environment variables. It defines defaults and retrieves configuration values needed for server setup.

#### `domain` Package

The `domain` package contains the core models and error handling for the business logic.

- **common**

  - **error.go**: Defines common error types used throughout the domain layer, facilitating error management and consistency.

- **error**

  - **domain_error.go**: Contains domain-specific error definitions, providing a structured way to handle errors in the business logic.

- **models**
  - **person.go**: Defines the `Person` model, representing the person entity with fields such as ID, name, age, and hobbies.

#### `infrastructure` Package

- **repo**
  - **person.go**: Contains the repository implementation for managing person data operations. This includes functions for CRUD operations interacting with a data store.

### Summary

This Go project implements a RESTful API for managing persons using clean architecture principles and CQRS.
