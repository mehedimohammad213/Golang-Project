package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	svc service.UserService
}

func NewAuthHandler(svc service.UserService) AuthHandler {
	return &authHandler{svc: svc}
}

// Login godoc
// @Summary User Login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "Login Credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/v1/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	resp, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		if err == utils.ErrUnauthorized {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid credentials", "")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "login failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "login successful", resp)
}
