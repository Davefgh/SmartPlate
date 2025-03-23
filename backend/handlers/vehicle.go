package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type vehicleHandler struct {
	pocketBaseURL string
	adminToken    string
	client        *http.Client
}

// Corrected constructor name to be exported
func NewVehicleHandler() *vehicleHandler {
	return &vehicleHandler{
		pocketBaseURL: "http://localhost:8090",
		adminToken:    os.Getenv("PB_ADMIN_TOKEN"),
		client:        &http.Client{},
	}
}

func (h *vehicleHandler) GetItems(c echo.Context) error {
	pbURL := h.pocketBaseURL + "/api/collections/vehicle/records"

	req, err := http.NewRequest("GET", pbURL, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Request creation failed"})
	}

	if h.adminToken != "" {
		req.Header.Add("Authorization", h.adminToken)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to contact PocketBase"})
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse response"})
	}

	return c.JSON(http.StatusOK, result)
}
