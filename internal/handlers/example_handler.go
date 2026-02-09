package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/user/car-project/internal/utils"
)

// ExampleHandler demonstrates the new response structure usage
type ExampleHandler struct{}

// Example 1: Basic success response
// Response will include: message, data, track_id
func (h *ExampleHandler) BasicSuccess(c *gin.Context) {
	data := map[string]interface{}{
		"id":   1,
		"name": "Example",
	}
	utils.SuccessResponse(c, http.StatusOK, "Operation successful", data)
}

// Example 2: Success response with hints
// Response will include: message, data, hints, track_id
func (h *ExampleHandler) SuccessWithHints(c *gin.Context) {
	data := map[string]interface{}{
		"status": "pending",
	}
	hints := []string{
		"This operation may take a few minutes to complete",
		"You can check the status using the /status endpoint",
	}
	utils.SuccessResponseWithHints(c, http.StatusAccepted, "Request accepted", data, hints)
}

// Example 3: Success response with documentation link
// Response will include: message, data, link, track_id
func (h *ExampleHandler) SuccessWithLink(c *gin.Context) {
	data := map[string]interface{}{
		"token": "abc123",
	}
	utils.SuccessResponseWithLink(
		c,
		http.StatusCreated,
		"Resource created successfully",
		data,
		"https://api.example.com/docs/resources",
	)
}

// Example 4: Basic error response
// Response will include: message, error, track_id
func (h *ExampleHandler) BasicError(c *gin.Context) {
	utils.ErrorResponse(
		c,
		http.StatusNotFound,
		"Resource not found",
		"The requested resource does not exist",
	)
}

// Example 5: Error response with hints
// Response will include: message, error, hints, track_id
func (h *ExampleHandler) ErrorWithHints(c *gin.Context) {
	hints := []string{
		"Ensure you have the correct permissions",
		"Check if the resource ID is valid",
		"Contact support if the issue persists",
	}
	utils.ErrorResponseWithHints(
		c,
		http.StatusForbidden,
		"Access denied",
		"You do not have permission to access this resource",
		hints,
	)
}

// Example 6: Error response with documentation link
// Response will include: message, error, link, track_id
func (h *ExampleHandler) ErrorWithLink(c *gin.Context) {
	utils.ErrorResponseWithLink(
		c,
		http.StatusBadRequest,
		"Invalid request format",
		"The request body is malformed",
		"https://api.example.com/docs/request-format",
	)
}

// Example 7: Paginated success response (with advanced query metadata)
// Response will include: message, data, pagination, filter, sort, orders, projection, search, track_id
func (h *ExampleHandler) PaginatedSuccess(c *gin.Context) {
	data := []map[string]interface{}{
		{"id": 1, "name": "Item 1", "links": []utils.ActionLink{{Rel: "self", Method: "GET", Href: "/api/v1/items/1"}}},
	}

	pagination := utils.Pagination{
		CurrentPage: 1,
		TotalPages:  1,
		Limit:       utils.DefaultPageSize,
		TotalItems:  1,
		Links: utils.PaginationLinks{
			Self:  "https://api.example.com/v1/items?page=1&limit=10&sort=name&search=item",
			First: "https://api.example.com/v1/items?page=1&limit=10&sort=name&search=item",
			Last:  "https://api.example.com/v1/items?page=1&limit=10&sort=name&search=item",
		},
	}

	trackID := uuid.New().String()
	c.JSON(http.StatusOK, utils.Response{
		Message:    "Items fetched successfully",
		Data:       data,
		Pagination: &pagination,
		Limit:      pagination.Limit,
		TrackID:    trackID,
	})
}

// Example 8: Product added successfully (Professional Feel)
// Links are nested under the data object
func (h *ExampleHandler) ProductAdded(c *gin.Context) {
	// Nested look (using helper that wraps)
	product := map[string]interface{}{
		"id":    123,
		"name":  "Professional Camera",
		"price": 1500,
	}

	links := []utils.ActionLink{
		{Rel: "self", Method: "GET", Href: "/api/v1/products/123"},
		{Rel: "next_step", Method: "GET", Href: "/api/v1/products/123/shipping-options"},
	}

	utils.SuccessResponseWithLinks(c, http.StatusCreated, "Product added successfully", product, links)
}

// Example 9: Product retrieved (Professional Feel)
// Links are flattened inside the data object for a sleeker look
func (h *ExampleHandler) ProductRetrieved(c *gin.Context) {
	links := []utils.ActionLink{
		{Rel: "self", Method: "GET", Href: "/api/v1/products/123"},
		{Rel: "next_step", Method: "POST", Href: "/api/v1/cart/add/123"},
	}

	// Flat look (manually merging for maximum control)
	product := gin.H{
		"id":    123,
		"name":  "Professional Camera",
		"price": 1500,
		"links": links,
	}

	utils.SuccessResponse(c, http.StatusOK, "Product fetched successfully", product)
}

// Example 10: Validation error (automatically includes hints)
// Response will include: message, error, hints, track_id
func (h *ExampleHandler) ValidationError(c *gin.Context) {
	type ExampleRequest struct {
		Email string `json:"email" binding:"required,email"`
		Age   int    `json:"age" binding:"required,min=18"`
	}

	var req ExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Validation passed", req)
}
