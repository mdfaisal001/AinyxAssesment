## Additional Features

### Pagination

The `GET /users` endpoint supports pagination using query parameters.

#### Example Request

```http
GET /users?page=1&limit=5
```

#### Example Response

```json
{
  "page": 1,
  "limit": 5,
  "data": [
    {
      "id": 1,
      "name": "Faisal",
      "dob": "2003-08-20",
      "age": 22
    }
  ]
}
```

---

## Docker Support

### Build Docker Image

```bash
docker build -t go-user-api .
```

### Run Docker Container

```bash
docker run -p 3000:3000 --env-file .env go-user-api
```

Application will be available at:

```text
http://localhost:3000
```

### Docker Note

If PostgreSQL is running on the host machine and the application is running inside Docker, update:

```env
DB_HOST=host.docker.internal
```

to allow the container to communicate with PostgreSQL.

---

## Logging

Structured logging is implemented using Uber Zap.

### Logged Events

- User Creation
- User Retrieval
- User Update
- User Deletion
- Validation Errors
- Database Errors

### Example Log

```json
{
  "level": "info",
  "msg": "User created successfully",
  "user_id": 1,
  "name": "Faisal"
}
```

---

## Validation

Request validation is implemented using:

```text
github.com/go-playground/validator/v10
```

### Validation Rules

- Name is required
- Date of Birth is required

### Example Invalid Request

```json
{
  "name": "",
  "dob": ""
}
```

### Example Response

```json
{
  "error": "validation failed"
}
```

---

## Bonus Features Implemented

- Docker Support
- Pagination (`page` and `limit`)
- Structured Logging (Uber Zap)
- Request Validation
- Dynamic Age Calculation
