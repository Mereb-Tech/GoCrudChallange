# Go CRUD Application 

This project is a CRUD application in Go that manages a "Person" entity using Go’s built-in `net/http` package, following a hexagonal architecture to maintain separation of concerns and scalability. The application is hosted on a VPS with a domain at [merebcrud.pro.et](https://merebcrud.pro.et), and it leverages several technologies to ensure efficient development, deployment, monitoring, and secure HTTPS handling.

## Table of Contents
- [Architecture](#architecture)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
  - [Run with Docker](#run-with-docker)
  - [Prometheus Monitoring](#prometheus-monitoring)
  - [CI/CD Workflow](#cicd-workflow)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [File Structure](#file-structure)

---

## Architecture

This application follows a **Hexagonal Architecture** (also known as ports and adapters architecture). This approach isolates core logic from external components like databases and HTTP handlers, making it modular, testable, and flexible. 

**Why Hexagonal Architecture?**
- Separation of concerns for ease of testing and maintenance.
- Flexibility to replace or upgrade components (e.g., swap database, replace HTTP server).
- Clearly defined contracts and dependencies, improving code readability and reducing coupling.

---

## Features

- CRUD operations on a `Person` entity (Create, Read, Update, Delete).
- Endpoint monitoring via Prometheus, with metrics for total request counts on each endpoint.
- Secure HTTPS setup using **Caddy** as a reverse proxy, managing TLS certificates.
- Continuous Integration and Deployment with **GitHub Actions**.
- Fully containerized environment with Docker and Docker Compose.

---

## Technologies Used

- **Go**: Core language for building the application, using `net/http` for HTTP handling.
- **Docker & Docker Compose**: For containerization of the application and Prometheus.
- **Caddy**: Reverse proxy for HTTPS/TLS management.
- **Prometheus**: Monitoring tool, used here to count total HTTP requests to each endpoint.
- **GitHub Actions**: Automates testing and deployment as part of the CI/CD pipeline.
- **VPS**: Hosted on a Virtual Private Server, available at [merebcrud.pro.et](https://merebcrud.pro.et).

---

## Setup

### Run with Docker

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Mahider-T/GoCrudChallange.git
   cd GoCrudChallange
   ```

2. **Build and start the services**:
   ```bash
   docker-compose up --build
   ```

   This command builds and runs the application in a Docker container along with Prometheus for monitoring.

3. **Access the application**:
   - Application endpoints: `http://localhost:3000`
   - Prometheus: `http://localhost:9090`

---

### Prometheus Monitoring

Prometheus is set up to monitor HTTP requests to the application’s endpoints. Metrics are exposed at `/metrics` and include counters for total requests per endpoint.

**Custom Metrics**:
- `http_requests_total`: Counts the total number of HTTP requests, labeled by method and endpoint.

### CI/CD Workflow

A GitHub Actions workflow file is included in the repository to automate deployment. Each push to the repository triggers deployment to the VPS if tests pass.

---

## API Endpoints

All endpoints manage `Person` data, which includes fields like `ID`, `Name`, `Age`, and `City`.

| Method | Endpoint                | Description            |
|--------|--------------------------|------------------------|
| POST   | `/person`               | Creates a new person   |
| GET    | `/person`               | Retrieves all persons  |
| GET    | `/person/{personId}`    | Retrieves a person by ID |
| PUT    | `/person/{personId}`    | Updates a person by ID |
| DELETE | `/person/{personId}`    | Deletes a person by ID |
| GET    | `/metrics`              | Exposes Prometheus metrics |

---

## Testing

This application includes tests for handlers to validate CRUD functionality. The tests are written in Go and check if the application responds as expected for each API endpoint.

To run the tests:
```bash
go test ./cmd/handlers_test.go
```

---

## File Structure

The project structure follows a modular organization for clarity and scalability.

```
├── Api Documentation.json
├── Caddyfile
├── Dockerfile
├── README.md
├── cmd
│   ├── handlers_test.go
│   └── main.go
├── config
│   └── conf.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── internal
    ├── adapters
    │   ├── db
    │   │   └── person.db.go
    │   └── handlers
    │       ├── jsonify.helper.go
    │       └── person.handler.go
    ├── application
    │   └── core
    │       ├── api
    │       │   └── person.api.go
    │       └── domain
    │           ├── errors.go
    │           └── person.domain.go
    ├── assert
    │   └── assert.go
    ├── pkg
    └── ports
        ├── apiport
        │   └── person.apiPort.go
        └── dbport
            └── person.dbPort.go
```

---

## Additional Notes

This application is designed to be lightweight, secure, and scalable, using minimal dependencies while maximizing performance and maintainability. 

