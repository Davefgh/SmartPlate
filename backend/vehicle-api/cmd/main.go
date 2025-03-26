package main

import (
	"log"
	"net/http"
	"vehicle-api/internal/database"
	"vehicle-api/internal/handlers"
	"vehicle-api/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories and handlers
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleHandler := handlers.NewVehicleHandler(vehicleRepo)



	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Enhanced CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           3600,
	}))
	// Vehicle routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Vehicle API is running")
	})
	e.POST("/vehicles", vehicleHandler.CreateVehicle)
	e.GET("/vehicles", vehicleHandler.GetAllVehicles)
	e.GET("/vehicles/:id", vehicleHandler.GetVehicle)
	e.PUT("/vehicles/:id", vehicleHandler.UpdateVehicle)
	e.DELETE("/vehicles/:id", vehicleHandler.DeleteVehicle)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}