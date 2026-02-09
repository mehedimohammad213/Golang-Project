package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	DefaultPageSize = 10
)

type ActionLink struct {
	Rel    string `json:"rel"`
	Method string `json:"method"`
	Href   string `json:"href"`
}

type PaginationLinks struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

type Pagination struct {
	CurrentPage int             `json:"current_page"`
	TotalPages  int             `json:"total_pages"`
	Limit       int             `json:"limit"`
	TotalItems  int64           `json:"total_items"`
	Links       PaginationLinks `json:"links"`
}

type Response struct {
	Message    string      `json:"message,omitempty"`
	Error      string      `json:"error,omitempty"`
	Hints      []string    `json:"hints,omitempty"`
	TrackID    string      `json:"track_id,omitempty"`
	Link       string      `json:"link,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// SuccessResponse sends a successful JSON response
func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	trackID := uuid.New().String()
	resp := Response{
		Message: message,
		Data:    data,
		TrackID: trackID,
	}

	if c.Request.Method == http.MethodGet {
		resp.Limit = 1
		resp.Pagination = &Pagination{
			CurrentPage: 1,
			TotalPages:  1,
			Limit:       1,
			TotalItems:  1,
			Links: PaginationLinks{
				Self: c.Request.URL.String(),
			},
		}
	}

	c.JSON(code, resp)
}

// SuccessResponseWithHints sends a successful JSON response with hints
func SuccessResponseWithHints(c *gin.Context, code int, message string, data interface{}, hints []string) {
	trackID := uuid.New().String()
	resp := Response{
		Message: message,
		Data:    data,
		Hints:   hints,
		TrackID: trackID,
	}

	if c.Request.Method == http.MethodGet {
		resp.Limit = 1
		resp.Pagination = &Pagination{
			CurrentPage: 1,
			TotalPages:  1,
			Limit:       1,
			TotalItems:  1,
			Links: PaginationLinks{
				Self: c.Request.URL.String(),
			},
		}
	}

	c.JSON(code, resp)
}

// SuccessResponseWithLink sends a successful JSON response with a link
func SuccessResponseWithLink(c *gin.Context, code int, message string, data interface{}, link string) {
	trackID := uuid.New().String()
	resp := Response{
		Message: message,
		Data:    data,
		Link:    link,
		TrackID: trackID,
	}

	if c.Request.Method == http.MethodGet {
		resp.Limit = 1
		resp.Pagination = &Pagination{
			CurrentPage: 1,
			TotalPages:  1,
			Limit:       1,
			TotalItems:  1,
			Links: PaginationLinks{
				Self: c.Request.URL.String(),
			},
		}
	}

	c.JSON(code, resp)
}

// PaginatedSuccessResponse sends a successful JSON response with pagination information
func PaginatedSuccessResponse(c *gin.Context, code int, message string, data interface{}, pagination Pagination) {
	trackID := uuid.New().String()
	resp := Response{
		Message:    message,
		Data:       data,
		Pagination: &pagination,
		Limit:      pagination.Limit,
		TrackID:    trackID,
	}
	c.JSON(code, resp)
}

// SuccessResponseWithLinks sends a successful JSON response with resource-level action links nested inside data
func SuccessResponseWithLinks(c *gin.Context, code int, message string, data interface{}, links []ActionLink) {
	trackID := uuid.New().String()

	payload := gin.H{
		"resource": data,
		"links":    links,
	}

	resp := Response{
		Message: message,
		Data:    payload,
		TrackID: trackID,
	}

	if c.Request.Method == http.MethodGet {
		resp.Limit = 1
		resp.Pagination = &Pagination{
			CurrentPage: 1,
			TotalPages:  1,
			Limit:       1,
			TotalItems:  1,
			Links: PaginationLinks{
				Self: c.Request.URL.String(),
			},
		}
	}

	c.JSON(code, resp)
}

// PaginatedSuccessWithLinksResponse sends a successful JSON response with pagination and nested action links
func PaginatedSuccessWithLinksResponse(c *gin.Context, code int, message string, data interface{}, pagination Pagination, links []ActionLink) {
	trackID := uuid.New().String()

	payload := gin.H{
		"items": data,
		"links": links,
	}

	resp := Response{
		Message:    message,
		Data:       payload,
		Pagination: &pagination,
		Limit:      pagination.Limit,
		TrackID:    trackID,
	}

	c.JSON(code, resp)
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c *gin.Context, code int, message string, err string) {
	trackID := uuid.New().String()
	c.JSON(code, Response{
		Message: message,
		Error:   err,
		TrackID: trackID,
	})
}

// ErrorResponseWithHints sends an error JSON response with hints
func ErrorResponseWithHints(c *gin.Context, code int, message string, err string, hints []string) {
	trackID := uuid.New().String()
	c.JSON(code, Response{
		Message: message,
		Error:   err,
		Hints:   hints,
		TrackID: trackID,
	})
}

// ErrorResponseWithLink sends an error JSON response with a documentation link
func ErrorResponseWithLink(c *gin.Context, code int, message string, err string, link string) {
	trackID := uuid.New().String()
	c.JSON(code, Response{
		Message: message,
		Error:   err,
		Link:    link,
		TrackID: trackID,
	})
}

// ValidationErrorResponse sends a validation error response with hints
func ValidationErrorResponse(c *gin.Context, err error) {
	trackID := uuid.New().String()
	var errs []string
	var hints []string

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			errs = append(errs, fmt.Sprintf("Field '%s' failed on the '%s' tag", e.Field(), e.Tag()))
		}
		hints = append(hints, "Check the API documentation for required field formats")
		hints = append(hints, "Ensure all required fields are provided")
	} else {
		errs = append(errs, err.Error())
	}

	c.JSON(http.StatusBadRequest, Response{
		Message: "Validation failed",
		Error:   fmt.Sprintf("%v", errs),
		Hints:   hints,
		TrackID: trackID,
	})
}
