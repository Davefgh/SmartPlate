package handlers

import (
	"log"
	"net/http"
	"strconv"
	"vehicle-api/internal/models"
	"vehicle-api/internal/repository"

	"github.com/labstack/echo/v4"
)

type VehicleHandler struct {
	repo *repository.VehicleRepository
}

func NewVehicleHandler(repo *repository.VehicleRepository) *VehicleHandler {
	return &VehicleHandler{repo: repo}
}

// CreateVehicle handles POST /vehicles
func (h *VehicleHandler) CreateVehicle(c echo.Context) error {
	var vehicle models.Vehicle
	if err := c.Bind(&vehicle); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if vehicle.Make == "" || vehicle.Model == "" || vehicle.Plate == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Make, model, and plate are required"})
	}

	if err := h.repo.Create(&vehicle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create vehicle"})
	}

	return c.JSON(http.StatusCreated, vehicle)
}

// GetAllVehicles handles GET /vehicles
// Change this in your VehicleHandler
func (h *VehicleHandler) GetAllVehicles(c echo.Context) error {
    vehicles, err := h.repo.GetAll()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    log.Printf("Returning %d vehicles", len(vehicles)) // Add this
    // Return direct array instead of wrapping in "items"
    return c.JSON(http.StatusOK, vehicles)
}

// GetVehicle handles GET /vehicles/:id
func (h *VehicleHandler) GetVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid vehicle ID"})
	}

	vehicle, err := h.repo.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Vehicle not found"})
	}

	return c.JSON(http.StatusOK, vehicle)
}

// UpdateVehicle handles PUT /vehicles/:id
func (h *VehicleHandler) UpdateVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid vehicle ID"})
	}

	var vehicle models.Vehicle
	if err := c.Bind(&vehicle); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	vehicle.ID = id
	if err := h.repo.Update(&vehicle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update vehicle"})
	}

	// Fetch the updated vehicle to return it with updated timestamp
	updatedVehicle, err := h.repo.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch updated vehicle"})
	}

	return c.JSON(http.StatusOK, updatedVehicle)
}

// DeleteVehicle handles DELETE /vehicles/:id
func (h *VehicleHandler) DeleteVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid vehicle ID"})
	}

	if err := h.repo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete vehicle"})
	}

	return c.NoContent(http.StatusNoContent)
}