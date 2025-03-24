package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type vehicleHandler struct {
	pocketBaseURL string
	client        *http.Client
}

func NewVehicleHandler() *vehicleHandler {
	return &vehicleHandler{
		pocketBaseURL: "http://localhost:8090",
		client:        &http.Client{},
	}
}

func (h *vehicleHandler) GetItems(c echo.Context) error {
	pbURL := h.pocketBaseURL + "/api/collections/vehicle/records"

	req, err := http.NewRequest("GET", pbURL, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create request: " + err.Error(),
		})
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to contact PocketBase: " + err.Error(),
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var pbError map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&pbError); err != nil {
			return c.JSON(resp.StatusCode, map[string]string{
				"error": "Received status " + resp.Status,
			})
		}
		return c.JSON(resp.StatusCode, pbError)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to parse response: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}