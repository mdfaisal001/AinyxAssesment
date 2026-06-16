# User Management API

A RESTful User Management API built using Go, Fiber, PostgreSQL, SQLC, Uber Zap, and Validator.

## Tech Stack

- Go
- Fiber
- PostgreSQL
- SQLC
- Uber Zap
- Go Playground Validator

---

## Project Architecture

The project follows a layered architecture:

```text
Handler -> Service -> Repository -> Database
```

### Layers

- Handler: Handles HTTP requests and responses.
- Service: Contains business logic.
- Repository: Handles database operations using SQLC.
- Database: PostgreSQL.

---

## Features

- Create User
- Get All Users
- Get User By ID
- Update User
- Delete User
- Request Validation
- Structured Logging using Uber Zap
- Dynamic Age Calculation from Date of Birth
- SQLC for type-safe SQL queries

---

## Project Structure

```text
go-user-api/
├── cmd/
│   └── server/
│       └── main.go
│
├── config/
│   └── database.go
│
├── db/
│   ├── migration/
│   ├── query/
│   └── sqlc/
│       └── generated/
│
├── internal/
│   ├── handler/
│   ├── logger/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── service/
│
├── .env
├── go.mod
├── go.sum
├── sqlc.yaml
└── README.md
```

---

## Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

## Environment Variables

Create a `.env` file in the project root.

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=user_api
```

---

## Installation

### Clone Repository

```bash
git clone <repository-url>
cd go-user-api
```

### Install Dependencies

```bash
go mod tidy
```

### Create Database

```sql
CREATE DATABASE user_api;
```

### Create Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

### Generate SQLC Code

```bash
sqlc generate
```

### Run Application

```bash
go run cmd/server/main.go
```

Server starts on:

```text
http://localhost:3000
```

---

# API Endpoints

## Create User

### Request

```http
POST /users
```

### Body

```json
{
  "name": "Faisal",
  "dob": "2003-08-20"
}
```

### Response

```json
{
  "ID": 1,
  "Name": "Faisal",
  "Dob": "2003-08-20T00:00:00Z"
}
```

---

## Get All Users

### Request

```http
GET /users
```

### Response

```json
[
  {
    "id": 1,
    "name": "Faisal",
    "dob": "2003-08-20",
    "age": 22
  }
]
```

---

## Get User By ID

### Request

```http
GET /users/1
```

### Response

```json
{
  "id": 1,
  "name": "Faisal",
  "dob": "2003-08-20",
  "age": 22
}
```

---

## Update User

### Request

```http
PUT /users/1
```

### Body

```json
{
  "name": "Mohammed Faisal",
  "dob": "2003-08-20"
}
```

### Response

```json
{
  "ID": 1,
  "Name": "Mohammed Faisal",
  "Dob": "2003-08-20T00:00:00Z"
}
```

---

## Delete User

### Request

```http
DELETE /users/1
```

### Response

```http
204 No Content
```

---

## Validation

Request validation is implemented using:

```text
github.com/go-playground/validator/v10
```

Validation Rules:

- Name is required
- Date of Birth is required

---

## Logging

Structured logging is implemented using:

```text
go.uber.org/zap
```

Logs include:

- User Creation
- User Fetch
- User Update
- User Deletion
- Error Logging

---

## Dynamic Age Calculation

Age is not stored in the database.

Age is calculated dynamically from the user's Date of Birth using Go's `time` package whenever user details are fetched.

---

## Author

**Mohammed Faisal**

Backend Development Assessment Submission for Ainyx Solutions.
