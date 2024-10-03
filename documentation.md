# Project Structure
**controller/**: Contains a single file for handling HTTP requests and responses.
**router/**: Contains a single file for defining routes and linking them to controllers.

**services**/: Contains a single file for business logic and interactions with the data layer.

**data**/: Acts as a mock database for your CRUD application, in memory storage is used 

## Setup Instructions

### Create a .env file

This file will store environment variables, such as the port number(use any port number) but 8080 will be used by default if port number is not provided 

```env
PORT=8080
```

# Install Dependencies
1. Ensure you have Go installed.

2. Navigate to your project directory in the terminal.
Run the following command to initialize and install dependencies:
```bash
go mod tidy
```

# Run the Application
To run the application, execute the following command:
```bash
go run main.go
```

