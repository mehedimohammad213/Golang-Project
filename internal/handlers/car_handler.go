package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/models"
	"github.com/user/car-project/internal/service"
	"github.com/user/car-project/internal/utils"
)

type CarHandler struct {
	Service service.CarService
}

func NewCarHandler(svc service.CarService) *CarHandler {
	return &CarHandler{Service: svc}
}

// CreateCar godoc
// @Summary      Create a new car
// @Description  Create a new car with the input payload
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        car  body      models.Car  true  "Car JSON"
// @Success      201  {object}  models.Car
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /cars [post]
// @Security     BearerAuth
func (h *CarHandler) CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	if err := h.Service.CreateCar(&car); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create car", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Car created successfully", car)
}

// GetCars godoc
// @Summary      List all cars
// @Description  Get a list of all cars
// @Tags         cars
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Car
// @Failure      500  {object}  utils.Response
// @Router       /cars [get]
// @Security     BearerAuth
func (h *CarHandler) GetCars(c *gin.Context) {
	cars, err := h.Service.GetCars()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch cars", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Cars fetched successfully", cars)
}

// GetCarByID godoc
// @Summary      Get a car by ID
// @Description  Get details of a specific car by its ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  models.Car
// @Failure      400  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /cars/{id} [get]
// @Security     BearerAuth
func (h *CarHandler) GetCarByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid car ID", err.Error())
		return
	}

	car, err := h.Service.GetCarByID(id)
	if err != nil {
		if err == utils.ErrNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Car not found", err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch car", err.Error())
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Car fetched successfully", car)
}

// UpdateCar godoc
// @Summary      Update a car
// @Description  Update an existing car by its ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id   path      int         true  "Car ID"
// @Param        car  body      models.Car  true  "Car JSON"
// @Success      200  {object}  models.Car
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /cars/{id} [put]
// @Security     BearerAuth
func (h *CarHandler) UpdateCar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid car ID", err.Error())
		return
	}

	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}
	car.ID = id

	if err := h.Service.UpdateCar(&car); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update car", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Car updated successfully", nil)
}

// DeleteCar godoc
// @Summary      Delete a car
// @Description  Delete a car by its ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  models.Car
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /cars/{id} [delete]
// @Security     BearerAuth
func (h *CarHandler) DeleteCar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid car ID", err.Error())
		return
	}

	if err := h.Service.DeleteCar(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete car", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Car deleted successfully", nil)
}
