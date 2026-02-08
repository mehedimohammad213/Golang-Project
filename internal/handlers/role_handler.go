package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"
)

type RoleHandler struct {
	Service service.RoleService
}

func NewRoleHandler(svc service.RoleService) *RoleHandler {
	return &RoleHandler{Service: svc}
}

// CreateRole godoc
// @Summary      Create a new role
// @Description  Create a new role with the input payload
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        role  body      dto.CreateRoleRequest  true  "Role JSON"
// @Success      201  {object}  dto.RoleResponse
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /roles [post]
// @Security     BearerAuth
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req dto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	role, err := h.Service.CreateRole(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create role", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Role created successfully", role)
}

// GetRoles godoc
// @Summary      List all roles
// @Description  Get a list of all roles
// @Tags         roles
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.RoleResponse
// @Failure      500  {object}  utils.Response
// @Router       /roles [get]
// @Security     BearerAuth
func (h *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := h.Service.GetRoles(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch roles", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Roles fetched successfully", roles)
}

// GetRoleByID godoc
// @Summary      Get role by ID
// @Description  Get details of a specific role by its ID
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  dto.RoleResponse
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /roles/{id} [get]
// @Security     BearerAuth
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid role ID", err.Error())
		return
	}

	role, err := h.Service.GetRoleByID(c.Request.Context(), id)
	if err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Role not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch role", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Role fetched successfully", role)
}

// UpdateRole godoc
// @Summary      Update role
// @Description  Update an existing role by its ID
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Param        role body      dto.UpdateRoleRequest true "Update Payload"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /roles/{id} [put]
// @Security     BearerAuth
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid role ID", err.Error())
		return
	}

	var req dto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	if err := h.Service.UpdateRole(c.Request.Context(), id, req); err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Role not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update role", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Role updated successfully", nil)
}

// DeleteRole godoc
// @Summary      Delete role
// @Description  Delete a role by its ID
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /roles/{id} [delete]
// @Security     BearerAuth
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid role ID", err.Error())
		return
	}

	if err := h.Service.DeleteRole(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete role", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Role deleted successfully", nil)
}

// AssignRole godoc
// @Summary      Assign role to user
// @Description  Assign a role to a user
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        assignment body dto.AssignRoleRequest true "Assignment Payload"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /roles/assign [post]
// @Security     BearerAuth
func (h *RoleHandler) AssignRole(c *gin.Context) {
	var req dto.AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	if err := h.Service.AssignRole(c.Request.Context(), req); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to assign role", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Role assigned successfully", nil)
}
