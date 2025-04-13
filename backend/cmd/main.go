package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"smartplate-api/internal/database"
	"smartplate-api/internal/handlers"
	"smartplate-api/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()
	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	// Initialize repositories and handlers
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

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
	//security suggestion??
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		ContentSecurityPolicy: "default-src 'self'",
	}))
	// Vehicle routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})
	e.POST("/users", userHandler.CreateUser)//working
	e.GET("/users", userHandler.GetAllUsers)//working
	e.GET("/users/:id", userHandler.GetUserByID)//working
	e.GET("/users/email/:email", userHandler.GetUserByEmail)
	e.PUT("/users/:id", userHandler.UpdateUser)	

	//for getting user by lto client id
	e.GET("/users/lto/:lto_client_id", userHandler.GetUserByLTOID)
	e.PUT("/users/by-lto/:lto_client_id", userHandler.UpdateUserByLTO)
	e.DELETE("/users/by-lto/:lto_client_id", userHandler.DeleteUserByLTO)
	//for generating lto client id
	e.GET("/generate-lto-id", userHandler.GenerateLTOID)  

	//for Vehicle routes
	e.POST("/api/vehicles", vehicleHandler.CreateVehicle)
	e.GET("/api/vehicles/:id", vehicleHandler.GetVehicle)
	e.PUT("/api/vehicles/:id", vehicleHandler.UpdateVehicle)
	e.DELETE("/api/vehicles/:id", vehicleHandler.DeleteVehicle)



	// Start server
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
    LogStatus: true,
    LogURI:    true,
    LogMethod: true,
    LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
        logger.Info().
            Str("URI", v.URI).
            Str("method", v.Method).
            Int("status", v.Status).
            Msg("request")
        return nil
    },
}))
fmt.Println("Registered routes:")
for _, route := range e.Routes() {
    fmt.Printf("%-6s %s\n", route.Method, route.Path)
}
// Then start the server
e.Logger.Fatal(e.Start(":8081"))
}

