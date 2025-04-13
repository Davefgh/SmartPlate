// internal/handlers/vehicle_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"smartplate-api/internal/models"
	"smartplate-api/internal/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VehicleHandler struct {
	repo repository.VehicleRepository // Changed from models to repository
}

func NewVehicleHandler(repo repository.VehicleRepository) *VehicleHandler { // Fixed here
	return &VehicleHandler{repo: repo}
}

func (h *VehicleHandler) CreateVehicle(c echo.Context) error {
	var vehicle models.Vehicle
	 // Use a strict decoder
	 decoder := json.NewDecoder(c.Request().Body)
	 decoder.DisallowUnknownFields()
	 
	 if err := decoder.Decode(&vehicle); err != nil {
		 return c.JSON(http.StatusBadRequest, map[string]string{
			 "error": "Invalid request payload: " + err.Error(),
		 })
	 }
	 
	 // Validate required fields
	 if vehicle.VEHICLE_CATEGORY == "" || vehicle.MV_FILE_NUMBER == "" {
		 return c.JSON(http.StatusBadRequest, map[string]string{
			 "error": "Missing required fields",
		 })
	 }

	createdVehicle, err := h.repo.CreateVehicle(c.Request().Context(), &vehicle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdVehicle)
}

func (h *VehicleHandler) GetVehicle(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle ID"})
	}

	vehicle, err := h.repo.GetVehicle(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "vehicle not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "vehicle not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, vehicle)
}

func (h *VehicleHandler) UpdateVehicle(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle ID"})
	}

	var vehicle models.Vehicle
	if err := c.Bind(&vehicle); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	if err := h.repo.UpdateVehicle(c.Request().Context(), id, &vehicle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, vehicle)
}

func (h *VehicleHandler) DeleteVehicle(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle ID"})
	}

	if err := h.repo.DeleteVehicle(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)

	
}



