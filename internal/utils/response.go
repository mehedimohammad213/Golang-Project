package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err string) {
	c.JSON(code, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func ValidationErrorResponse(c *gin.Context, err error) {
	var errs []string
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			errs = append(errs, fmt.Sprintf("Field '%s' failed on the '%s' tag", e.Field(), e.Tag()))
		}
	} else {
		errs = append(errs, err.Error())
	}

	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: "Validation failed",
		Data:    errs,
	})
}
