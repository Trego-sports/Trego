# User API

This package implements the user API endpoints for the Trego API Gateway.

## Architecture

The user API follows a layered architecture:

1. **Handler Layer** (`handler.go`) - HTTP request handling
2. **Service Layer** (`service.go`) - Business logic processing
3. **Repository Layer** (`repo.go`) - Data access (to be implemented)

### Handler Layer
- Located in: `trego-backend/api-gateway/service/v1/user/handler.go`
- Responsibilities:
  - HTTP request parsing and validation
  - Input validation (email format, required fields, data types)
  - Error handling and HTTP response formatting
  - Logging requests and responses
  - Calling service layer for business logic

### Service Layer
- Located in: `trego-backend/api-gateway/service/v1/user/service.go`
- Responsibilities:
  - Business logic processing
  - Business rule enforcement
  - Data transformation and enrichment
  - Calling repository layer for data access

### Repository Layer
- Location: To be created as `repo.go`
- Responsibilities (future):
  - Database query execution
  - Data persistence operations
  - Transaction management

## API Endpoints

### GET `/api/v1/user/email/:email`
Query a user by their email address (username).

**Parameters:**
- `email` (path parameter) - User's email address

**Example:**
```bash
GET /api/v1/user/email/john@example.com
```

**Response:**
```json
{
  "user_id": "uuid-here",
  "name": "John Doe",
  "email": "john@example.com",
  "reputation": 100,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### GET `/api/v1/user/:user_id`
Query a user by their user ID.

**Parameters:**
- `user_id` (path parameter) - User's unique identifier

### POST `/api/v1/user`
Create a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "picture_url": "https://example.com/picture.jpg",
  "phone_number": "+1234567890",
  "location": "Toronto, ON"
}
```

**Response:**
```json
{
  "user_id": "uuid-here",
  "name": "John Doe",
  "email": "john@example.com",
  "reputation": 0,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### PUT `/api/v1/user/:user_id`
Update an existing user's information.

**Parameters:**
- `user_id` (path parameter) - User's unique identifier

**Request Body** (all fields optional):
```json
{
  "name": "John Smith",
  "phone_number": "+1234567890",
  "location": "Vancouver, BC",
  "picture_url": "https://example.com/new-picture.jpg"
}
```

**Response:**
```json
{
  "user_id": "uuid-here",
  "name": "John Smith",
  "email": "john@example.com",
  "reputation": 100,
  "updated_at": "2023-01-02T00:00:00Z"
}
```

### GET `/api/v1/users`
List all users with pagination.

**Query Parameters:**
- `limit` (default: 10, max: 100) - Number of users per page
- `offset` (default: 0) - Number of users to skip

**Example:**
```bash
GET /api/v1/users?limit=20&offset=0
```

**Response:**
```json
{
  "users": [
    {
      "user_id": "uuid-1",
      "name": "John Doe",
      "email": "john@example.com"
    },
    {
      "user_id": "uuid-2",
      "name": "Jane Smith",
      "email": "jane@example.com"
    }
  ],
  "limit": 20,
  "offset": 0
}
```

## Implementation Status

- ✅ Handler layer with input validation
- ✅ Service layer with business logic placeholders
- ⏳ Repository layer (not yet implemented)

### Next Steps

1. Implement repository layer in `repo.go`:
   - Database connection handling
   - Query execution
   - Result mapping
   - Error handling

2. Implement service layer business logic:
   - Email uniqueness validation
   - User ID generation
   - Data enrichment
   - Business rule enforcement

3. Add additional endpoints if needed:
   - User sports management
   - User reputation updates
   - User preferences

## Business Logic Placeholders

The service layer currently contains placeholder comments indicating where business logic will be implemented:

- User validation and data processing
- Business rule enforcement
- Data transformation
- Repository layer integration

See `service.go` for detailed comments on what each function should do.

