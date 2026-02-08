package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/dto"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"
)

type PermissionHandler struct {
	Service service.PermissionService
}

func NewPermissionHandler(svc service.PermissionService) *PermissionHandler {
	return &PermissionHandler{Service: svc}
}

// CreatePermission godoc
// @Summary      Create a new permission
// @Description  Create a new permission with the input payload
// @Tags         permissions
// @Accept       json
// @Produce      json
// @Param        permission  body      dto.CreatePermissionRequest  true  "Permission JSON"
// @Success      201  {object}  dto.PermissionResponse
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /permissions [post]
// @Security     BearerAuth
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req dto.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	perm, err := h.Service.CreatePermission(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create permission", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Permission created successfully", perm)
}

// GetPermissions godoc
// @Summary      List all permissions
// @Description  Get a list of all permissions
// @Tags         permissions
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.PermissionResponse
// @Failure      500  {object}  utils.Response
// @Router       /permissions [get]
// @Security     BearerAuth
func (h *PermissionHandler) GetPermissions(c *gin.Context) {
	perms, err := h.Service.GetPermissions(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch permissions", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Permissions fetched successfully", perms)
}

// GetPermissionByID godoc
// @Summary      Get permission by ID
// @Description  Get details of a specific permission by its ID
// @Tags         permissions
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Permission ID"
// @Success      200  {object}  dto.PermissionResponse
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /permissions/{id} [get]
// @Security     BearerAuth
func (h *PermissionHandler) GetPermissionByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid permission ID", err.Error())
		return
	}

	perm, err := h.Service.GetPermissionByID(c.Request.Context(), id)
	if err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Permission not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch permission", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Permission fetched successfully", perm)
}

// UpdatePermission godoc
// @Summary      Update permission
// @Description  Update an existing permission by its ID
// @Tags         permissions
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Permission ID"
// @Param        permission body      dto.UpdatePermissionRequest true "Update Payload"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /permissions/{id} [put]
// @Security     BearerAuth
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid permission ID", err.Error())
		return
	}

	var req dto.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	if err := h.Service.UpdatePermission(c.Request.Context(), id, req); err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Permission not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update permission", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Permission updated successfully", nil)
}

// DeletePermission godoc
// @Summary      Delete permission
// @Description  Delete a permission by its ID
// @Tags         permissions
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Permission ID"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /permissions/{id} [delete]
// @Security     BearerAuth
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid permission ID", err.Error())
		return
	}

	if err := h.Service.DeletePermission(c.Request.Context(), id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete permission", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Permission deleted successfully", nil)
}
