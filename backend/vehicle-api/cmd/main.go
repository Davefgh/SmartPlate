package main

import (
	"log"
	"vehicle-api/internal/database"
	"vehicle-api/internal/handlers"
	"vehicle-api/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	vehicleRepo := repository.NewVehicleRepository(db)

	// Initialize handlers
	vehicleHandler := handlers.NewVehicleHandler(vehicleRepo)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for Vehicle CRUD
	e.POST("/vehicles", vehicleHandler.CreateVehicle)
	e.GET("/vehicles", vehicleHandler.GetAllVehicles)
	e.GET("/vehicles/:id", vehicleHandler.GetVehicle)
	e.PUT("/vehicles/:id", vehicleHandler.UpdateVehicle)
	e.DELETE("/vehicles/:id", vehicleHandler.DeleteVehicle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
