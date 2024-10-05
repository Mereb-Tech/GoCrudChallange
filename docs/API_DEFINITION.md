# API Definition

## Endpoints

### 1. Create Person

- **Endpoint**: `{host}/api/v1/person`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "name": "string (min: 4, max: 100)",
    "age": "int8 (min: 0)",
    "hobbies": "[]string (optional)"
  }
  ```
- **Response**:

  - **201 Created**:

    - **Body**:

    ```json
    {
      "id": "uuid",
      "name": "string",
      "age": "int8",
      "hobbies": "[]string"
    }
    ```

    - **Headers**:
      Location: URI of the newly created person

  - **400 Bad Request**: If input validation fails.

---

### 2. Get All Persons

- **Endpoint**: `{host}/api/v1/person`
- **Method**: `GET`
- **Response**:
  - **200 OK**
    ```json
    [
      {
        "id": "uuid",
        "name": "string",
        "age": "int8",
        "hobbies": "[]string"
      },
      ...
    ]
    ```

---

### 3. Get Person by ID

- **Endpoint**: `{host}/api/v1/person/:id`
- **Method**: `GET`
- **Path Parameters**:
  - `id` (uuid): The ID of the person to retrieve.
- **Response**:
  - **200 OK**
    ```json
    {
      "id": "uuid",
      "name": "string",
      "age": "int8",
      "hobbies": "[]string"
    }
    ```
  - **400 Bad Request**: If ID format is invalid.
  - **404 Not Found**: If the person does not exist.

---

### 4. Update Person

- **Endpoint**: `{host}/api/v1/person/:id`
- **Method**: `PUT`
- **Path Parameters**:
  - `id` (uuid): The ID of the person to update.
- **Request Body**:

  ```json
  {
    "name": "string (min: 4, max: 100)",
    "age": "int8 (min: 0)",
    "hobbies": "[]string (optional)"
  }
  ```

- **Response**:
  - **200 OK**
    ```json
    {
      "id": "uuid",
      "name": "string",
      "age": "int8",
      "hobbies": "[]string"
    }
    ```
  - **400 Bad Request**: If input validation fails or ID format is invalid.
  - **404 Not Found**: If the person does not exist.

---

### 5. Delete Person

- **Endpoint**: `{host}/api/v1/person/:id`
- **Method**: `DELETE`
- **Path Parameters**:
  - `id` (uuid): The ID of the person to delete.
- **Response**:
  - **204 No Content**: Successfully deleted.
  - **400 Bad Request**: If ID format is invalid.
  - **404 Not Found**: If the person does not exist.

---

### CORS Configuration

- **Allow Origins**: All origins (to be restricted later).
- **Allowed Methods**: `GET`, `POST`, `PUT`, `DELETE`.
- **Allowed Headers**: `Origin`, `Content-Type`.
- **Expose Headers**: `Content-Length`.
- **Max Age**: 12 hours.
