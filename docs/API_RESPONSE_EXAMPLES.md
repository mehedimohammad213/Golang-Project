# API Response Examples

This document shows real JSON responses from the standardized API response structure.

## Success Responses

### 1. Product Added Successfully (Professional Feel)
**Request:** `POST /api/v1/products`
```json
{
  "message": "Product added successfully",
  "track_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "data": {
    "id": 123,
    "name": "Professional Camera",
    "price": 1500,
    "links": [
      { "rel": "self", "method": "GET", "href": "/api/v1/products/123" },
      { "rel": "next_step", "method": "GET", "href": "/api/v1/products/123/shipping-options" }
    ]
  }
}
```

### 2. Product Retrieved (Professional Feel)
**Request:** `GET /api/v1/products/123`
```json
{
  "message": "Product fetched successfully",
  "track_id": "b2c3d4e5-f6a7-8901-bcde-f12345678901",
  "data": {
    "id": 123,
    "name": "Professional Camera",
    "price": 1500,
    "links": [
      { "rel": "self", "method": "GET", "href": "/api/v1/products/123" },
      { "rel": "next_step", "method": "POST", "href": "/api/v1/cart/add/123" }
    ]
  }
}
```

### 3. Users List Retrieved (Paginated List)
**Request:** `GET /api/users?page=1&limit=10`
```json
{
  "message": "Users fetched successfully",
  "track_id": "c3d4e5f6-a7b8-9012-cdef-123456789012",
  "limit": 10,
  "pagination": {
    "current_page": 1,
    "total_pages": 2,
    "limit": 10,
    "total_items": 20,
    "links": {
      "self": "http://localhost:8080/api/v1/users?page=1&limit=10",
      "first": "http://localhost:8080/api/v1/users?page=1&limit=10",
      "last": "http://localhost:8080/api/v1/users?page=2&limit=10",
      "next": "http://localhost:8080/api/v1/users?page=2&limit=10"
    }
  },
  "data": [
    {
      "id": 1,
      "username": "john_doe",
      "email": "john@example.com",
      "links": [
        { "rel": "self", "method": "GET", "href": "/api/v1/users/1" },
        { "rel": "next_step", "method": "GET", "href": "/api/v1/users/1/activity" }
      ]
    }
  ]
}
```

### 4. Product Retrieved (Mandatory Pagination on GET)
**Request:** `GET /api/v1/products/123`
```json
{
  "message": "Product fetched successfully",
  "track_id": "d4e5f6a7-b8c9-0123-def1-234567890123",
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "limit": 1,
    "total_items": 1,
    "links": {
      "self": "/api/v1/products/123"
    }
  },
  "data": {
    "id": 123,
    "name": "Professional Camera",
    "price": 1500,
    "links": [
      { "rel": "self", "method": "GET", "href": "/api/v1/products/123" },
      { "rel": "next_step", "method": "POST", "href": "/api/v1/cart/add/123" }
    ]
  }
}
```

### 5. Async Operation Accepted (with hints)
**Request:** `POST /api/export`
```json
{
  "message": "Export request accepted",
  "hints": [
    "This operation may take a few minutes to complete",
    "You can check the status using GET /api/export/{id}/status",
    "You will receive an email when the export is ready"
  ],
  "track_id": "d4e5f6a7-b8c9-0123-def1-234567890123",
  "data": {
    "export_id": "exp_123456",
    "status": "pending",
    "estimated_time": "5 minutes"
  }
}
```

## Error Responses

### 1. Resource Not Found
**Request:** `GET /api/users/999`
```json
{
  "message": "User not found",
  "error": "No user exists with ID 999",
  "track_id": "e5f6a7b8-c9d0-1234-ef12-345678901234"
}
```

### 2. Validation Error
**Request:** `POST /api/users` (with invalid data)
```json
{
  "message": "Validation failed",
  "error": "[Field 'Email' failed on the 'email' tag Field 'Age' failed on the 'min' tag]",
  "hints": [
    "Check the API documentation for required field formats",
    "Ensure all required fields are provided"
  ],
  "track_id": "f6a7b8c9-d0e1-2345-f123-456789012345"
}
```

## Field Presence Matrix

| Response Type | message | error | hints | track_id | link | pagination | links (body) | data |
|--------------|---------|-------|-------|----------|------|------------|--------------|------|
| Success | ✅ | ❌ | ❌ | ✅ | ❌ | ❌ | ❌ | ✅ |
| Success + Hints | ✅ | ❌ | ✅ | ✅ | ❌ | ❌ | ❌ | ✅ |
| Success + Link | ✅ | ❌ | ❌ | ✅ | ✅ | ❌ | ❌ | ✅ |
| Paginated Success | ✅ | ❌ | ❌ | ✅ | ❌ | ✅ | ❌ | ✅ |
| Success + Links | ✅ | ❌ | ❌ | ✅ | ❌ | ❌ | ✅ | ✅ |
| Error | ✅ | ✅ | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| Error + Hints | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ | ❌ |
| Error + Link | ✅ | ✅ | ❌ | ✅ | ✅ | ❌ | ❌ | ❌ |
| Validation Error | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ | ❌ |

## Notes

- `track_id` is **always present** in every response for debugging and tracking
- `message` is **always present** to provide a human-readable description
- `error` is **only present** in error responses (4xx, 5xx status codes)
- `data` is **only present** in successful responses (2xx status codes)
- `hints`, `link`, `pagination`, and `links` are **optional** and used when additional context is helpful
