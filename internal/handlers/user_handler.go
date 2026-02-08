package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserRequest  true  "User JSON"
// @Success      201  {object}  dto.UserResponse
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	user, err := h.Service.CreateUser(c.Request.Context(), req)
	if err != nil {
		if err == utils.ErrAlreadyExists {
			utils.ErrorResponse(c, http.StatusConflict, "User already exists", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User created successfully", user)
}

// GetUsers godoc
// @Summary      List all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.UserResponse
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.Service.GetUsers(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Users fetched successfully", users)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Tags         users
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dto.UserResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	user, err := h.Service.GetUserByID(c.Request.Context(), id)
	if err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch user", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User fetched successfully", user)
}

// UpdateUser godoc
// @Summary      Update user
// @Tags         users
// @Param        id   path      int  true  "User ID"
// @Param        user body      dto.UpdateUserRequest true "Update Payload"
// @Success      200  {object}  utils.Response
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	if err := h.Service.UpdateUser(c.Request.Context(), id, req); err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User updated successfully", nil)
}

// DeleteUser godoc
// @Summary      Delete user
// @Tags         users
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  utils.Response
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	if err := h.Service.DeleteUser(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}
