# Fiber REST Minimal

A clean, modular RESTful API boilerplate built with [Go](https://go.dev/) and [Fiber v3](https://github.com/gofiber/fiber), following Clean / Layered Architecture principles. It features GORM with SQLite, request validation using `go-playground/validator`, UUID primary keys, and consistent JSON response envelopes.

## Features

- **Framework**: High-performance HTTP server powered by [Fiber v3](https://github.com/gofiber/fiber).
- **Architecture**: Clean / Layered Architecture separating HTTP delivery, business logic (use cases), data access (repositories), and database models.
- **ORM & Database**: [GORM](https://gorm.io/) with SQLite driver, auto-migration, and automatic timestamp tracking.
- **Validation**: Request validation via [validator/v10](https://github.com/go-playground/validator) integrated directly into Fiber's struct validator.
- **Pagination**: Built-in reusable pagination helper and metadata response envelope (`page`, `per_page`, `total`, `total_pages`).
- **Standardized API Envelopes**: Consistent API responses for single entities, paginated collections, empty actions, and validation/resource errors.
- **UUID Identifiers**: Entity IDs auto-generated using [google/uuid](https://github.com/google/uuid).

## Tech Stack

- **Language**: [Go](https://go.dev/) (1.26+)
- **Web Framework**: [Fiber v3](https://github.com/gofiber/fiber) (v3.6.1)
- **ORM**: [GORM](https://gorm.io/) (v1.31.2)
- **Database**: [SQLite](https://www.sqlite.org/) (driver/sqlite v1.6.0)
- **Validation**: [validator/v10](https://github.com/go-playground/validator) (v10.30.5)
- **Testing**: [testify](https://github.com/stretchr/testify) (v1.12.1)
- **UUID Generator**: [google/uuid](https://github.com/google/uuid) (v1.6.0)

## Project Structure

```text
├── cmd/
│   └── main.go                     # Application entry point
├── internal/
│   ├── config/
│   │   ├── app.go                  # Dependency injection & bootstrap wiring
│   │   ├── fiber.go                # Fiber v3 app initialization & validator config
│   │   └── gorm.go                 # Database connection & auto-migration
│   ├── delivery/
│   │   ├── handler/
│   │   │   └── todo.go             # HTTP handlers (request binding & response mapping)
│   │   ├── request/
│   │   │   └── todo.go             # Request DTOs with validation rules
│   │   ├── response/
│   │   │   ├── error.go            # Standard error responses & validation error formatting
│   │   │   └── success.go          # Standard success & paginated collection responses
│   │   └── routes/
│   │       └── todo.go             # Route declarations
│   ├── entity/
│   │   ├── error.go                # Domain error definitions
│   │   ├── pagination.go           # Generic pagination domain entity
│   │   └── todo.go                 # Todo domain entity
│   ├── helper/
│   │   └── pagination.go           # Pagination helper functions and GORM scopes
│   ├── model/
│   │   └── todo.go                 # Database model with GORM hooks & entity mappers
│   ├── repository/
│   │   └── todo.go                 # Database operations layer
│   └── usecase/
│       └── todo/
│           ├── mock/
│           │   └── todo.go         # Use case mock for testing
│           ├── dto.go              # Use case parameter transfer objects
│           └── todo.go             # Core business logic implementation
├── test/
│   └── todo_test.go                # Todo API Test
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Architecture Flow

The project follows unidirectional dependency flow:

```
HTTP Request
     │
     ▼
[Routes] (internal/delivery/routes)
     │
     ▼
[Handler] (internal/delivery/handler)
  - Binds Request Body / Query & extracts route parameters
  - Validates structs via Fiber validator
     │
     ▼
[Usecase] (internal/usecase)
  - Executes business rules
  - Calculates pagination
     │
     ▼
[Repository] (internal/repository)
  - Interacts with DB via GORM
  - Applies pagination scopes
     │
     ▼
[Database] (SQLite / database.db)
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- A C compiler (e.g. `gcc`) and `CGO_ENABLED=1` — required by the CGo-based SQLite driver (`mattn/go-sqlite3`)

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/andruwizz/fiber-rest-minimal.git
   cd fiber-rest-minimal
   ```

2. Download Go dependencies:
   ```bash
   go mod download
   ```

### Running the Application

To run the application directly:

```bash
go run cmd/main.go
```

To build and run as a standalone binary:

```bash
go build -o bin/server cmd/main.go
./bin/server
```

The server listens on port `3000` (hardcoded in `cmd/main.go`): `http://localhost:3000`. On first start, the database file `database.db` is created and migrated automatically in the **current working directory**, so run the server from the same directory each time.

## API Endpoints

Base URL: `http://localhost:3000`

### Summary of Endpoints

| Method | Endpoint | Description |
|:-------|:---------|:------------|
| `POST` | `/todos` | Create a new todo |
| `GET` | `/todos` | List paginated todos |
| `GET` | `/todos/:id` | Get todo details by ID |
| `PUT` | `/todos/:id` | Update an existing todo |
| `DELETE` | `/todos/:id` | Delete a todo by ID |

---

### 1. Create Todo

- **Endpoint**: `POST /todos`
- **Headers**: `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "title": "Buy groceries",
    "done": false
  }
  ```
- **Response** (`201 Created`):
  ```json
  {
    "success": true,
    "data": {
      "id": "e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b",
      "title": "Buy groceries",
      "done": false,
      "created_at": 1730455721,
      "updated_at": 1730455721
    }
  }
  ```
- **cURL Example**:
  ```bash
  curl -X POST http://localhost:3000/todos \
    -H "Content-Type: application/json" \
    -d '{"title": "Buy groceries", "done": false}'
  ```

---

### 2. List Todos (Paginated)

- **Endpoint**: `GET /todos`
- **Query Parameters**:
  - `page` (optional, default: `1`, max: `100`): Page number
  - `limit` (optional, default: `10`, max: `100`): Items per page
  - Values `<= 0` or omitted fall back to the defaults; values above `100` or non-numeric values return `400`.
  - Results are not explicitly ordered.
- **Response** (`200 OK`):
  ```json
  {
    "success": true,
    "data": [
      {
        "id": "e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b",
        "title": "Buy groceries",
        "done": false,
        "created_at": 1730455721,
        "updated_at": 1730455721
      }
    ],
    "meta": {
      "page": 1,
      "per_page": 10,
      "total": 1,
      "total_pages": 1
    }
  }
  ```
- **cURL Example**:
  ```bash
  curl -X GET "http://localhost:3000/todos?page=1&limit=10"
  ```

---

### 3. Get Todo by ID

- **Endpoint**: `GET /todos/:id`
- **Response** (`200 OK`):
  ```json
  {
    "success": true,
    "data": {
      "id": "e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b",
      "title": "Buy groceries",
      "done": false,
      "created_at": 1730455721,
      "updated_at": 1730455721
    }
  }
  ```
- **cURL Example**:
  ```bash
  curl -X GET http://localhost:3000/todos/e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b
  ```

---

### 4. Update Todo

- **Endpoint**: `PUT /todos/:id`
- **Note**: This is a full replace — `title` is required, and an omitted `done` defaults to `false`.
- **Headers**: `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "title": "Buy groceries and cook dinner",
    "done": true
  }
  ```
- **Response** (`200 OK`):
  ```json
  {
    "success": true,
    "data": {
      "id": "e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b",
      "title": "Buy groceries and cook dinner",
      "done": true,
      "created_at": 1730455721,
      "updated_at": 1730459321
    }
  }
  ```
- **cURL Example**:
  ```bash
  curl -X PUT http://localhost:3000/todos/e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b \
    -H "Content-Type: application/json" \
    -d '{"title": "Buy groceries and cook dinner", "done": true}'
  ```

---

### 5. Delete Todo

- **Endpoint**: `DELETE /todos/:id`
- **Response** (`200 OK`):
  ```json
  {
    "success": true
  }
  ```
- **cURL Example**:
  ```bash
  curl -X DELETE http://localhost:3000/todos/e4b6bf5a-2735-4d08-a5ea-79fbb9fb811b
  ```

## Response Schema & Error Handling

### Validation Error (`400 Bad Request`)

When request validation fails (e.g. missing required `title`), the API returns structured error details:

```json
{
  "success": false,
  "error": {
    "message": "Field validation for 'Title' failed on the 'required' tag.",
    "errors": {
      "title": "Field validation for 'Title' failed on the 'required' tag."
    }
  }
}
```

If multiple fields fail validation, `message` will highlight the first error along with the count of remaining errors (e.g. `Field validation for 'Title' failed on the 'required' tag. (and 1 more)`).

### Invalid Payload Error (`400 Bad Request`)

When the request payload contains malformed JSON or cannot be bound:

```json
{
  "success": false,
  "error": {
    "message": "request payload is invalid"
  }
}
```

### Resource Not Found (`404 Not Found`)

When an entity is not found by ID (e.g. `GET`, `PUT`, or `DELETE` with an invalid ID):

```json
{
  "success": false,
  "error": {
    "message": "record not found"
  }
}
```

### Server Error (`500 Internal Server Error`)

When an operation fails (e.g. database failure). The underlying error is logged server-side and not exposed to the client:

```json
{
  "success": false,
  "error": {
    "message": "failed to execute operation"
  }
}
```

## Testing

Run the test suite:

```bash
go test ./test/
```

The tests exercise the HTTP handlers and routes against an in-memory mock use case (`internal/usecase/todo/mock`); the repository and SQLite layers are not covered.

## License

This project is licensed under the [MIT license](LICENSE).