// internal/handlers/vehicle_handler.go
package handlers

import (
	"net/http"
	"smartplate-api/internal/models"
	"smartplate-api/internal/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VehicleHandler struct {
	repo repository.VehicleRepository
}

// NewVehicleHandler creates a new VehicleHandler
func NewVehicleHandler(repo repository.VehicleRepository) *VehicleHandler {
	return &VehicleHandler{repo: repo}
}

func (h *VehicleHandler) GetAllVehicles(c echo.Context) error {
	vehicles, err := h.repo.GetAllVehicles(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, vehicles)
}

// CreateVehicle handles POST /api/vehicles
func (h *VehicleHandler) CreateVehicle(c echo.Context) error {
	var vehicle models.Vehicle
	if err := c.Bind(&vehicle); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	created, err := h.repo.CreateVehicle(c.Request().Context(), &vehicle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

// GetVehicle handles GET /api/vehicles/:id
func (h *VehicleHandler) GetVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle id"})
	}
	vehicle, err := h.repo.GetVehicle(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, vehicle)
}

// UpdateVehicle handles PUT /api/vehicles/:id
func (h *VehicleHandler) UpdateVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle id"})
	}
	var vehicle models.Vehicle
	if err := c.Bind(&vehicle); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.repo.UpdateVehicle(c.Request().Context(), id, &vehicle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// DeleteVehicle handles DELETE /api/vehicles/:id
func (h *VehicleHandler) DeleteVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle id"})
	}
	err = h.repo.DeleteVehicle(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *VehicleHandler) UpdatePlate(c echo.Context) error {
    // parse vehicle_id from URL
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid vehicle id"})
    }

    // bind into models.Plate
    var plate models.Plate
    if err := c.Bind(&plate); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    plate.VEHICLE_ID = id // ensure it matches URL

    if err := h.repo.UpdatePlate(c.Request().Context(), &plate); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.NoContent(http.StatusNoContent)
}
