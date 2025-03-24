// main.go
package main

import (
	"backend/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// CORS Configuration (Must be FIRST middleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	vehicleHandler := handlers.NewVehicleHandler()

	// Explicit OPTIONS handler
	e.OPTIONS("/vehicle", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	

	// Main endpoint
	e.GET("/vehicle", vehicleHandler.GetItems)

	e.Logger.Fatal(e.Start(":1323"))
}