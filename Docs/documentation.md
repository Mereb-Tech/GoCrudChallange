# Go CRUD Challenge Documentation

## Overview

This document outlines the implementation of a simple CRUD (Create, Read, Update, Delete) API for managing person records using Go. The API utilizes an in-memory database and adheres to clean architecture principles.

## API Endpoints

### Base URL




### Endpoints

1. **GET** `/person`
   - **Description:** Retrieve all persons.
   - **Response:**
     - **200 OK**: Returns a list of persons.
     - **Example Response:**
       ```json
       [
         {
           "id": "uuid-1",
           "name": "John Doe",
           "age": 30,
           "hobbies": ["reading", "swimming"]
         },
         {
           "id": "uuid-2",
           "name": "Jane Smith",
           "age": 25,
           "hobbies": []
         }
       ]
       ```

2. **GET** `/person/${personId}`
   - **Description:** Retrieve a specific person by `personId`.
   - **Response:**
     - **200 OK**: Returns the person object if found.
     - **404 Not Found**: If the person does not exist.
     - **Example Response:**
       ```json
       {
         "id": "uuid-1",
         "name": "John Doe",
         "age": 30,
         "hobbies": ["reading", "swimming"]
       }
       ```

3. **POST** `/person`
   - **Description:** Create a new person record.
   - **Request Body:**
     - **Required Fields:** `name` (string), `age` (number), `hobbies` (array of strings).
     - **Example Request:**
       ```json
       {
         "name": "Alice Johnson",
         "age": 28,
         "hobbies": ["hiking", "cooking"]
       }
       ```
   - **Response:**
     - **201 Created**: Returns the created person object.
     - **400 Bad Request**: If required fields are missing.
     - **Example Response:**
       ```json
       {
         "id": "uuid-3",
         "name": "Alice Johnson",
         "age": 28,
         "hobbies": ["hiking", "cooking"]
       }
       ```

4. **PUT** `/person/${personId}`
   - **Description:** Update an existing person record.
   - **Request Body:**
     - **Required Fields:** `name` (string), `age` (number), `hobbies` (array of strings).
     - **Example Request:**
       ```json
       {
         "name": "Alice Johnson",
         "age": 29,
         "hobbies": ["hiking", "traveling"]
       }
       ```
   - **Response:**
     - **200 OK**: Returns the updated person object.
     - **404 Not Found**: If the person does not exist.
     - **400 Bad Request**: If required fields are missing.
     - **Example Response:**
       ```json
       {
         "id": "uuid-3",
         "name": "Alice Johnson",
         "age": 29,
         "hobbies": ["hiking", "traveling"]
       }
       ```

5. **DELETE** `/person/${personId}`
   - **Description:** Delete a person record.
   - **Response:**
     - **200 OK**: If deletion is successful.
     - **404 Not Found**: If the person does not exist.

### Error Handling

- **404 Not Found**: Returned for requests to non-existing endpoints (e.g., `/some-non/existing/resource`).

### CORS Support

The API is configured to allow cross-origin requests, enabling accessibility by frontend applications hosted on different domains.

## Data Model

### Person Object

- **id**: Unique identifier (`string`, `uuid`) generated on the server side.
- **name**: Person's name (`string`, **required**).
- **age**: Person's age (`number`, **required**).
- **hobbies**: Person's hobbies (`array` of `strings`, **required**).


## Testing

- Unit tests are written to ensure the functionality of the API and the correctness of the CRUD operations.
- Tests can be found in the `Test` directory.

## Conclusion

This documentation serves as a guide to understanding and using the Person API built in Go. 


